package main

import (
	"github.com/uzimbra/orch-process-transactions/cmd/services"
	"github.com/uzimbra/orch-process-transactions/internal/config"
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal(err.Error())
	}
	zap.ReplaceGlobals(logger)
}

func main() {
	if err := config.Load(); err != nil {
		zap.L().Fatal(err.Error())
	}

	c := services.StartKafkaConsumer()
	defer c.Close()
	p := services.StartKafkaProducer()
	defer p.Close()

	services.StartApi()
}
