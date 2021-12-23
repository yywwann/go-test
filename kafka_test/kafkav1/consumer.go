package kafkav1

import (
	cluster "github.com/bsm/sarama-cluster"
	"github.com/rs/zerolog/log"
)

type Process interface {
	Deal(m string) error
}

type Consumer struct {
	*cluster.Consumer
}

func (c *Consumer) Listen(index string, p Process) {
	for {
		select {
		case err := <-c.Errors():
			log.Error().Err(err).Msg("consumer error")
		case n := <-c.Notifications():
			log.Info().Interface("number", n).Msg("consumer rebalanced")
		case msg, ok := <-c.Messages():
			if !ok {
				return
			}
			log.Info().Str("index", index).Int32("partition", msg.Partition).Int64("offset", msg.Offset).Bytes("key", msg.Key).Msg("")

			_ = p.Deal(string(msg.Value))
			c.MarkOffset(msg, "")
		}
	}
}
