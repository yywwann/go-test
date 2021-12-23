package kafka

import (
	"github.com/rs/zerolog/log"
	kafkav1 "gopkg.in/Shopify/sarama.v1"
)

type Pubv1 struct {
	KafkaPub kafkav1.SyncProducer
}

func NewPubv1(brokers []string) kafkav1.SyncProducer {
	kc := kafkav1.NewConfig()
	kc.Producer.RequiredAcks = kafkav1.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                    // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafkav1.NewSyncProducer(brokers, kc)
	if err != nil {
		panic(err)
	}
	return pub
}

func (k *Pubv1) Send(key, topic, val string) error {
	m := &kafkav1.ProducerMessage{
		Key:   kafkav1.StringEncoder(key),
		Topic: topic,
		Value: kafkav1.ByteEncoder(val),
	}
	if _, _, err := k.KafkaPub.SendMessage(m); err != nil {
		log.Error().Err(err).Interface("msg", m).Msg("kafka v1 send")
	}
	return nil
}
