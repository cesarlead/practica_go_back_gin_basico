package config

import (
	"os"
	"strings"
)

// Config agrupa la configuración de la aplicación
type Config struct {
	Port        string
	CORSOrigins []string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
}

// Load lee variables de entorno con valores por defecto
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8185"),
		CORSOrigins: parseList(getEnv("CORS_ORIGINS", "*")),
		DBUser:      getEnv("DB_USER", "go_test_user"),
		DBPassword:  getEnv("DB_PASSWORD", "go_test_pass"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5435"),
		DBName:      getEnv("DB_NAME", "app_test"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseList(s string) []string {
	if s == "*" {
		return []string{"*"}
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
