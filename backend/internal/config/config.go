package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	ServerPort         string
	JWTSecret          string
	JWTExpiry          string
	RefreshTokenExpiry string
	QueryTimeout       time.Duration
	DBMaxRetries       int
	DBRetryDelay       time.Duration
}

func Load() *Config {
	// Load .env file if it exists (try multiple paths)
	_ = godotenv.Load()
	_ = godotenv.Load("../../.env")

	// Parse query timeout (default: 10 seconds)
	queryTimeoutStr := getEnv("QUERY_TIMEOUT", "10s")
	queryTimeout, err := time.ParseDuration(queryTimeoutStr)
	if err != nil || queryTimeout <= 0 {
		queryTimeout = 10 * time.Second
	}

	// Parse retry delay (default: 2 seconds)
	retryDelayStr := getEnv("DB_RETRY_DELAY", "2s")
	retryDelay, err := time.ParseDuration(retryDelayStr)
	if err != nil || retryDelay <= 0 {
		retryDelay = 2 * time.Second
	}

	// Parse max retries (default: 5)
	maxRetries := 5
	if retriesStr := getEnv("DB_MAX_RETRIES", ""); retriesStr != "" {
		if retries, err := parseInt(retriesStr); err == nil && retries > 0 {
			maxRetries = retries
		}
	}

	jwtSecret := getEnv("JWT_SECRET", "")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}

	return &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "app"),
		DBSSLMode:          getEnv("DB_SSLMODE", "disable"),
		ServerPort:         getEnv("SERVER_PORT", "8080"),
		JWTSecret:          jwtSecret,
		JWTExpiry:          getEnv("JWT_EXPIRY", "15m"),
		RefreshTokenExpiry: getEnv("REFRESH_TOKEN_EXPIRY", "168h"),
		QueryTimeout:       queryTimeout,
		DBMaxRetries:       maxRetries,
		DBRetryDelay:       retryDelay,
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}
