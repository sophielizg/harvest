package topic

import (
	"errors"
)

type Topic[V any, PV MessageValuePtr[V]] struct {
	Producer *Producer[PV]
	Consumer *Consumer[V, PV]
}

func (t *Topic[V, PV]) Close() error {
	return errors.Join(t.Producer.Close(), t.Consumer.Close())
}
