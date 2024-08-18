package consumers

import (
	"context"
	"encoding/json"
	"fmt"

	dto "solution_ch_part1/src/app/dto/transaction"
	usecase "solution_ch_part1/src/app/usecase/ohlc"
	"solution_ch_part1/src/infra/kafka"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type ohlcMessageConsumer struct {
	usecase usecase.OHLCUCInterface
}

type ConsumerInteraface interface {
	Consume(msg *sarama.ConsumerMessage) error
}

func NewOHLCMessageConsumer(u usecase.OHLCUCInterface) kafka.ClientConsumer {
	return &ohlcMessageConsumer{
		usecase: u,
	}
}

func (c *ohlcMessageConsumer) Consume(msg *sarama.ConsumerMessage) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover %+v", r)
		}
	}()
	ctx := context.Background()
	var data []*dto.TransactionDTO
	logrus.Infof("New message consumed from Partition %d", msg.Partition)

	err := json.Unmarshal(msg.Value, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.usecase.ConsumeAndCalculate(ctx, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
