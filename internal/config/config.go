package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `env:"ENV" yaml:"env" env-dafault:"local" env-required:"true"`
	StoragePath string `env:"STORAGE_PATH" yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"https_server"`
}

type HTTPServer struct {
	Address     string        `env:"HTTP_SERVER_ADDRESS" yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT" yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDDLE_TIMEOUT" yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Check is file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	return cfg
}
