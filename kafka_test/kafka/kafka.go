package mq

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"

	"github.com/Shopify/sarama"
)

type KConsumerTopic struct {
	// Topics 主题
	Topics []string
	// GroupHandler 消费者组 handler
	GroupHandler sarama.ConsumerGroupHandler
}

type KConsumer struct {
	ctx context.Context
	// key 为 kafka 消费者组
	consumerGroupTopic map[string]KConsumerTopic
	// kafka 客户端
	client sarama.Client
	// kafka 配置
	config *sarama.Config

	consumerGroup []sarama.ConsumerGroup
}

//func InitKafkaConsumer() *KConsumer {
//	//cfg := config.GetTomlConfig()
//	//gt := cfg.Kafka.GroupTopics
//	//gts := strings.Split(gt,";")
//	//ct := make(map[string]KConsumerTopic)
//	//
//	//for _,v := range gts {
//	//	tt := strings.Split(v,":")
//	//	group,topics := tt[0],strings.Split(tt[1],",")
//	//	groupHanlder,ok := sync.Syncs[group]
//	//	if !ok {
//	//		panic(fmt.Sprintf("kafka consumer group [%s] handler not register",group))
//	//	}
//	//	consumerTopic := KConsumerTopic{
//	//		Topics: topics,
//	//		GroupHandler: groupHanlder,
//	//	}
//	//	ct[group] = consumerTopic
//	//}
//
//	kafkaCfg := sarama.NewConfig()
//	kafkaCfg.Version = sarama.V2_7_0_0
//	kafkaCfg.Consumer.Return.Errors = true
//
//	if cfg.Kafka.ClientID != "" {
//		kafkaCfg.ClientID = cfg.Kafka.ClientID
//	}
//
//	if cfg.Kafka.Oldest {
//		kafkaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
//	}
//
//	return NewKafkaConsumer(context.Background(),[]string{cfg.Kafka.Addr},ct,kafkaCfg)
//}

func NewKafkaConsumer(ctx context.Context, addrs []string, groupTopic map[string]KConsumerTopic, config *sarama.Config) *KConsumer {
	consumer := &KConsumer{
		ctx:                ctx,
		consumerGroupTopic: groupTopic,
		config:             config,
		consumerGroup:      make([]sarama.ConsumerGroup, 0, len(groupTopic)),
	}

	if consumer.config == nil {
		cfg := sarama.NewConfig()
		cfg.ClientID = "sarama_test"
		cfg.Version = sarama.V0_11_0_0                     // kafka 版本号
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest // 重连时从最开始的地方消费
		cfg.Consumer.Return.Errors = true                  // 手动处理消费者错误
		consumer.config = cfg
	}

	client, err := sarama.NewClient(addrs, consumer.config)
	if err != nil {
		panic(err)
	}

	consumer.client = client

	consumer.setupGroup()

	return consumer
}

func (c *KConsumer) setupGroup() {
	for groupKey, gt := range c.consumerGroupTopic {
		consumerGroup, err := sarama.NewConsumerGroupFromClient(groupKey, c.client)
		if err != nil {
			panic(err)
		}

		c.consumerGroup = append(c.consumerGroup, consumerGroup)

		go c.trackError(consumerGroup)
		go c.consume(consumerGroup, gt.Topics, gt.GroupHandler)
	}
}

func (c *KConsumer) consume(group sarama.ConsumerGroup, topics []string, handler sarama.ConsumerGroupHandler) {
	for {
		err := group.Consume(c.ctx, topics, handler)
		if errors.Is(err, sarama.ErrClosedConsumerGroup) {
			return
		}

		if err != nil {
			log.Error().Err(err).Msg("KConsumer.consume Consume")
		}

		if err := c.ctx.Err(); err != nil {
			log.Error().Err(err).Msg("KConsumer consume ctx err")
			return
		}
	}
}

func (c *KConsumer) trackError(group sarama.ConsumerGroup) {
	for err := range group.Errors() {
		log.Error().Err(err).Msg("KConsumer.trackError Errors")
	}
}

func (c *KConsumer) Close() {
	for _, g := range c.consumerGroup {
		if err := g.Close(); err != nil {
			log.Error().Err(err).Msg("KConsumer.Close Close")
		}
	}
	_ = c.client.Close()
}
