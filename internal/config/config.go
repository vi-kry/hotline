package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	ISSO       `yaml:"isso"`
	Token      `yaml:"token"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type ISSO struct {
	URL       string `yaml:"url" env-required:"true"`
	ClientID  string `yaml:"client_id" env-required:"true"`
	GrantType string `yaml:"grant_type" env-required:"true"`
}

type Token struct {
	Secret    string `yaml:"secret" env-default:"0000000-0000-0000-0000-000000000000"`
	ExpiresAt string `yaml:"expires_at" env-default:"10080"`
}

func MustLoad() *Config {
	//configPath := os.Getenv("CONFIG_PATH")
	configPath := "./configs/config.yaml"

	//if configPath == "" {
	//	log.Fatal("CONFIG_PATH is not set")
	//}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
