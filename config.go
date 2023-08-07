package goceng

import (
	"github.com/ikhsan892/goceng/config"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DB config.Database
}

func newConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		DB: config.LoadDatabaseConfiguration(),
	}
}
