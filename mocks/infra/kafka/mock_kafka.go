package mock_kafka

import (
	"solution_ch_part1/src/infra/kafka"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/mock"
)

type MockKafka struct {
	mock.Mock
}

func NewMocKafka() *MockKafka {
	return &MockKafka{}
}

// implement interface
var _ kafka.Producer = &MockKafka{}

func (m *MockKafka) SendMessage(topic string, value []byte, key []byte, headers []sarama.RecordHeader) error {
	args := m.Called(topic, value, key, headers)
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockKafka) Close() {

}
