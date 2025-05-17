package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	//Load .env files
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, relying on OS ENV vars")
	}

	viper.AutomaticEnv() // use env vars from system or .env
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JWT_SECRET", "supersecret")
	viper.SetDefault("JWT_EXPIRE_HOURS", 24)
}

// Helper to get a config value
func Get(key string) string {
	return viper.GetString(key)
}

// Or as int
func GetInt(key string) int {
	return viper.GetInt(key)
}
