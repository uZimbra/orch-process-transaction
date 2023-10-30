package kafka

import (
	"fmt"

	goKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/uzimbra/orch-process-transactions/internal/config"
	"go.uber.org/zap"
)

var kafkaProducer *goKafka.Producer
var deliveryChannel = make(chan goKafka.Event)

func StartKafkaProducer() (*goKafka.Producer ,error) {
	kafkaEnv := config.GetKafkaEnv()
	p, err := goKafka.NewProducer(&goKafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", kafkaEnv.Host, kafkaEnv.Port),
	})
	if err != nil {
		return nil, err
	}

	kafkaProducer = p

	go startMessageDelivery()

	return kafkaProducer, nil
}

func startMessageDelivery() {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *goKafka.Message:
			if ev.TopicPartition.Error != nil {
				zap.L().Sugar().Infow("kafka delivery failed", "topic", ev.TopicPartition)
			} else {
				zap.L().Sugar().Infow("kafka delivery message to", "topic", ev.TopicPartition)
			}
		}
	}
}
