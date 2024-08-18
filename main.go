package main

import (
	"context"
	usecase "solution_ch_part1/src/app/usecase"
	ohlcUC "solution_ch_part1/src/app/usecase/ohlc"
	"solution_ch_part1/src/infra/config"
	ms_log "solution_ch_part1/src/infra/log"
	"solution_ch_part1/src/infra/persistence/redis"
	"solution_ch_part1/src/interface/rest"

	"solution_ch_part1/src/infra/kafka"
	ohlcRedis "solution_ch_part1/src/infra/persistence/redis/usecase"
	consumer "solution_ch_part1/src/interface/kafka"
	avro "solution_ch_part1/src/interface/kafka/avro"

	_ "github.com/joho/godotenv/autoload"
)

type Transaction struct {
	Number   int
	Type     string
	Stock    string
	Quantity int
	Price    int
}

func main() {
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(conf.Redis, logger)
	ohlcRdb := ohlcRedis.NewsRedis(redisClient)

	// KAFKA
	avro := avro.GetKafkaAvroConfig() //avro topic
	kProducer, err := kafka.RegisterKafkaProducer(logger, kafka.KafkaConf(conf.Kafka), avro)
	if err != nil {
		logger.Fatal(err)
	}
	defer kProducer.Close()

	useCases := usecase.AllUseCases{
		OHLCUC: ohlcUC.NewOHLCUseCase(ohlcRdb, kProducer),
	}

	consumers := consumer.GetKafkaConsumers(logger, useCases)
	kConsumer, err := kafka.RegisterKafkaConsumer(logger, kafka.KafkaConf(conf.Kafka), consumers)
	if err != nil {
		logger.Fatal(err)
	}
	kafka.StartConsumer(kConsumer)
	defer kConsumer.Close()

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		useCases,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
