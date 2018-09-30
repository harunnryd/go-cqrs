package broker

import (
	"fmt"
	"github.com/Shopify/sarama"
	"quizes/util"
	"reflect"
)

func Dispatch(topic string, command interface{}) {
	producer := util.NewKafkaSyncProducer()
	defer func() {
		if err:= producer.Close(); err != nil {
			fmt.Printf("Kafka err: %s\n", err)
		}
	}()

	cmd := util.Encode(command)

	prodMessage := new(sarama.ProducerMessage)
	prodMessage.Topic = topic
	prodMessage.Key = sarama.StringEncoder(reflect.TypeOf(command).Name())
	prodMessage.Value = sarama.StringEncoder(cmd)

	partition, offset, err := producer.SendMessage(prodMessage)
	if err != nil {
		fmt.Printf("Kafka err: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", command)
	fmt.Printf("Message stored in partition: %d, offset: %d\n", partition, offset)
}