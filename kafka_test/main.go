package main

import (
	"context"
	mq "example/kafka_test/kafka"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type A struct {
}

func (a A)Setup(_ sarama.ConsumerGroupSession) error {
	log.Info().Msg("setup")
	return nil
}

func (a A)Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Info().Msg("clean")
	return nil
}

func (a A)ConsumeClaim(cgs sarama.ConsumerGroupSession, cgc sarama.ConsumerGroupClaim) error {

	for msg := range cgc.Messages(){
		//fmt.Printf("%s\n", msg.Value)
		//
		log.Info().Int32("Partition", msg.Partition).
			Str("Topic", msg.Topic).
			Str("Value", string(msg.Value)).
			Msg("")
		//fmt.Println()
		cgs.MarkMessage(msg, "")
	}

	return nil
}

func main() {
	topic := mq.KConsumerTopic{
		Topics: []string{"test9093"},
		GroupHandler: A{},
	}

	kafkaAddress := []string{
		"42.192.50.232:9092",
		"42.192.50.232:9093",
		"42.192.50.232:9094",
		//"172.16.101.126:9092",
	}

	groupMap := map[string]mq.KConsumerTopic{
		"test": topic,
	}

	consumer := mq.NewKafkaConsumer(context.Background(), kafkaAddress, groupMap, nil)
	defer consumer.Close()
	//for {
	//	time.Sleep(time.Second)
	//	fmt.Println("xxx")
	//}
	//time.Sleep(10 * time.Second)
	select {

	}
}
