package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Env struct {
	AppPort string `env:"APP_PORT" envDefault:"8080"`
	DBHost  string `env:"DB_HOST" envDefault:"localhost"`
	DBPort  string `env:"DB_PORT" envDefault:"5432"`
	DBUser  string `env:"DB_USER" envDefault:"postgres"`
	DBPass  string `env:"DB_PASS" envDefault:"postgres"`
	DBName  string `env:"DB_NAME" envDefault:"hmdtif_todo"`
}

func NewEnv() *Env {
	e := Env{}
	if err := env.Parse(&e); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}
	return &e
}
