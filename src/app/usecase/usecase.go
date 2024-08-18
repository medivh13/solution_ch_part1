package usecases

import (
	ohlcUC "solution_ch_part1/src/app/usecase/ohlc"
)

type AllUseCases struct {
	OHLCUC ohlcUC.OHLCUCInterface
}
