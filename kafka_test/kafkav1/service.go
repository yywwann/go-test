package kafkav1

import (
	"github.com/rs/zerolog/log"
)

type Service struct {
	consumers map[string]*Consumer
}

func New(topic string, group string, consumerNum int, brokers []string) *Service {
	s := &Service{
		consumers: NewKafkaConsumers(topic, group, consumerNum, brokers),
	}
	return s
}

func (s *Service) ListenMQ() {
	for i, c := range s.consumers {
		log.Info().Str("index", i).Msg("")
		go c.Listen(i, s)
	}
}

func (s *Service) Deal(msg string) error {
	return nil
}
