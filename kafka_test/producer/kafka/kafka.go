package kafka

type Producer interface {
	Send(key, topic, val string) error
}