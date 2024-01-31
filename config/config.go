package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type FileConfig struct {
	FilePath string
}

type SessionConfig struct {
	SessionKey string
	MaxAge     int
}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
	SessionConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return errors.New("error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}

	c.FileConfig = FileConfig{
		FilePath: os.Getenv("LOG_FILE"),
	}

	maxAge, _ := strconv.Atoi(os.Getenv("SESSION_MAX_AGE"))

	c.SessionConfig = SessionConfig{
		SessionKey: os.Getenv("SESSION_KEY"),
		MaxAge:     maxAge,
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}
