package ohlc

import (
	"context"
	"encoding/json"
	"errors"
	mockKafka "solution_ch_part1/mocks/infra/kafka"
	mockRedis "solution_ch_part1/mocks/infra/redis"
	dto "solution_ch_part1/src/app/dto/transaction"
	"time"

	"testing"

	transaction "solution_ch_part1/src/app/proto"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockArticleUseCase struct {
	mock.Mock
}

type OHLCUseCaseList struct {
	suite.Suite
	useCase          OHLCUCInterface
	mockRedis        *mockRedis.MockRedisUC
	dtoCreate        []*dto.TransactionDTO
	dtoCreateQ0      []*dto.TransactionDTO
	dtoCreateTypeE   []*dto.TransactionDTO
	dtoResp          *dto.TransactionRespDTO
	dtoRespQ0        *dto.TransactionRespDTO
	dtoRespQTypeE    *dto.TransactionRespDTO
	mockKafka        *mockKafka.MockKafka
	dtoRespGRPC      *transaction.TransactionResp
	dtoRespGRPCQ0    *transaction.TransactionResp
	dtoRespGRPCTypeE *transaction.TransactionResp
}

func (suite *OHLCUseCaseList) SetupTest() {
	suite.mockRedis = new(mockRedis.MockRedisUC)
	suite.mockKafka = new(mockKafka.MockKafka)
	suite.useCase = NewOHLCUseCase(suite.mockRedis, suite.mockKafka)

	suite.dtoCreate = []*dto.TransactionDTO{
		&dto.TransactionDTO{
			Number:    "test",
			Type:      "A",
			Stock:     "test",
			OrderBook: "test",
			Quantity:  500,
			Price:     500,
		},
	}

	suite.dtoCreateQ0 = []*dto.TransactionDTO{
		&dto.TransactionDTO{
			Number:    "test",
			Type:      "A",
			Stock:     "test",
			OrderBook: "test",
			Quantity:  0,
			Price:     500,
		},
	}

	suite.dtoResp = &dto.TransactionRespDTO{
		PreviousPrice: 0,
		OpenPrice:     500,
		HighestPrice:  500,
		LowestPrice:   500,
		ClosePrice:    0,
		Volume:        0,
		Value:         0,
		AveragePrice:  0,
	}

	suite.dtoCreateTypeE = []*dto.TransactionDTO{
		&dto.TransactionDTO{
			Number:    "test",
			Type:      "E",
			Stock:     "test",
			OrderBook: "test",
			Quantity:  200,
			Price:     500,
		},
	}

	suite.dtoRespQ0 = &dto.TransactionRespDTO{
		PreviousPrice: 500,
		OpenPrice:     0,
		HighestPrice:  0,
		LowestPrice:   0,
		ClosePrice:    0,
		Volume:        0,
		Value:         0,
		AveragePrice:  0,
	}

	suite.dtoRespQTypeE = &dto.TransactionRespDTO{
		PreviousPrice: 0,
		OpenPrice:     0,
		HighestPrice:  500,
		LowestPrice:   500,
		ClosePrice:    500,
		Volume:        200,
		Value:         100000,
		AveragePrice:  500,
	}

	suite.dtoRespGRPC = &transaction.TransactionResp{
		PreviousPrice: 0,
		OpenPrice:     500,
		HighestPrice:  500,
		LowestPrice:   500,
		ClosePrice:    0,
		Volume:        0,
		Value:         0,
		AveragePrice:  0,
	}

	suite.dtoRespGRPCQ0 = &transaction.TransactionResp{
		PreviousPrice: 500,
		OpenPrice:     0,
		HighestPrice:  0,
		LowestPrice:   0,
		ClosePrice:    0,
		Volume:        0,
		Value:         0,
		AveragePrice:  0,
	}

	suite.dtoRespGRPCTypeE = &transaction.TransactionResp{
		PreviousPrice: 0,
		OpenPrice:     0,
		HighestPrice:  500,
		LowestPrice:   500,
		ClosePrice:    500,
		Volume:        200,
		Value:         100000,
		AveragePrice:  500,
	}

}

func (u *OHLCUseCaseList) TestCalculateSuccess() {
	marshal, _ := json.Marshal(u.dtoCreate)

	u.mockKafka.Mock.On("SendMessage", "ohlc.transaction", marshal, []byte{}, []sarama.RecordHeader{}).Return(nil)
	err := u.useCase.Calculate(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestCalculateFail() {
	marshal, _ := json.Marshal(u.dtoCreate)

	u.mockKafka.Mock.On("SendMessage", "ohlc.transaction", marshal, []byte{}, []sarama.RecordHeader{}).Return(errors.New(mock.Anything))
	err := u.useCase.Calculate(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *OHLCUseCaseList) TestConsumeCalculateSuccessFromRedis() {
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreate)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", ctx, key).Return(string(dataresp), nil)
	err := u.useCase.ConsumeAndCalculate(ctx, u.dtoCreate)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestConsumeCalculateSuccessNotFromRedis() {
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreate)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoResp)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	err := u.useCase.ConsumeAndCalculate(ctx, u.dtoCreate)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestConsumeCalculateSuccessNotFromRedisQuantity0() {
	// for quantity 0
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateQ0)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespQ0)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	err := u.useCase.ConsumeAndCalculate(ctx, u.dtoCreateQ0)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestConsumeCalculateSuccessNotFromRedisTypeE() {
	// for type E or P
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateTypeE)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespQTypeE)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	err := u.useCase.ConsumeAndCalculate(ctx, u.dtoCreateTypeE)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestConsumeCalculateErrorSetDataRedis() {
	// for type E or P
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateTypeE)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespQTypeE)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(errors.New(mock.Anything))
	err := u.useCase.ConsumeAndCalculate(ctx, u.dtoCreateTypeE)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestCalculateFromRPCFromRedis() {
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreate)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", ctx, key).Return(string(dataresp), nil)
	_, err := u.useCase.CalculateFromRPC(ctx, u.dtoCreate)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestCalculateFromRPCSuccessNotFromRedis() {
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreate)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespGRPC)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	_, err := u.useCase.CalculateFromRPC(ctx, u.dtoCreate)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestCalculateFromRPCSuccessNotFromRedisQuantity0() {
	// for quantity 0
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateQ0)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespGRPCQ0)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	_, err := u.useCase.CalculateFromRPC(ctx, u.dtoCreateQ0)
	u.Equal(nil, err)
}
func (u *OHLCUseCaseList) TestCalculateFromRPCSuccessNotFromRedisTypeE() {
	// for type E or P
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateTypeE)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespGRPCTypeE)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(nil)
	_, err := u.useCase.CalculateFromRPC(ctx, u.dtoCreateTypeE)
	u.Equal(nil, err)
}

func (u *OHLCUseCaseList) TestCalculateFromRPCErrorSetDataRedis() {
	// for type E or P
	var key string
	ctx := context.Background()
	dataKey, _ := json.Marshal(u.dtoCreateTypeE)
	key = string(dataKey)

	dataresp, _ := json.Marshal(u.dtoRespGRPCTypeE)
	ttl := time.Duration(2) * time.Minute
	u.mockRedis.Mock.On("GetData", ctx, key).Return("", errors.New(mock.Anything))
	u.mockRedis.Mock.On("SetData", ctx, key, dataresp, ttl).Return(errors.New(mock.Anything))
	_, err := u.useCase.CalculateFromRPC(ctx, u.dtoCreateTypeE)
	u.Equal(nil, err)
}
func TestUsecase(t *testing.T) {
	suite.Run(t, new(OHLCUseCaseList))
}
