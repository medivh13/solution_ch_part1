package ohlc

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	dto "solution_ch_part1/src/app/dto/transaction"
	transaction "solution_ch_part1/src/app/proto"
	"solution_ch_part1/src/infra/kafka"
	redis "solution_ch_part1/src/infra/persistence/redis/usecase"
	"time"

	"github.com/Shopify/sarama"
)

type OHLCUCInterface interface {
	Calculate(data []*dto.TransactionDTO) error
	ConsumeAndCalculate(ctx context.Context, data []*dto.TransactionDTO) error
	CalculateFromRPC(ctx context.Context, data []*dto.TransactionDTO) (*transaction.TransactionResp, error)
}

type OHLCUseCase struct {
	Redis    redis.RedisInt
	Producer kafka.Producer
}

func NewOHLCUseCase(rds redis.RedisInt, prd kafka.Producer) *OHLCUseCase {
	return &OHLCUseCase{
		Redis:    rds,
		Producer: prd,
	}
}

func (uc *OHLCUseCase) Calculate(data []*dto.TransactionDTO) error {

	msg, _ := json.Marshal(data)

	//publish to topic ohlc.transaction
	if err := uc.Producer.SendMessage("ohlc.transaction", msg, []byte(""), []sarama.RecordHeader{}); err != nil {
		return err
	}

	return nil
}

func (uc *OHLCUseCase) ConsumeAndCalculate(ctx context.Context, data []*dto.TransactionDTO) error {
	result := &dto.TransactionRespDTO{}

	var key string

	dataKey, _ := json.Marshal(data)
	key = string(dataKey)

	dataRedis, err := uc.Redis.GetData(ctx, key)
	if err != nil {
		log.Printf("Data Not Found in Redis. error: %v\n", err)
	}

	if dataRedis != "" {
		_ = json.Unmarshal([]byte(dataRedis), &result)

		log.Println("Data From Redis")
	} else {
		result = calculateTransaction(data)

		redisData, _ := json.Marshal(result)
		ttl := time.Duration(2) * time.Minute

		// set data to redis
		err := uc.Redis.SetData(ctx, key, redisData, ttl)
		if err != nil {
			log.Printf("unable to SET data. error: %v", err)
		}
	}

	fmt.Println("Result:")
	fmt.Println("Previous Price:", result.PreviousPrice)
	fmt.Println("Open Price:", result.OpenPrice)
	fmt.Println("Highest Price:", result.HighestPrice)
	fmt.Println("Lowest Price:", result.LowestPrice)
	fmt.Println("Close Price:", result.ClosePrice)
	fmt.Println("Volume:", result.Volume)
	fmt.Println("Value:", result.Value)
	fmt.Println("Average Price:", result.AveragePrice)

	return nil
}

func (uc *OHLCUseCase) CalculateFromRPC(ctx context.Context, data []*dto.TransactionDTO) (*transaction.TransactionResp, error) {

	res := &dto.TransactionRespDTO{}
	var key string

	dataKey, _ := json.Marshal(data)
	key = string(dataKey)

	dataRedis, err := uc.Redis.GetData(context.Background(), key)
	if err != nil {
		log.Printf("Data Not Found in Redis. error: %v\n", err)
	}

	if dataRedis != "" {
		result := &transaction.TransactionResp{}
		_ = json.Unmarshal([]byte(dataRedis), &result)

		log.Println("Data From Redis")
		return result, nil
	}

	res = calculateTransaction(data)
	result := &transaction.TransactionResp{
		PreviousPrice: res.PreviousPrice,
		OpenPrice:     res.OpenPrice,
		HighestPrice:  res.HighestPrice,
		LowestPrice:   res.LowestPrice,
		ClosePrice:    res.ClosePrice,
		Volume:        res.Volume,
		Value:         res.Value,
		AveragePrice:  res.AveragePrice,
	}
	redisData, _ := json.Marshal(result)
	ttl := time.Duration(2) * time.Minute

	// set data to redis
	err = uc.Redis.SetData(context.Background(), key, redisData, ttl)
	if err != nil {
		log.Printf("unable to SET data. error: %v", err)
	}

	return result, nil
}

func calculateTransaction(data []*dto.TransactionDTO) *dto.TransactionRespDTO {
	result := &dto.TransactionRespDTO{}
	for _, t := range data {
		if t.Quantity == 0 { // Previous Price

			result.PreviousPrice = t.Price

		} else if t.Type == "A" { // Open Price

			result.OpenPrice = t.Price

			if result.HighestPrice == 0 || t.Price > result.HighestPrice { // Highest Price
				result.HighestPrice = t.Price
			}
			if result.LowestPrice == 0 || t.Price < result.LowestPrice { // Lowest Price
				result.LowestPrice = t.Price
			}

		} else if t.Type == "E" || t.Type == "P" { // Close Price
			result.ClosePrice = t.Price
			if result.HighestPrice == 0 || t.Price > result.HighestPrice { // Highest Price
				result.HighestPrice = t.Price
			}
			if result.LowestPrice == 0 || t.Price < result.LowestPrice { // Lowest Price
				result.LowestPrice = t.Price
			}

			result.Volume += t.Quantity          // Volume
			result.Value += t.Quantity * t.Price // Value
		}
	}

	if result.Value != 0 && result.Volume != 0 {
		result.AveragePrice = result.Value / result.Volume // Average Price (rounded down)
	}

	return result
}
