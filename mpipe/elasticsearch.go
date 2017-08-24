package mpipe

import (
	"fmt"
	"time"

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
}

type elasticIndexer struct {
	// ElasticOptions []elastic.ClientOptionFunc

	config esConfig

	processor *elastic.BulkProcessor
	input     chan *MetricData
	quit      chan int
}

func NewElasticIndexer(config *esConfig) *elasticIndexer {
	e := elasticIndexer{config: *config}

	e.quit = make(chan int, 1)

	return &e
}

func (es *elasticIndexer) Index(m *MetricData) {
	es.input <- m
}

func (es *elasticIndexer) Start() {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL(es.config.url))
	if err != nil {
		panic(err)
	}

	es.processor, err = client.BulkProcessor().
		Name("MetricIndexer").
		Workers(es.config.workers).
		BulkActions(es.config.batchSize).
		FlushInterval(es.config.flushInterval).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	es.input = make(chan *MetricData, es.config.bufferSize)

	go func() {
		for {
			select {
			case m := <-es.input:
				indexName := fmt.Sprintf("%s-%s", es.config.prefix, m.Timestamp.Format("2006.01.02"))
				r := elastic.NewBulkIndexRequest().Index(indexName).Type(es.config.template).Doc(m.Data)
				es.processor.Add(r)
			case <-es.quit:
				return
			}
		}

	}()
}

func (es *elasticIndexer) Stop() {
	es.quit <- 0
	err := es.processor.Close()
	if err != nil {
		panic(err)
	}
	close(es.input)
	close(es.quit)
}
