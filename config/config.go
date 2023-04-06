package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string
	ServiceName string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int

	LogLevel string
	HttpPort string
	HttpHost string

	NotificationService string

	KafkaUrl string
}

func Load() Config {
	envFileName := cast.ToString(getOrReturnDefault("ENV_FILE_PATH", "./app/.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "info"))
	config.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "invan_auth_service"))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "invan_auth_service"))

	config.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	config.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", "6379"))
	config.RedisDB = cast.ToInt(getOrReturnDefault("REDIS_DB", 0))
	config.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))

	config.HttpHost = cast.ToString(getOrReturnDefault("LISTEN_HOST", "localhost"))
	config.HttpPort = cast.ToString(getOrReturnDefault("GRPC_PORT", ":8005"))

	config.KafkaUrl = cast.ToString(getOrReturnDefault("KAFKA_URL", "localhost:9092"))

	config.NotificationService = fmt.Sprintf("%s%s", cast.ToString(getOrReturnDefault("NOTIFICATION_SERVICE_HOST", "localhost")), cast.ToString(getOrReturnDefault("NOTIFICATION_GRPC_PORT", ":80")))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
