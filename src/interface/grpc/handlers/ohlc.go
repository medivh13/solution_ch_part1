package handler

import (
	"context"
	"log"

	dto "solution_ch_part1/src/app/dto/transaction"
	transaction "solution_ch_part1/src/app/proto"
)

func (c *Handler) Calculate(ctx context.Context, req *transaction.TransactionReq) (*transaction.TransactionResp, error) {

	data := dto.TransformRPC(req)

	resp, err := c.usecase.CalculateFromRPC(ctx, data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
