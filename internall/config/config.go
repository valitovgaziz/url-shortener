package config

import "time"

type Config struct {
	Env     string `yaml:"env" env:"ENV" env-default:"local"`
	Storage string `yaml:"storage_path" env:"STORAGE" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}
