package topic

import (
	"fmt"

	"github.com/sophielizg/go-libs/utils"
)

type ConsumerHandler[V any, PV MessageValuePtr[V]] interface {
	HandleMessage(message *ConsumerMessage[V, PV]) error
}

type ConsumerSession[V any, PV MessageValuePtr[V]] interface {
	AckSuccess(message *ConsumerMessage[V, PV]) error
	AckError(message *ConsumerMessage[V, PV]) error
}

type ConsumerImplementation[V any, PV MessageValuePtr[V]] interface {
	Start(handler ConsumerHandler[V, PV]) <-chan error
	Pause() error
	Resume() error
	Errors() <-chan error
	Close() error
}

type ConsumerOption[V any, PV MessageValuePtr[V]] func(*ConsumerOptions[V, PV]) error

type ConsumerOptions[V any, PV MessageValuePtr[V]] struct {
	Implementation ConsumerImplementation[V, PV]
}

type Consumer[V any, PV MessageValuePtr[V]] struct {
	options *ConsumerOptions[V, PV]
}

func NewConsumer[V any, PV MessageValuePtr[V]](options ...ConsumerOption[V, PV]) (*Consumer[V, PV], error) {
	o := &ConsumerOptions[V, PV]{}

	for _, option := range options {
		if err := option(o); err != nil {
			return nil, err
		}
	}

	if utils.AnyNil(o.Implementation) {
		return nil, fmt.Errorf("NewConsumer: %w", ErrNilOption)
	}
	return &Consumer[V, PV]{o}, nil
}

func WithConsumerImplementation[V any, PV MessageValuePtr[V]](implementation ConsumerImplementation[V, PV]) func(*ConsumerOptions[V, PV]) error {
	return func(o *ConsumerOptions[V, PV]) error {
		o.Implementation = implementation
		return nil
	}
}

func (c *Consumer[V, PV]) Start(handler ConsumerHandler[V, PV]) <-chan error {
	return c.options.Implementation.Start(handler)
}

func (c *Consumer[V, PV]) Pause() error {
	return c.options.Implementation.Pause()
}

func (c *Consumer[V, PV]) Resume() error {
	return c.options.Implementation.Resume()
}

func (c *Consumer[V, PV]) Errors() <-chan error {
	return c.options.Implementation.Errors()
}

func (c *Consumer[V, PV]) Close() error {
	return c.options.Implementation.Close()
}
