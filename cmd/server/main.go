package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	config := loadConfig()

	_ = config.GetString("grpc_address")
	_ = config.GetString("db_path")

}

func loadConfig() *viper.Viper {
	config := viper.New()

	config.SetDefault("grpc_port", "50051")
	config.SetDefault("db_path", "data/database.db")

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")

	config.AutomaticEnv()

	config.BindEnv("grpc_port", "GRPC_PORT")
	config.BindEnv("db_path", "DB_PATH")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	log.Printf("Loading configuration: %s", config.ConfigFileUsed())
	return config
}
