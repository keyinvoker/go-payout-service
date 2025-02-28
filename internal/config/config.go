package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found. Falling back to system environment variables.")
		}

		cfg = &Config{
			DBUser:     getEnv("POSTGRES_DB_USER", ""),
			DBPassword: getEnv("POSTGRES_DB_PASSWORD", ""),
			DBHost:     getEnv("POSTGRES_DB_HOST", ""),
			DBPort:     getEnv("POSTGRES_DB_PORT", ""),
			DBName:     getEnv("POSTGRES_DB_NAME", "payout"),
			ServerPort: getEnv("SERVER_PORT", "8888"),
		}
	})
	return cfg
}

func GetConfig() *Config {
	if cfg == nil {
		return LoadConfig()
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
