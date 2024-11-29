package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"simple-api/internal/utils"
	"sync"
)

var AppValidator = utils.NewValidator()

type Config struct {
	AccessKey string
	SecretKey string
}

var (
	once     sync.Once
	instance *Config
)

func Inst() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found, loading from OS environment variables.")
		}

		instance = &Config{
			AccessKey: getEnv("ACCESS_KEY", "Aboba"),
			SecretKey: getEnv("JWT_SECRET", "Asus"),
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
