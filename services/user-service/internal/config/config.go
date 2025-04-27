package config

import (
	"errors"
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	Port     string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, errors.New("DB_HOST environment variable is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432" // Default PostgreSQL port
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.New("DB_USER environment variable is required")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		return nil, errors.New("DB_PASSWORD environment variable is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, errors.New("DB_NAME environment variable is required")
	}

	return &Config{
		Port:     port,
		DBHost:   dbHost,
		DBPort:   dbPort,
		DBUser:   dbUser,
		DBPass:   dbPass,
		DBName:   dbName,
	}, nil
}

// DatabaseURL returns the database connection URL
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}