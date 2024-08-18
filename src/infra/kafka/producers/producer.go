package kafka_producers

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type CommonProducer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(
	producer sarama.SyncProducer,
) *CommonProducer {
	return &CommonProducer{
		Producer: producer,
	}
}

func (p *CommonProducer) Close() {
	p.Producer.Close()
}

func (p *CommonProducer) SendMessage(topic string, value []byte, key []byte, headers []sarama.RecordHeader) error {
	kafkaMsg := SetProducerMessage(
		topic,
		sarama.StringEncoder(value),
		key,
		headers,
	)

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return errors.WithStack(err)
}
