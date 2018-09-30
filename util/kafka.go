package util

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"os"
)

var (
	brokers = []string{"kafka:9092"}
)

func producerConfig() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func NewKafkaSyncProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(brokers, producerConfig())
	if err != nil {
		fmt.Printf("Kafka err: %s\n", err)
		os.Exit(-1)
	}
	return producer
}

func consumerConfig() *cluster.Config {
	conf := cluster.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Group.Return.Notifications = true
	return conf
}

func NewKafkaConsumer(groupID string, topics ...string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, groupID, topics, consumerConfig())
	if err != nil {
		fmt.Printf("Kafka err: %s\n", err)
		os.Exit(-1)
	}
	return consumer
}