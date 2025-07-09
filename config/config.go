package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type BotConfig struct {
	APIKey string
	Owner int64
	Timeout int
}

func GetConfig() *BotConfig {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf := BotConfig{}
	var ok bool
	if conf.APIKey, ok=os.LookupEnv("APIKey"); !ok {
		log.Fatal("Cant read .env")
	}

	conf.Owner = 123456789 
	conf.Timeout = 60
	return &conf
}
