package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	AppPort string

	MongoURI    string
	MongoDBName string

	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioUseSSL    bool
	MinioBucket    string
	MinioPublicURL string

	JWTSecret      string
	JWTExpireHours int
}

var Env EnvConfig

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
	Env = EnvConfig{
		AppPort: getEnv("APP_PORT", "8080"),

		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName: getEnv("MONGO_DB_NAME", "movies"),

		MinioEndpoint:  getEnv("MINIO_ENDPOINT", "http://localhost:9000"),
		MinioAccessKey: getEnv("MINIO_ACCESS_KEY", ""),
		MinioSecretKey: getEnv("MINIO_SECRET_KEY", ""),
		MinioBucket:    getEnv("MINIO_BUCKET", "movies"),
		MinioUseSSL:    getEnvBool("MINIO_USE_SSL", false),
		MinioPublicURL: getEnv("MINIO_PUBLIC_URL", ""),

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpireHours: getEnvInt("JWT_EXPIRE_HOURS", 24),
	}

	if Env.MinioAccessKey == "" || Env.MinioSecretKey == "" {
		log.Fatal("Minio access key or secret key required")
	}
	if Env.JWTSecret == "" {
		log.Fatal("JWT secret key required")
	}
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist && value != "" {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	value, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return value == "true" || value == "1"
}

func getEnvInt(key string, fallback int) int {
	value, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return i
}
