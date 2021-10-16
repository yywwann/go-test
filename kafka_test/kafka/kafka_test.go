package mq

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

type A struct {
}

func (a A) Setup(_ sarama.ConsumerGroupSession) error {
	log.Info().Msg("setup")
	return nil
}

func (a A) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Info().Msg("clean")
	return nil
}

func (a A) ConsumeClaim(cgs sarama.ConsumerGroupSession, cgc sarama.ConsumerGroupClaim) error {

	for msg := range cgc.Messages() {
		//fmt.Printf("%s\n", msg.Value)
		//
		log.Info().Int32("Partition", msg.Partition).Int64("time", time.Now().UnixNano()).Msg("")
		//fmt.Println()
		cgs.MarkMessage(msg, "")
	}

	return nil
}

func TestNewKafkaConsumer(t *testing.T) {
	topic := KConsumerTopic{
		Topics:       []string{"test-topic"},
		GroupHandler: A{},
	}

	kafkaAddress := []string{
		"172.16.101.107:9092",
	}

	groupMap := map[string]KConsumerTopic{
		"group1": topic,
	}

	consumer := NewKafkaConsumer(context.Background(), kafkaAddress, groupMap, nil)
	defer consumer.Close()
	//for {
	//	time.Sleep(time.Second)
	//	fmt.Println("xxx")
	//}
	//time.Sleep(10 * time.Second)
	select {}
}

func BenchmarkNewKafkaConsumer(b *testing.B) {
	b.ResetTimer()
	topic := KConsumerTopic{
		Topics:       []string{"pubv2"},
		GroupHandler: A{},
	}

	kafkaAddress := []string{
		"127.0.0.1:9092",
	}

	groupMap := map[string]KConsumerTopic{
		"conv2": topic,
	}

	consumer := NewKafkaConsumer(context.Background(), kafkaAddress, groupMap, nil)
	defer consumer.Close()
}
