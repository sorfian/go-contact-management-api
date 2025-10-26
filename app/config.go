package app

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sorfian/go-contact-management-api/helper"
)

type Config struct {
	AppEnv   string
	AppPort  string
	Database DatabaseConfig
	LogLevel string
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

var AppConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if exists (for development)
	if err := godotenv.Load(); err != nil {
		// Load .env file for testing environment
		if err := godotenv.Load("../.env"); err != nil {
			log.Println("No .env file found, using environment variables")
		}
	}

	config := &Config{
		AppEnv:  helper.GetEnv("APP_ENV", "development"),
		AppPort: helper.GetEnv("APP_PORT", "3000"),
		Database: DatabaseConfig{
			Host:            helper.GetEnv("DB_HOST", "localhost"),
			Port:            helper.GetEnv("DB_PORT", "3306"),
			User:            helper.GetEnv("DB_USER", "root"),
			Password:        helper.GetEnv("DB_PASSWORD", ""),
			Name:            helper.GetEnv("DB_NAME", "go_todo_list"),
			MaxIdleConns:    helper.GetEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns:    helper.GetEnvAsInt("DB_MAX_OPEN_CONNS", 100),
			ConnMaxLifetime: helper.GetEnvAsDuration("DB_CONN_MAX_LIFETIME", 30*time.Minute),
			ConnMaxIdleTime: helper.GetEnvAsDuration("DB_CONN_MAX_IDLE_TIME", 10*time.Minute),
		},
		LogLevel: helper.GetEnv("LOG_LEVEL", "info"),
	}

	AppConfig = config
	return config
}

// GetDSN returns database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}
