package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	Env          string      `env:"ENV"`
	DbUrl        string      `env:"DB_URL"`
	LoggerLevel  *slog.Level `yaml:"loggerLevel"`
	WorkersCount int         `env:"WORKERS_COUNT"`
}

func MustLoad() *Config {
	godotenv.Load()

	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(configPath); err != nil {
		panic(err)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic(err.Error())
	}
	return cfg
}

func fetchConfigPath() (res string) {
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "config/config_local.yml"
	}
	return
}
