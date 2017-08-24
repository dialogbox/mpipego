package mpipe

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

type kafkaConfig struct {
	brokers []string
	groupid string
	format  string
	topics  []string
	skip    bool
}

type mpipConfig struct {
	kafkaConfig
	esConfig
	test bool
}

var config mpipConfig

var mpipeCmd = &cobra.Command{
	Use:   "mpipe",
	Short: "Metric forwarder",
	Long:  "mpipe is metric forwarder from karfa to elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		runpipe(&config)
	},
	Args: cobra.NoArgs,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := mpipeCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	mpipeCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mpipe.yaml)")
	mpipeCmd.PersistentFlags().BoolVar(&config.test, "test", false, "Test mode. No elasticsearch output. No kafka commit. Just print out to stdout.")
	mpipeCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Verbose output")

	mpipeCmd.PersistentFlags().StringP("groupid", "g", "group01", "Kafka consumer group ID")
	mpipeCmd.PersistentFlags().StringP("format", "f", "telegram_json", "Type of metric data format")
	mpipeCmd.PersistentFlags().StringSliceP("topics", "t", []string{}, "list of topics")
	mpipeCmd.PersistentFlags().Bool("skip", false, "Skip all data in the topics")

	viper.BindPFlag("kafka.groupid", mpipeCmd.PersistentFlags().Lookup("groupid"))
	viper.BindPFlag("kafka.format", mpipeCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("kafka.topics", mpipeCmd.PersistentFlags().Lookup("topics"))
	viper.BindPFlag("kafka.skip", mpipeCmd.PersistentFlags().Lookup("skip"))

	viper.SetDefault("elasticsearch.url", "elasticsearch://localhost:9300?cluster.name=my-application")
	viper.SetDefault("elasticsearch.template", "telegraf_json")
	viper.SetDefault("elasticsearch.prefix", "metricpipe")

	viper.SetDefault("elasticsearch.batchSize", 500)
	viper.SetDefault("elasticsearch.flushInterval", 5*time.Second)
	viper.SetDefault("elasticsearch.workers", 4)
	viper.SetDefault("elasticsearch.bufferSize", 1000)
}

func initConfig() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".mpipego" (without extension).
		viper.SetConfigName("mpipe")
		viper.AddConfigPath("/etc/mpipego/")
		viper.AddConfigPath("$HOME/.mpipego")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	config.kafkaConfig.brokers = viper.GetStringSlice("kafka.brokers")
	config.kafkaConfig.groupid = viper.GetString("kafka.groupid")
	config.kafkaConfig.format = viper.GetString("kafka.format")
	config.kafkaConfig.topics = viper.GetStringSlice("kafka.topics")

	config.esConfig.url = viper.GetString("elasticsearch.url")
	config.esConfig.template = viper.GetString("elasticsearch.template")
	config.esConfig.prefix = viper.GetString("elasticsearch.prefix")

	config.esConfig.batchSize = viper.GetInt("elasticsearch.batchSize")
	config.esConfig.flushInterval = viper.GetDuration("elasticsearch.flushInterval")
	config.esConfig.workers = viper.GetInt("elasticsearch.workers")
	config.esConfig.bufferSize = viper.GetInt("elasticsearch.bufferSize")

	logrus.Debugf("%v\n", config)
}

// func getConverters(cl []interface{}) ([]Converter, error) {
// 	var converters []Converter
// 	for i := range cl {
// 		c := cl[i].(map[string]interface{})
// 		name, ok := c["name"]
// 		if !ok {
// 			logrus.Fatalln("name must be provided for converter")
// 		}
// 		enabled, ok := c["enabled"]
// 		if ok && !enabled.(bool) {
// 			continue
// 		}
// 		converter, err := NewConverter(name.(string), c)
// 		if err != nil {
// 			logrus.Fatalln(err)
// 		}
// 		converters = append(converters, converter)
// 	}
// 	return converters, nil
// }

func kafkaConsumer(conf *cluster.Config, brokers []string, group string, topics []string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, group, topics, conf)
	if err != nil {
		logrus.Fatal(err)
	}

	return consumer
}

func runpipe(config *mpipConfig) {
	if len(config.kafkaConfig.brokers) == 0 {
		logrus.Fatal("No broker is provided")
		os.Exit(1)
	}

	if len(config.kafkaConfig.topics) == 0 {
		logrus.Fatal("No topic is specified")
		os.Exit(1)
	}

	indexer := NewElasticIndexer(&config.esConfig)
	if !config.test {
		indexer.Start()
		defer indexer.Stop()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumerConf := cluster.NewConfig()
	if config.kafkaConfig.skip {
		consumerConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		consumerConf.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumerConf.Consumer.Fetch.Default = 1024 * 1024
	consumerConf.Consumer.Return.Errors = true
	consumerConf.Group.Return.Notifications = true

	consumer := kafkaConsumer(consumerConf, config.kafkaConfig.brokers, config.groupid, config.topics)
	defer consumer.Close()

	converter, err := NewConverter(config.format)
	if err != nil {
		logrus.Fatalln(err)
	}

	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				metricData, err := converter.Convert(msg.Value)
				if err != nil {
					logrus.Errorf("can not convert data using %s, [%v]", converter.Name(), msg.Value)
					continue
				}

				if !config.test {
					indexer.Index(metricData)
					consumer.MarkOffset(msg, "") // mark message as processed
				} else {
					fmt.Println(metricData)
				}
			}
		case err, more := <-consumer.Errors():
			if more {
				logrus.Errorf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				logrus.Infof("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			return
		}

	}

}
