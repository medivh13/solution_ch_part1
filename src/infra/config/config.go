package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type LogConf struct {
	Name string
}

type RedisConf struct {
	Host string
	Port string
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

// Config ...
type Config struct {
	App   AppConf
	Http  HttpConf
	Log   LogConf
	Redis RedisConf
	Kafka KafkaConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	redis := RedisConf{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	kafka := KafkaConf{
		Username:              os.Getenv("KAFKA_USERNAME"),
		Password:              os.Getenv("KAFKA_PASSWORD"),
		KafkaSchemaRegistry:   os.Getenv("KAFKA_SCHEMA_REGISTRY"),
		ConsumerGroupPrefix:   os.Getenv("KAFKA_CONSUMER_GROUP_PREFIX"),
		Brokers:               os.Getenv("KAFKA_BROKERS"),
		PartitionNo:           os.Getenv("KAFKA_PARTITION_NO"),
		ReplicationMultiplier: os.Getenv("KAFKA_REPLICATION_MULTIPLIER"),
	}
	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	config := Config{
		App:   app,
		Http:  http,
		Log:   log,
		Redis: redis,
		Kafka: kafka,
	}

	return config
}
