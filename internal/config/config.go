package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"hotline/internal/models"
	"log"
	"os"
)

func MustLoad() *models.Config {
	//configPath := os.Getenv("CONFIG_PATH")
	configPath := "./configs/config.yaml"

	//if configPath == "" {
	//	log.Fatal("CONFIG_PATH is not set")
	//}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg models.Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
