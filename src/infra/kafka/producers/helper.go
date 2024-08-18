package kafka_producers

import (
	"encoding/binary"
	"time"

	"github.com/Shopify/sarama"
)

func SetProducerMessage(
	topic string,
	value sarama.Encoder,
	key []byte,
	headers []sarama.RecordHeader,
) *sarama.ProducerMessage {
	if len(key) == 0 {
		randomTime := time.Now().UnixNano()
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(randomTime))
		key = b
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic:   topic,
		Value:   value,
		Key:     sarama.StringEncoder(key),
		Headers: headers,
	}
	return kafkaMsg
}
