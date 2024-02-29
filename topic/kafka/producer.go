package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/sophielizg/go-libs/utils"
	"github.com/sophielizg/harvest/topic"
)

type ProducerOption = func(o *ProducerOptions)

type ProducerOptions struct {
	BaseOptions
}

type KafkaProducer struct {
	options  *ProducerOptions
	producer sarama.SyncProducer
}

func NewProducerImplementation(options ...ProducerOption) (*KafkaProducer, error) {
	o := &ProducerOptions{}
	utils.ApplyOptions(o, options...)
	if utils.AnyNil(o.Topic, o.Brokers) {
		return nil, fmt.Errorf("NewProducerImplementation: %w", ErrNilOption)
	}

	if o.Config == nil {
		o.Config = sarama.NewConfig()
	}

	p := &KafkaProducer{
		options: o,
	}
	var err error
	p.producer, err = sarama.NewSyncProducer(o.Brokers, o.Config)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func WithProducerImplementation[V topic.MessageValue](options ...ProducerOption) topic.ProducerOption {
	return func(producerOpts *topic.ProducerOptions) error {
		kafkaProducer, err := NewProducerImplementation(options...)
		if err != nil {
			return err
		}

		return topic.WithProducerImplementation[V](kafkaProducer)(producerOpts)
	}
}

func Producer(option BaseOption) ProducerOption {
	return func(o *ProducerOptions) {
		option(&o.BaseOptions)
	}
}

func (p *KafkaProducer) SendMessages(messages ...*topic.ProducerMessage) error {
	kafkaMessages := utils.Map(messages, func(message *topic.ProducerMessage) *sarama.ProducerMessage {
		return &sarama.ProducerMessage{
			Topic: p.options.Topic,
			Value: sarama.ByteEncoder(message.Value),
		}
	})

	return p.producer.SendMessages(kafkaMessages)
}

func (p *KafkaProducer) Close() error {
	return p.producer.Close()
}
