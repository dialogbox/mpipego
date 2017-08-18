package mpipego

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

type ElasticIndexer struct {
	ElasticOptions []elastic.ClientOptionFunc

	IndexPrefix string
	Template    string

	BatchSize     int
	FlushInterval time.Duration
	Workers       int

	BufferSize int
	processor  *elastic.BulkProcessor
	input      chan *MetricData
	quit       chan int
}

func NewElasticIndexer() *ElasticIndexer {
	e := ElasticIndexer{}

	e.IndexPrefix = "metricpipe"
	e.Template = "telegraf_json"

	e.BatchSize = 500
	e.FlushInterval = 5 * time.Second
	e.Workers = 4

	e.BufferSize = 1000

	e.quit = make(chan int, 1)

	return &e
}

func (es *ElasticIndexer) Index(m *MetricData) {
	es.input <- m
}

func (es *ElasticIndexer) Start() {
	ctx := context.Background()

	client, err := elastic.NewClient(es.ElasticOptions...)
	if err != nil {
		panic(err)
	}

	es.processor, err = client.BulkProcessor().
		Name("MetricIndexer").
		Workers(es.Workers).
		BulkActions(es.BatchSize).
		FlushInterval(es.FlushInterval).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	es.input = make(chan *MetricData, es.BufferSize)

	go func() {
		for {
			select {
			case m := <-es.input:
				indexName := fmt.Sprintf("%s-%s", es.IndexPrefix, m.Timestamp.Format("2006.01.02"))
				r := elastic.NewBulkIndexRequest().Index(indexName).Type(es.Template).Doc(m.Data)
				es.processor.Add(r)
			case <-es.quit:
				return
			}
		}

	}()
}

func (es *ElasticIndexer) Stop() {
	es.quit <- 0
	err := es.processor.Close()
	if err != nil {
		panic(err)
	}
	close(es.input)
	close(es.quit)
}
