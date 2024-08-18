package avro

import (
	"solution_ch_part1/src/infra/kafka"
)

func GetKafkaAvroConfig() map[string]kafka.KafkaProducersAvroConfig {
	return map[string]kafka.KafkaProducersAvroConfig{
		"cho.common.otp.sendEmail": {
			Path: "./src/interface/kafka/avro/ohlc/ohlc.asvc",
		},
	}
}
