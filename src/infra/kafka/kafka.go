package kafka

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	kafka_producer "solution_ch_part1/src/infra/kafka/producers"
	schema "solution_ch_part1/src/infra/kafka/producers/schema_registry"
	"github.com/linkedin/goavro/v2"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type Consumer interface {
	Close()
	Consume(wait bool)
}

type Producer interface {
	Close()
	SendMessage(topic string, value []byte, key []byte, headers []sarama.RecordHeader) error
}

type KafkaConf struct {
	Username              string
	Password              string
	KafkaSchemaRegistry   string
	ConsumerGroupPrefix   string
	Brokers               string
	PartitionNo           string
	ReplicationMultiplier string
}

var _ Producer = &kafka_producer.AvroProducer{}
var _ Producer = &kafka_producer.CommonProducer{}

func RegisterKafkaProducer(
	logger *logrus.Logger,
	clientConf KafkaConf,
	avro map[string]KafkaProducersAvroConfig,
) (Producer, error) {
	conf := getKafkaConfig(clientConf.Username, clientConf.Password)

	brokers := strings.Split(clientConf.Brokers, ",")
	if clientConf.KafkaSchemaRegistry != "" {
		schema := strings.Split(clientConf.KafkaSchemaRegistry, ",")
		producers, err := kafka_producer.NewAvroProducer(logger, conf, brokers, schema)
		if err != nil {
			logrus.Errorf("Unable to create kafka producer got error %v", err)
			return nil, err
		}

		for topic, conf := range avro {
			avroFile, err := os.Open(conf.Path)
			if err != nil {
				logger.Fatalf("err: %s", errors.WithStack(err))
			}
			byteValue, _ := ioutil.ReadAll(avroFile)
			avroCodec, err := goavro.NewCodec(string(byteValue))
			if err != nil {
				logger.Fatal(err)
			}

			producers.SchemaRegistry.CreateSubject(topic+"-value", avroCodec)
			logger.Infof("%s schema has been registered", topic)
		}

		return producers, nil
	} else {
		producers, err := sarama.NewSyncProducer(brokers, conf)
		if err != nil {
			logrus.Errorf("Unable to create kafka producer got error %v", err)
			return nil, err
		}

		return kafka_producer.NewKafkaProducer(producers), nil
	}

}

func RegisterKafkaConsumer(
	logger *logrus.Logger,
	clientConf KafkaConf,
	consumerGroups map[string]KafkaConsumersConfig,
) (Consumer, error) {
	conf := getKafkaConfig(clientConf.Username, clientConf.Password)

	brokers := strings.Split(clientConf.Brokers, ",")
	client, err := sarama.NewClient(brokers, conf)
	if err != nil {
		logrus.Errorf("Unable to create kafka consumer got error %v", err)
		return nil, err
	}

	consumers, err := sarama.NewConsumerGroupFromClient(clientConf.ConsumerGroupPrefix, client)
	if err != nil {
		logrus.Errorf("Unable to create kafka consumer got error %v", err)
		return nil, err
	}

	schemaRegistryClient := schema.NewCachedSchemaRegistryClient(logger, []string{clientConf.KafkaSchemaRegistry})
	return NewKafkaConsumer(logger, clientConf, consumers, consumerGroups, make(chan os.Signal, 1), schemaRegistryClient), nil
}

func StartConsumer(cons Consumer, wait ...bool) {
	if len(wait) > 0 {
		cons.Consume(wait[0])
	} else {
		cons.Consume(true)
	}
}

func getKafkaConfig(username, password string) *sarama.Config {
	idInt := time.Now().UnixNano() / int64(time.Millisecond)
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 10
	kafkaConfig.ClientID = strconv.FormatInt(idInt, 10)

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
