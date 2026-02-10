package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Database string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		MongoURI: os.Getenv("MONGO_URI"),
		Database: os.Getenv("MONGO_DB"),
	}

	if cfg.MongoURI == "" || cfg.Database == "" {
		log.Fatal("Variables de entorno Mongo no configuradas")
	}

	return cfg
}
