package topic

import "context"

type MessageValue interface {
	Encode() ([]byte, error)
	Decode(bytes []byte) error
}

type MessageValuePtr[V any] interface {
	*V
	MessageValue
}

type ProducerMessage struct {
	Value []byte
}

func NewProducerMessage[V MessageValue](value V) (*ProducerMessage, error) {
	bytes, err := value.Encode()
	if err != nil {
		return nil, err
	}

	return &ProducerMessage{
		Value: bytes,
	}, nil
}

type ConsumerMessage[V any, PV MessageValuePtr[V]] struct {
	Value   PV
	session ConsumerSession[V, PV]
	ctx     context.Context
}

func NewConsumerMessage[V any, PV MessageValuePtr[V]](ctx context.Context, session ConsumerSession[V, PV], encodedValue []byte) (*ConsumerMessage[V, PV], error) {
	var value V
	valuePtr := PV(&value)
	message := &ConsumerMessage[V, PV]{valuePtr, session, ctx}

	if err := message.Value.Decode(encodedValue); err != nil {
		return nil, err
	}
	return message, nil
}

func (m *ConsumerMessage[V, PV]) Success() error {
	if m.session == nil {
		return ErrNilSessionAck
	}
	return m.session.AckSuccess(m)
}

func (m *ConsumerMessage[V, PV]) Error() error {
	if m.session == nil {
		return ErrNilSessionAck
	}
	return m.session.AckError(m)
}

func (m *ConsumerMessage[V, PV]) Context() context.Context {
	return m.ctx
}
