package kafka_producers

import (
	"encoding/binary"

	schema "solution_ch_part1/src/infra/kafka/producers/schema_registry"

	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro/v2"
	"github.com/sirupsen/logrus"
)

const Suffix = "-value"

type AvroProducer struct {
	Logger         *logrus.Logger
	Producer       sarama.SyncProducer
	SchemaRegistry *schema.CachedSchemaRegistryClient
}

func NewAvroProducer(logger *logrus.Logger, conf *sarama.Config, kafkaServers []string, schemaRegistryServers []string) (*AvroProducer, error) {
	producer, err := sarama.NewSyncProducer(kafkaServers, conf)
	if err != nil {
		return nil, err
	}
	schemaRegistryClient := schema.NewCachedSchemaRegistryClient(logger, schemaRegistryServers)
	return &AvroProducer{logger, producer, schemaRegistryClient}, nil
}

func (p *AvroProducer) Close() {
	p.Producer.Close()
}

func (ap *AvroProducer) GetSchemaId(topic string, avroCodec *goavro.Codec) (int, error) {
	schemaId, err := ap.SchemaRegistry.CreateSubject(topic+Suffix, avroCodec)
	if err != nil {
		return 0, err
	}
	return schemaId, nil
}

func (ap *AvroProducer) SendMessage(topic string, value []byte, key []byte, headers []sarama.RecordHeader) error {
	avroCodec, err := ap.SchemaRegistry.GetLatestSchema(topic + Suffix)
	if err != nil {
		return err
	}

	schemaId, err := ap.GetSchemaId(topic, avroCodec)
	if err != nil {
		return err
	}

	native, _, err := avroCodec.NativeFromTextual(value)
	if err != nil {
		return err
	}

	binaryValue, err := avroCodec.BinaryFromNative(nil, native)
	if err != nil {
		return err
	}

	binaryMsg := &AvroEncoder{
		SchemaID: schemaId,
		Content:  binaryValue,
	}
	msg := SetProducerMessage(topic, binaryMsg, key, headers)
	partition, offset, err := ap.Producer.SendMessage(msg)
	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return err
}

// AvroEncoder encodes schemaId and Avro message.
type AvroEncoder struct {
	SchemaID int
	Content  []byte
}

// Notice: the Confluent schema registry has special requirements for the Avro serialization rules,
// not only need to serialize the specific content, but also attach the Schema ID and Magic Byte.
// Ref: https://docs.confluent.io/current/schema-registry/serializer-formatter.html#wire-format
func (a *AvroEncoder) Encode() ([]byte, error) {
	var binaryMsg []byte
	// Confluent serialization format version number; currently always 0.
	binaryMsg = append(binaryMsg, byte(0))
	// 4-byte schema ID as returned by Schema Registry
	binarySchemaId := make([]byte, 4)
	binary.BigEndian.PutUint32(binarySchemaId, uint32(a.SchemaID))
	binaryMsg = append(binaryMsg, binarySchemaId...)
	// Avro serialized data in Avro's binary encoding
	binaryMsg = append(binaryMsg, a.Content...)
	return binaryMsg, nil
}

// Length of schemaId and Content.
func (a *AvroEncoder) Length() int {
	return 5 + len(a.Content)
}
