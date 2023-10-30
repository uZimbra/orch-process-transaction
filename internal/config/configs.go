package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *configurations

type configurations struct {
	SERVER       ServerConfig
	KAFKA        KafkaConfig
	KAFKA_TOPICS KafkaTopics
}

type ServerConfig struct {
	Port string
}

type KafkaConfig struct {
	Host string
	Port string
}

type KafkaTopics struct {
	ProcessTransactions     string
	UpdateTransactionStatus string
}

func Load() error {
	profile, isProfileSet := os.LookupEnv("ENV")

	if (isProfileSet && profile == "local") || (isProfileSet && profile == "development") || !isProfileSet {
		if !isProfileSet {
			return loadLocalEnv("local")
		}
		return loadLocalEnv(profile)
	}

	zap.L().Sugar().Infow("loading custom environment", "profile", profile)

	config = new(configurations)

	config.SERVER = ServerConfig{
		Port: os.Getenv("PORT"),
	}

	config.KAFKA = KafkaConfig{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}

	config.KAFKA_TOPICS = KafkaTopics{
		ProcessTransactions:     os.Getenv("KAFKA_TOPIC_PROCESS_TRANSACTION"),
		UpdateTransactionStatus: os.Getenv("KAFKA_TOPIC_UPDATE_TRANSACTION_STATUS"),
	}

	return nil
}

func loadLocalEnv(profile string) error {
	zap.L().Sugar().Infow("loading local environment", "profile", profile)

	viper.SetConfigName(fmt.Sprintf("profile-%s", profile))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	config = new(configurations)

	config.SERVER = ServerConfig{
		Port: viper.GetString("server.port"),
	}

	config.KAFKA = KafkaConfig{
		Host: viper.GetString("kafka.host"),
		Port: viper.GetString("kafka.port"),
	}

	config.KAFKA_TOPICS = KafkaTopics{
		ProcessTransactions:     viper.GetString("kafka.process-transaction-topic"),
		UpdateTransactionStatus: viper.GetString("kafka.update-transaction-status-topic"),
	}

	return nil
}

func GetServerEnv() ServerConfig {
	return config.SERVER
}

func GetKafkaEnv() KafkaConfig {
	return config.KAFKA
}

func GetKafkaTopic() KafkaTopics {
	return config.KAFKA_TOPICS
}
