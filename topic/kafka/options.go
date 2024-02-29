package kafka

import "github.com/IBM/sarama"

type BaseOption = func(o *BaseOptions)

type BaseOptions struct {
	Topic   string
	Config  *sarama.Config
	Brokers []string
}

func WithConfig(config *sarama.Config) BaseOption {
	return func(o *BaseOptions) {
		o.Config = config
	}
}

func WithBrokers(brokerAddr ...string) BaseOption {
	return func(o *BaseOptions) {
		o.Brokers = brokerAddr
	}
}

func WithTopic(topic string) BaseOption {
	return func(o *BaseOptions) {
		o.Topic = topic
	}
}
