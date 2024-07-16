package configs

// здесь будут обрабатываться конфиг файлы

import (
	"os"

	"github.com/joho/godotenv"
)
// общая структура конфига
type Config struct {
	Server Server
}

type Server struct {
	Host string `env:"SERVER_HOST" default:"127.0.0.1"`
	Port string `env:"SERVER_PORT" default:"8080"`
}

func ReadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}
