package ohlc

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	dto "solution_ch_part1/src/app/dto/transaction"
	usecases "solution_ch_part1/src/app/usecase/ohlc"
	common_error "solution_ch_part1/src/infra/errors"
	"solution_ch_part1/src/interface/rest/response"
)

type OHLCHandlerInterface interface {
	Calculation(w http.ResponseWriter, r *http.Request)
}

type ohlcHandler struct {
	response response.IResponseClient
	usecase  usecases.OHLCUCInterface
}

func NewOHLCHandler(r response.IResponseClient, h usecases.OHLCUCInterface) OHLCHandlerInterface {
	return &ohlcHandler{
		response: r,
		usecase:  h,
	}
}

func (h *ohlcHandler) Calculation(w http.ResponseWriter, r *http.Request) {

	var postDTOs []*dto.TransactionDTO
	var postDTO *dto.TransactionReqDTO
	// err := json.NewDecoder(r.Body).Decode(&postDTO)
	d := json.NewDecoder(r.Body)
	for {

		err := d.Decode(&postDTO)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		postDTOs = append(postDTOs, postDTO.Transform())

	}

	err := h.usecase.Calculate(postDTOs)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Make Calculation",
		nil,
		nil,
	)
}
