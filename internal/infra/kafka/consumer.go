package kafka

import (
	"fmt"
	"time"

	goKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/uzimbra/orch-process-transactions/internal/config"
	"go.uber.org/zap"
)

var kafkaConsumer *goKafka.Consumer

func StartKafkaConsumer() (*goKafka.Consumer, error) {
	kafkaEnv := config.GetKafkaEnv()
	c, err := goKafka.NewConsumer(&goKafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", kafkaEnv.Host, kafkaEnv.Port),
		"group.id":          "github.com/uzimbra",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{config.GetKafkaTopic().ProcessTransactions}, nil)
	kafkaConsumer = c

	go startKafkaConsumer()

	return kafkaConsumer, nil
}

func startKafkaConsumer() {
	for {
		msg, err := kafkaConsumer.ReadMessage(time.Second)
		if err == nil {
			zap.L().Sugar().Infow("Message received to be processed", "partition", msg.TopicPartition, "message", string(msg.Value))

			processReceivedTransactionMessage(string(msg.Value))
		} else if !err.(goKafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			return
		}
	}
}
