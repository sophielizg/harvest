package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sophielizg/harvest/topic"
)

type KafkaConsumerGroup[V any, PV topic.MessageValuePtr[V]] struct {
	handler topic.ConsumerHandler[V, PV]
}

func (h *KafkaConsumerGroup[V, PV]) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaConsumerGroup[V, PV]) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaConsumerGroup[V, PV]) ConsumeClaim(groupSession sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	session := &KafkaConsumerSession[V, PV]{groupSession}
	for {
		select {
		case claimMessage, more := <-claim.Messages():
			ctx := newContextWithMessage(groupSession.Context(), claimMessage)
			message, err := topic.NewConsumerMessage[V, PV](ctx, session, claimMessage.Value)
			if err != nil {
				return err
			}

			if err := h.handler.HandleMessage(message); err != nil {
				return err
			}

			if !more {
				return nil
			}
		case <-groupSession.Context().Done():
			return nil
		}
	}
}
