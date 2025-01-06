package main

import (
	"github.com/pyeremenko/klausapp-scoring/pkg/db"
	"github.com/spf13/viper"
	"log"
)

func main() {
	config := loadConfig()

	_ = config.GetString("grpc_address")
	dbPath := config.GetString("db_path")

	dbConn, err := db.InitSQLite(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbConn.Close()
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
