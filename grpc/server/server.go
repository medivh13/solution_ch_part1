package server

import (
	"solution_ch_part1/src/infra/config"
	
	handler "solution_ch_part1/src/interface/grpc/handlers"
	
	transaction "solution_ch_part1/src/app/proto"
	usecases "solution_ch_part1/src/app/usecase/ohlc"

	"solution_ch_part1/src/infra/persistence/redis"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"solution_ch_part1/src/infra/kafka"
	ms_log "solution_ch_part1/src/infra/log"
	ohlcRedis "solution_ch_part1/src/infra/persistence/redis/usecase"
	avro "solution_ch_part1/src/interface/kafka/avro"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
}

type ServerGrpcOption func(*Server)

// NewGRPCServer is constructor
// func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
func NewGRPCServer(options ...ServerGrpcOption) *Server {
	server := &Server{}

	for _, option := range options {
		option(server)
	}

	return server
}

// Run is a method gRPC server
func (s *Server) Run(port int) error {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(

		),
		grpc.ChainStreamInterceptor(
		
		),
	)

	m := make(map[string]interface{})
	m["env"] = s.config.App.Environment
	m["service"] = s.config.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(s.config.Log.Name),
		ms_log.IsProduction(false),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(s.config.Redis, logger)
	ohlcRdb := ohlcRedis.NewsRedis(redisClient)

	// KAFKA
	avro := avro.GetKafkaAvroConfig() //avro topic
	kProducer, err := kafka.RegisterKafkaProducer(logger, kafka.KafkaConf(s.config.Kafka), avro)
	if err != nil {
		logger.Fatal(err)
	}
	defer kProducer.Close()

	uc := usecases.NewOHLCUseCase(ohlcRdb, kProducer)
	handlers := handler.NewHandler(s.config, uc)

	// register from proto
	transaction.RegisterOHLCServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
