package handler

import (
	transaction "solution_ch_part1/src/app/proto"
	usecases "solution_ch_part1/src/app/usecase/ohlc"
	"solution_ch_part1/src/infra/config"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	// book.BookServiceServer
	transaction.OHLCServiceServer
}

// Handler is struct
type Handler struct {
	config  *config.Config
	usecase usecases.OHLCUCInterface
	transaction.UnimplementedOHLCServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, uc usecases.OHLCUCInterface) *Handler {
	return &Handler{
		config: conf,
		// repo:       repo,
		// grpcClient: grpcClient,
		usecase: uc,
	}
}

var _ Interface = &Handler{}
