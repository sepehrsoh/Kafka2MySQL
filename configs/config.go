package configs

import (
	"log"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type KafkaConfig struct {
	Brokers []string
}

type DatabaseConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     int
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct %s", err)
	}

	return config
}
