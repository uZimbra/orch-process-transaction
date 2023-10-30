package services

import (
	goKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/uzimbra/orch-process-transactions/internal/infra/kafka"

	"go.uber.org/zap"
)

func StartKafkaProducer() *goKafka.Producer {
	zap.L().Info("Starting kafka message producer!")
	kafkaProducer, err := kafka.StartKafkaProducer()
	if err != nil {
		zap.L().Sugar().Fatalw("Failed to start kafka consummer", "error", err)
	}
	return kafkaProducer
}

func StartKafkaConsumer() *goKafka.Consumer {
	zap.L().Info("Starting kafka message consumer!")
	kafkaConsumer, err := kafka.StartKafkaConsumer()
	if err != nil {
		zap.L().Sugar().Fatalw("Failed to start kafka producer", "error", err)
	}
	return kafkaConsumer
}
