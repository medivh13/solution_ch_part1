package transaction

import (
	transaction "solution_ch_part1/src/app/proto"
	"strconv"
)

type TransactionReqDTO struct {
	Number           string `json:"order_number, omitempty"`
	Type             string `json:"type"`
	Stock            string `json:"stock_code"`
	OrderBook        string `json:"order_book"`
	Quantity         string `json:"quantity"`
	Price            string `json:"price"`
	ExecutedQuantity string `json:"executed_quantity, omitempty"`
	ExecutionPrice   string `json:"execution_price, omitempty"`
}

type TransactionDTO struct {
	Number    string `json:"order_number, omitempty"`
	Type      string `json:"type"`
	Stock     string `json:"stock_code"`
	OrderBook string `json:"order_book"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
}

func (dto *TransactionReqDTO) Transform() *TransactionDTO {
	var price, quantity int64
	if dto.Price != "" {
		i, _ := strconv.ParseInt(dto.Price, 10, 64)
		price = i
	} else {
		i, _ := strconv.ParseInt(dto.ExecutionPrice, 10, 64)
		price = i
	}

	if dto.Quantity != "" {
		i, _ := strconv.ParseInt(dto.Quantity, 10, 64)
		quantity = i
	} else {
		i, _ := strconv.ParseInt(dto.ExecutedQuantity, 10, 64)
		quantity = i
	}

	data := &TransactionDTO{
		Number:    dto.Number,
		Type:      dto.Type,
		Stock:     dto.Stock,
		OrderBook: dto.OrderBook,
		Quantity:  quantity,
		Price:     price,
	}
	return data
}

type TransactionRespDTO struct {
	PreviousPrice int64
	OpenPrice     int64
	HighestPrice  int64
	LowestPrice   int64
	ClosePrice    int64
	Volume        int64
	Value         int64
	AveragePrice  int64
}

func TransformRPC(data *transaction.TransactionReq) []*TransactionDTO {
	var resp []*TransactionDTO

	for _, val := range data.Transaction {
		var price, quantity int64
		if val.Price != "" {
			i, _ := strconv.ParseInt(val.Price, 10, 64)
			price = i
		} else {
			i, _ := strconv.ParseInt(val.ExecutionPrice, 10, 64)
			price = i
		}

		if val.Quantity != "" {
			i, _ := strconv.ParseInt(val.Quantity, 10, 64)
			quantity = i
		} else {
			i, _ := strconv.ParseInt(val.ExecutedQuantity, 10, 64)
			quantity = i
		}

		res := &TransactionDTO{
			Number:    val.Number,
			Type:      val.Type,
			Stock:     val.Stock,
			OrderBook: val.OrderBook,
			Quantity:  quantity,
			Price:     price,
		}

		resp = append(resp, res)
	}

	return resp
}
