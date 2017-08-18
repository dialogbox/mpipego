package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/dialogbox/mpipego"
)

func kafkaConsumer(conf *cluster.Config, brokers []string, group string, topics []string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, group, topics, conf)
	if err != nil {
		panic(err)
	}

	return consumer
}

func main() {
	flag.Parse()

	brokersPtr := flag.String("brokers", "localhost:9092", "CSV initial broker list")
	groupIDPtr := flag.String("groupid", "group01", "Group ID")
	converterNamePtr := flag.String("converter", "telegraf_json", "Format converter name")
	brokers := strings.Split(*brokersPtr, ",")

	if len(brokers) == 0 {
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		os.Exit(1)
	}

	topics := flag.Args()

	// indexer := mpipego.NewElasticIndexer()
	// indexer.Start()
	// defer indexer.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	conf := cluster.NewConfig()
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Consumer.Fetch.Default = 1024 * 1024
	conf.Consumer.Return.Errors = true
	conf.Group.Return.Notifications = true

	consumer := kafkaConsumer(conf, brokers, *groupIDPtr, topics)
	defer consumer.Close()

	converter := mpipego.ConverterForType(*converterNamePtr)

	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				metricData := converter(msg.Value)
				// indexer.Index(metricData)
				fmt.Println(metricData.Data)

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

	}

}
