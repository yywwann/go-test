package kafka

import (
	kafkav2 "github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Pubv2 struct {
	KafkaPub kafkav2.SyncProducer
	KafkaPubA kafkav2.AsyncProducer
}

func NewPubv2(brokers []string) kafkav2.SyncProducer {
	kc := kafkav2.NewConfig()
	kc.Producer.RequiredAcks = kafkav2.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                  // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafkav2.NewSyncProducer(brokers, kc)
	if err != nil {
		panic(err)
	}
	return pub
}

func NewPubv2A(brokers []string) kafkav2.AsyncProducer {
	kc := kafkav2.NewConfig()
	kc.Producer.RequiredAcks = kafkav2.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                  // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafkav2.NewAsyncProducer(brokers, kc)
	if err != nil {
		panic(err)
	}
	return pub
}

func (k *Pubv2) Send(key, topic, val string) error {
	m := &kafkav2.ProducerMessage{
		Key:   kafkav2.StringEncoder(key),
		Topic: topic,
		Value: kafkav2.ByteEncoder(val),
	}
	if _, _, err := k.KafkaPub.SendMessage(m); err != nil {
		log.Error().Err(err).Interface("msg", m).Msg("kafka v1 send")
	}
	return nil
}

//func (k *Pubv2) SendA(key, topic, val string) error {
//	msg := &kafkav2.ProducerMessage{
//		Key:   kafkav2.StringEncoder(key),
//		Topic: topic,
//		Value: kafkav2.ByteEncoder(val),
//	}
//	select {
//	case k.KafkaPubA.Input() <- msg:
//	case err := <-k.KafkaPubA.Errors():
//	case <-k.KafkaPubA:
//		doneCh <- struct{}{}
//	}
//	return nil
//}