package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sophielizg/harvest/topic"
)

type KafkaConsumerSession[V any, PV topic.MessageValuePtr[V]] struct {
	groupSession sarama.ConsumerGroupSession
}

func (s *KafkaConsumerSession[V, PV]) markMessage(message *topic.ConsumerMessage[V, PV]) error {
	metadata, err := GetMessageMetadataFromContext(message.Context())
	if err != nil {
		return err
	}

	s.groupSession.MarkOffset(metadata.GetTopic(), metadata.GetPartition(), metadata.GetOffset()+1, "")
	return nil
}

func (s *KafkaConsumerSession[V, PV]) AckSuccess(message *topic.ConsumerMessage[V, PV]) error {
	return s.markMessage(message)
}

func (s *KafkaConsumerSession[V, PV]) AckError(message *topic.ConsumerMessage[V, PV]) error {
	return s.markMessage(message)
}
