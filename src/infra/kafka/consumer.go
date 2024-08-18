package kafka

import (
	"context"
	"encoding/binary"
	"errors"
	"os"
	"time"

	"github.com/Shopify/sarama"
	schema "solution_ch_part1/src/infra/kafka/producers/schema_registry"
	"github.com/sirupsen/logrus"
)

type ClientConsumer interface {
	Consume(m *sarama.ConsumerMessage) error
}

type KafkaConsumersConfig struct {
	Consumer ClientConsumer
}

type KafkaProducersAvroConfig struct {
	Path string
}

type kafkaConsumer struct {
	Logger         *logrus.Logger
	Consumer       sarama.ConsumerGroup
	ConsumerGroups map[string]KafkaConsumersConfig
	Signals        chan os.Signal
	ClientConfig   KafkaConf
	SchemaRegistry *schema.CachedSchemaRegistryClient
}

func NewKafkaConsumer(
	logger *logrus.Logger,
	conf KafkaConf,
	consumer sarama.ConsumerGroup,
	consumerGroups map[string]KafkaConsumersConfig,
	signals chan os.Signal,
	schema *schema.CachedSchemaRegistryClient,
) Consumer {
	return &kafkaConsumer{
		Logger:         logger,
		Consumer:       consumer,
		ConsumerGroups: consumerGroups,
		Signals:        signals,
		ClientConfig:   conf,
		SchemaRegistry: schema,
	}
}

func (c *kafkaConsumer) Close() {
	c.Consumer.Close()
}

func (c *kafkaConsumer) Consume(wait bool) {
	var topics []string
	for k := range c.ConsumerGroups {
		topics = append(topics, k)
	}

	handler := consumerGroupHandler{
		Logger:         c.Logger,
		Signal:         c.Signals,
		ConsumerGroups: c.ConsumerGroups,
		IsUsingSchema:  c.ClientConfig.KafkaSchemaRegistry != "",
		Schema:         c.SchemaRegistry,
		Ready:          make(chan bool),
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				time.Sleep(time.Second)
				c.Logger.Info("Rejoin Kafka Consumer")
				c.Consume(wait)
			}
		}()
		if err := c.Consumer.Consume(context.Background(), topics, &handler); err != nil {
			// When setup fails, error will be returned here
			c.Logger.Errorf("Error from consumer: %v", err)
			panic(err)
		}
	}()
	if wait {
		<-handler.Ready
	}

}

type ConsumerGroupReady struct {
	Ready bool
}

type consumerGroupHandler struct {
	Logger         *logrus.Logger
	Signal         chan os.Signal
	ConsumerGroups map[string]KafkaConsumersConfig
	IsUsingSchema  bool
	Schema         *schema.CachedSchemaRegistryClient
	Ready          chan bool
}

func (h consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	h.Logger.Infof("Setup Consumer")
	h.Ready <- true
	return nil
}
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	h.Logger.Infof("Kafka consumer %s is started . . .", claim.Topic())

ConsumerLoop:
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				break ConsumerLoop
			}
			if msg != nil {
				sess.MarkMessage(msg, "")
			}
			if h.IsUsingSchema {
				schemaId := binary.BigEndian.Uint32(msg.Value[1:5])
				codec, err := h.Schema.GetSchema(int(schemaId))
				if err != nil {
					h.Logger.Error(err)
					return err
				}

				native, _, err := codec.NativeFromBinary(msg.Value[5:])
				if err != nil {
					h.Logger.Error(err)
					return err
				}

				textual, err := codec.TextualFromNative(nil, native)

				if err != nil {
					h.Logger.Error(err)
					return err
				}

				msg = &sarama.ConsumerMessage{
					Headers:        msg.Headers,
					Timestamp:      msg.Timestamp,
					BlockTimestamp: msg.BlockTimestamp,
					Key:            msg.Key,
					Value:          textual,
					Topic:          msg.Topic,
					Partition:      msg.Partition,
					Offset:         msg.Offset,
				}
			}

			h.ConsumerGroups[claim.Topic()].Consumer.Consume(msg)
		case sig := <-h.Signal:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
	return nil

}
func (h consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	panic(errors.New("cleanup"))
}
