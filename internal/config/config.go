package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func ConfigLoad() (*Config, error) {
	path, err := fetchConfigPath()

	if err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	return &cfg, nil
}

func fetchConfigPath() (string, error) {
	var res string

	flag.StringVar(&res, "config", "", "path to cfg file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	if res == "" {
		return "", fmt.Errorf("cannot fetch config file path")
	}

	return res, nil
}
