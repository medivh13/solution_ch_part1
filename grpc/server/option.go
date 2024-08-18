package server

import (
	"solution_ch_part1/src/infra/config"
	
)

// WithConfig is function
func WithConfig(config *config.Config) ServerGrpcOption {
	return func(r *Server) {
		r.config = config
	}
}

