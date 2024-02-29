package topic

import (
	"fmt"

	"github.com/sophielizg/go-libs/utils"
)

type ProducerImplementation interface {
	SendMessages(messages ...*ProducerMessage) error
	Close() error
}

type ProducerOptions struct {
	Implementation ProducerImplementation
}

type ProducerOption func(*ProducerOptions) error

type Producer[V MessageValue] struct {
	options *ProducerOptions
}

func NewProducer[V MessageValue](options ...ProducerOption) (*Producer[V], error) {
	o := &ProducerOptions{}

	for _, option := range options {
		if err := option(o); err != nil {
			return nil, err
		}
	}

	if utils.AnyNil(o.Implementation) {
		return nil, fmt.Errorf("NewProducer: %w", ErrNilOption)
	}
	return &Producer[V]{o}, nil
}

func WithProducerImplementation[V MessageValue](implementation ProducerImplementation) ProducerOption {
	return func(o *ProducerOptions) error {
		o.Implementation = implementation
		return nil
	}
}

func (p *Producer[V]) SendMessages(values ...V) error {
	messages, err := utils.MapWithError(values, NewProducerMessage[V])
	if err != nil {
		return err
	}
	return p.options.Implementation.SendMessages(messages...)
}

func (p *Producer[V]) Close() error {
	return p.options.Implementation.Close()
}
