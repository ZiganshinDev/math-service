package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Notifier   `yaml:"notifier"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
	Mathmaker  `yaml:"mathmaker"`
}

type HTTPServer struct {
	Port        string        `yaml:"port" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	PostgresURI           string `yaml:"postgres_uri" env-required:"true"`
	RedisConnectionString string `yaml:"redis_connection_string" env-required:"true"`
	RabbitMQ              string `yaml:"amqp_url" env-required:"true"`
}

type Notifier struct {
	Queue string `yaml:"queue" env-required:"true"`
}

type Mathmaker struct {
	URL string `yaml:"url" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
