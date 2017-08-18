package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/dialogbox/mpipego/converters"
	"github.com/dialogbox/mpipego/outputs"
)

func benchmarker(signal chan int) {
	prev := time.Now()
	acc := 0

	for {
		count := <-signal

		if count < 0 {
			break
		}

		elapsed := time.Since(prev)
		prev = time.Now()

		fmt.Printf("%d in %s, total %d\n", count-acc, elapsed, count)
		acc = count
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func kafkaConsumer(conf *cluster.Config, brokers []string, group string, topics []string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, group, topics, conf)
	if err != nil {
		panic(err)
	}

	return consumer
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	indexer := outputs.NewElasticIndexer()
	indexer.Start()
	defer indexer.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	stopwatch := make(chan int)
	go benchmarker(stopwatch)

	conf := cluster.NewConfig()
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Fetch.Default = 1024 * 1024
	conf.Consumer.Return.Errors = true
	conf.Group.Return.Notifications = true

	consumer := kafkaConsumer(conf, []string{"localhost:9092"}, "GOTestGroup01", []string{"test"})
	defer consumer.Close()

	converter := converters.ConverterForType("telegraf_json")

	count := 0
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				metricData := converter(msg.Value)
				indexer.Index(metricData)

				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case err, more := <-consumer.Errors():
			if more {
				log.Printf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				log.Printf("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			return
		}

		count = count + 1
		if count%10000 == 0 {
			stopwatch <- count
		}
	}

}
