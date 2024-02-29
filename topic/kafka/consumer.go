package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/sophielizg/go-libs/utils"
	"github.com/sophielizg/harvest/topic"
)

type ConsumerOption = func(o *ConsumerOptions)

type ConsumerOptions struct {
	Group string
	BaseOptions
}

type KafkaConsumer[V any, PV topic.MessageValuePtr[V]] struct {
	options  *ConsumerOptions
	handler  *KafkaConsumerGroup[V, PV]
	consumer sarama.ConsumerGroup
}

func NewConsumerImplementation[V any, PV topic.MessageValuePtr[V]](options ...ConsumerOption) (*KafkaConsumer[V, PV], error) {
	o := &ConsumerOptions{}
	utils.ApplyOptions(o, options...)
	if utils.AnyNil(o.Topic, o.Brokers, o.Group) {
		return nil, fmt.Errorf("NewConsumerImplementation: %w", ErrNilOption)
	}

	if o.Config == nil {
		o.Config = sarama.NewConfig()
		o.Config.Consumer.Return.Errors = true
		o.Config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	c := &KafkaConsumer[V, PV]{options: o}
	var err error
	c.consumer, err = sarama.NewConsumerGroup(o.Brokers, o.Group, o.Config)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func WithGroup(group string) ConsumerOption {
	return func(o *ConsumerOptions) {
		o.Group = group
	}
}

func WithConsumerImplementation[V any, PV topic.MessageValuePtr[V]](options ...ConsumerOption) topic.ConsumerOption[V, PV] {
	return func(consumerOpts *topic.ConsumerOptions[V, PV]) error {
		kafkaConsumer, err := NewConsumerImplementation[V, PV](options...)
		if err != nil {
			return err
		}

		return topic.WithConsumerImplementation[V, PV](kafkaConsumer)(consumerOpts)
	}
}

func Consumer(option BaseOption) ConsumerOption {
	return func(o *ConsumerOptions) {
		option(&o.BaseOptions)
	}
}

func (c *KafkaConsumer[V, PV]) Start(handler topic.ConsumerHandler[V, PV]) <-chan error {
	c.handler = &KafkaConsumerGroup[V, PV]{handler}

	errs := make(chan error)
	go func() {
		for {
			err := c.consumer.Consume(context.Background(), []string{c.options.Topic}, c.handler)
			if err != nil {
				errs <- err
			}
		}
	}()

	return errs
}

func (c *KafkaConsumer[V, PV]) Pause() error {
	c.consumer.PauseAll()
	return nil
}

func (c *KafkaConsumer[V, PV]) Resume() error {
	c.consumer.ResumeAll()
	return nil
}

func (c *KafkaConsumer[V, PV]) Errors() <-chan error {
	return c.consumer.Errors()
}

func (c *KafkaConsumer[V, PV]) Close() error {
	return c.consumer.Close()
}
