package config

import "os"

type Config struct {
	ServerPort  string
	MongoURI    string
	MongoDBName string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		MongoURI:    getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBName: getEnv("MONGODB_NAME", "crud-go"),
		JWTSecret:   getEnv("JWT_SECRET", "change-this-secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
