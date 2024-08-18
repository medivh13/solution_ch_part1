package main

import (
	"solution_ch_part1/grpc/server"
	"solution_ch_part1/src/infra/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf := config.Make()
	grpcServer := server.NewGRPCServer(

		server.WithConfig(&conf),
	)
	grpcServer.Run(9005)
}
