package mpipe

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/context"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

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

func runpipe(config *mpipConfig) {
	if len(config.kafkaConfig.brokers) == 0 {
		logrus.Fatal("No broker is provided")
		os.Exit(1)
	}

	if len(config.kafkaConfig.topics) == 0 {
		logrus.Fatal("No topic is specified")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	// trap SIGINT to trigger a shutdown and cancellation of context
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		cancel()
	}()

	indexer := &elasticIndexer{esConfig: config.esConfig}
	if !config.test {
		indexer.start(ctx)
	}

	consumer := kafkaConsumer(&config.kafkaConfig)
	converter, err := NewConverter(config.format)
	if err != nil {
		logrus.Fatalln(err)
	}

	for {
		select {
		case <-ctx.Done():
			go consumer.Close()
			return
		case msg, more := <-consumer.Messages():
			if more {
				go func() {
					metricData, err := converter.Convert(msg.Value)
					if err != nil {
						logrus.Errorf("can not convert data using %s, [%v]", converter.Name(), msg.Value)
						return
					}

					if !config.test {
						indexer.index(metricData)
						consumer.MarkOffset(msg, "") // mark message as processed
					} else {
						fmt.Println(metricData)
					}
				}()
			}
		case err, more := <-consumer.Errors():
			if more {
				logrus.Errorf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				logrus.Infof("Rebalanced: %+v\n", ntf)
			}
		}
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	mpipeCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mpipe.yaml)")
	mpipeCmd.PersistentFlags().BoolVar(&config.test, "test", false, "Test mode. No elasticsearch output. No kafka commit. Just print out to stdout.")
	mpipeCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Verbose output")

	mpipeCmd.PersistentFlags().StringP("groupid", "g", "group01", "Kafka consumer group ID")
	mpipeCmd.PersistentFlags().StringP("format", "f", "telegraf_json", "Type of metric data format")
	mpipeCmd.PersistentFlags().StringSliceP("topics", "t", []string{}, "list of topics")
	mpipeCmd.PersistentFlags().Bool("skip", false, "Skip all data in the topics")

	viper.BindPFlag("kafka.groupid", mpipeCmd.PersistentFlags().Lookup("groupid"))
	viper.BindPFlag("kafka.format", mpipeCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("kafka.topics", mpipeCmd.PersistentFlags().Lookup("topics"))
	viper.BindPFlag("kafka.skip", mpipeCmd.PersistentFlags().Lookup("skip"))

	viper.SetDefault("kafka.brokers", []string{"localhost:9092"})

	viper.SetDefault("elasticsearch.url", "http://localhost:9200?cluster.name=my-application")
	viper.SetDefault("elasticsearch.template", "telegraf_json")
	viper.SetDefault("elasticsearch.prefix", "metricpipe")

	viper.SetDefault("elasticsearch.batchSize", 500)
	viper.SetDefault("elasticsearch.flushInterval", 5*time.Second)
	viper.SetDefault("elasticsearch.workers", 4)
	viper.SetDefault("elasticsearch.bufferSize", 1000)
	viper.SetDefault("elasticsearch.reportInterval", time.Minute)
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
	config.esConfig.reportInterval = viper.GetDuration("elasticsearch.reportInterval")

	logrus.Debugf("Processed config: %v\n", config)
}
