package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducerInterface interface {
	SendService(key string, value []byte) error
}

type KafkaProducer struct {
	KafkaWriter *kafka.Writer
	Ctx         context.Context
}

// NewService adalah fungsi yang membuat instance baru dari Service yang terhubung dengan Kafka.
func NewProducerService(kafkaWriter *kafka.Writer, ctx context.Context) KafkaProducerInterface {
	return &KafkaProducer{KafkaWriter: kafkaWriter, Ctx: ctx}
}

// SendService mengirimkan pesan ke Kafka dengan key dan value yang diberikan.
func (s *KafkaProducer) SendService(key string, value []byte) error {

	// Buat pesan Kafka
	pesan := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	// Tulis pesan ke Kafka
	err := s.KafkaWriter.WriteMessages(s.Ctx, pesan)

	if err != nil {
		return fmt.Errorf("terjadi kesalahan saat menulis ke Kafka: %w", err)
	}

	return nil

}
