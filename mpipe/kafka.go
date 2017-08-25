package mpipe

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
)

type kafkaConfig struct {
	brokers []string
	groupid string
	format  string
	topics  []string
	skip    bool
}

func kafkaConsumer(config *kafkaConfig) *cluster.Consumer {
	consumerConf := cluster.NewConfig()
	if config.skip {
		consumerConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		consumerConf.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumerConf.Consumer.Fetch.Default = 1024 * 1024
	consumerConf.Consumer.Return.Errors = true
	consumerConf.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(config.brokers, config.groupid, config.topics, consumerConf)
	if err != nil {
		logrus.Fatal(err)
	}

	return consumer
}
