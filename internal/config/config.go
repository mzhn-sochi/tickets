package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	App struct {
		Host string `env:"APP_HOST" env-default:"0.0.0.0"`
		Port int    `env:"APP_PORT" env-default:"8080"`
	}

	DB struct {
		User string `env:"DB_USER" env-default:"postgres"`
		Pass string `env:"DB_PASS" env-default:"postgres"`
		Host string `env:"DB_HOST" env-default:"localhost"`
		Port int    `env:"DB_PORT" env-default:"5434"`
		Name string `env:"DB_NAME" env-default:"tickets"`
	}

	Nats struct {
		URL    string `env:"NATS_URL" env-required:"true"`
		Queues struct {
			OCR        string `env:"NATS_QUEUE_OCR" env-required:"true"`
			Errors     string `env:"NATS_QUEUE_ERRORS" env-required:"true"`
			Validation string `env:"NATS_QUEUE_VALIDATION" env-required:"true"`
			Status     string `env:"NATS_QUEUE_STATUS" env-required:"true"`
			Overprice  string `env:"NATS_QUEUE_OVERPRICE" env-required:"true"`
		}
	}
}

func New() *Config {
	config := &Config{}

	if err := cleanenv.ReadEnv(config); err != nil {
		header := "TICKET SERVICE ENVs"
		f := cleanenv.FUsage(os.Stdout, config, &header)
		f()
		panic(err)
	}

	return config
}
