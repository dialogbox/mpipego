package mpipe

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

type esConfig struct {
	url      string
	template string
	prefix   string

	batchSize     int
	flushInterval time.Duration
	workers       int

	bufferSize int

	reportInterval time.Duration
}

type elasticIndexer struct {
	esConfig

	input chan *MetricData
}

func (es *elasticIndexer) index(m *MetricData) {
	es.input <- m
}

type logrusWrapper struct {
	log func(format string, v ...interface{})
}

func (lw *logrusWrapper) Printf(format string, v ...interface{}) {
	lw.log(format, v...)
}

func (es *elasticIndexer) start(ctx context.Context) {
	errLogger := &logrusWrapper{logrus.Errorf}
	infoLogger := &logrusWrapper{logrus.Debugf}

	client, err := elastic.NewClient(
		elastic.SetURL(es.url),
		elastic.SetErrorLog(errLogger),
		elastic.SetInfoLog(infoLogger),
		elastic.SetHealthcheckInterval(20*time.Second),
	)
	if err != nil {
		panic(err)
	}

	processor, err := client.BulkProcessor().
		Name("MetricIndexer").
		Workers(es.workers).
		BulkActions(es.batchSize).
		FlushInterval(es.flushInterval).
		Stats(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	es.input = make(chan *MetricData, es.bufferSize)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(es.input)
				go func() {
					processor.Close()
				}()
				return
			case m := <-es.input:
				indexName := fmt.Sprintf("%s-%s", es.prefix, m.Timestamp.Format("2006.01.02"))
				r := elastic.NewBulkIndexRequest().Index(indexName).Type(es.template).Doc(m.Data)
				processor.Add(r)
			}
		}
	}()

	// reporter goroutine
	go func() {
		var prevStat elastic.BulkProcessorStats
		for {
			select {
			case <-time.After(es.reportInterval):
				stat := processor.Stats()
				reportStat(&prevStat, &stat)
				prevStat = stat
			case <-ctx.Done():
				return
			}
		}
	}()
}

func reportStat(prev *elastic.BulkProcessorStats, stats *elastic.BulkProcessorStats) {
	logrus.Infof("------------------------------------------")
	logrus.Infof("# of flush: %d, commit: %d",
		stats.Flushed-prev.Flushed,
		stats.Committed-prev.Committed,
	)
	logrus.Infof("# of reqs indexed: %d, created: %d, updated: %d,  successed: %d",
		stats.Indexed-prev.Indexed,
		stats.Created-prev.Created,
		stats.Updated-prev.Updated,
		stats.Succeeded-prev.Succeeded,
	)

	if stats.Failed-prev.Failed > 0 {
		logrus.Infof("# of reqs failed: %d", stats.Failed-prev.Failed)
	}

	for i, w := range stats.Workers {
		logrus.Infof("[Worker %3d] # of reqs queued: %d, Last response time: %v", i, w.Queued, w.LastDuration)
	}
}
