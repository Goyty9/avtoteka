package config

import "os"

//конфигурация(env, настройки бд и логгера)

// for save app config
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// download config from env vars
func Load() *Config {
	return &Config{
		DBHost:     getEnv("DB_Host", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_User", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "-----"),
		DBName:     getEnv("DB_NAME", "avtoteka"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// additional for get env vars
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
