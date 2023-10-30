package kafka

import (
	"encoding/json"

	goKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/uzimbra/orch-process-transactions/internal/config"
	"github.com/uzimbra/orch-process-transactions/internal/entity"
	"github.com/uzimbra/orch-process-transactions/internal/infra/usecase"
	"go.uber.org/zap"
)

func processReceivedTransactionMessage(t string) {
	transaction := &entity.Transaction{}

	if err := json.Unmarshal([]byte(t), transaction); err != nil {
		zap.L().Sugar().Errorw("Unable to unmarshal transaction", "transaction", t, "error", err)
		return
	}

	processedTransaction, err := usecase.ProcessTransactionUseCase(transaction)
	if err != nil {
		zap.L().Sugar().Errorw("An error occured when trying to process transaction", "id", transaction.Id, "error", err)
		return
	}

	msg, err := json.Marshal(processedTransaction)
	if err != nil {
		zap.L().Sugar().Errorw("Unable to marshal transaction", "id", transaction.Id, "error", err)
		return
	}

	topic := config.GetKafkaTopic().UpdateTransactionStatus
	if err := ProduceMessage(topic, msg); err != nil {
		zap.L().Sugar().Errorw("Fail to produce transaction message", "id", transaction.Id, "error", err)
		return
	}
}

func ProduceMessage(topic string, message []byte) error {
	msg := &goKafka.Message{
		TopicPartition: goKafka.TopicPartition{Topic: &topic, Partition: goKafka.PartitionAny},
		Value:          message,
	}

	return kafkaProducer.Produce(msg, deliveryChannel)
}
