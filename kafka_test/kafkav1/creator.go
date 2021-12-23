package kafkav1

import (
	cluster "github.com/bsm/sarama-cluster"
	"strconv"
)

func NewKafkaConsumers(topic string, group string, consumerNum int, brokers []string) map[string]*Consumer {
	store := make(map[string]*Consumer)
	num := int(consumerNum)
	for i := 0; i < num; i++ {
		store[strconv.Itoa(i)] = &Consumer{Consumer: newKafkaSub(topic, group, brokers)}
	}
	return store
}

func newKafkaSub(topic string, group string, brokers []string) *cluster.Consumer {
	c := cluster.NewConfig()
	c.Consumer.Return.Errors = true
	c.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(brokers, group, []string{topic}, c)
	if err != nil {
		panic(err)
	}
	return consumer
}
