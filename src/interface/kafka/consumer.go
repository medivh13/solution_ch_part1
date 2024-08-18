package kafka

import (
	"os"

	consumers "solution_ch_part1/src/interface/kafka/consumer"

	usecases "solution_ch_part1/src/app/usecase"
	"solution_ch_part1/src/infra/kafka"

	"github.com/sirupsen/logrus"
)

func GetKafkaConsumers(logger *logrus.Logger, usecase usecases.AllUseCases) map[string]kafka.KafkaConsumersConfig {
	topic := os.Getenv("TOPIC_CONSUMER")

	topicMap := map[string]kafka.KafkaConsumersConfig{}

	if _, ok := topicMap[topic]; !ok {
		topicMap[topic] = kafka.KafkaConsumersConfig{
			Consumer: consumers.NewOHLCMessageConsumer(usecase.OHLCUC),
		}
	}

	return topicMap
}
