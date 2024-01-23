package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)



type Config struct {
	PostgresHost 	 string
	PostgresPort 	 string
	PostgresUser   	 string
	PostgresPassword string
	PostgressDB 	 string
}

func Load()  {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error: ", err)
	}


	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "host"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "1234"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "user"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "mypassword"))
	cfg.PostgressDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "DB"))
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}