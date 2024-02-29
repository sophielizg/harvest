package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

const MessageMetadataKey = "kafkaMessageMetadata"

type MessageMetadata struct {
	topic     string
	partition int32
	offset    int64
}

func (m MessageMetadata) GetTopic() string {
	return m.topic
}

func (m MessageMetadata) GetPartition() int32 {
	return m.partition
}

func (m MessageMetadata) GetOffset() int64 {
	return m.offset
}

func newContextWithMessage(ctx context.Context, message *sarama.ConsumerMessage) context.Context {
	return context.WithValue(ctx, MessageMetadataKey, MessageMetadata{
		message.Topic,
		message.Partition,
		message.Offset,
	})
}

func GetMessageMetadataFromContext(ctx context.Context) (MessageMetadata, error) {
	value := ctx.Value(MessageMetadataKey)
	if value == nil {
		return MessageMetadata{}, fmt.Errorf("GetMessageMetadataFromContext: %w", ErrNilContextMetadata)
	}

	if metadata, ok := value.(MessageMetadata); !ok {
		return MessageMetadata{}, fmt.Errorf("GetMessageMetadataFromContext: %w", ErrWrongContextMetadataType)
	} else {
		return metadata, nil
	}
}
