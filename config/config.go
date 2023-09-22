package config

import (
	"fmt"
	"os"
)

var config Config

type Config struct {
	Database
}

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

func Instance() Config {
	return config
}

func InitDbConfig() *Config {
	var ok bool
	user, ok := os.LookupEnv("POSTGRES_HOST")
	password, ok := os.LookupEnv("POSTGRES_USER")
	host, ok := os.LookupEnv("POSTGRES_PASSWORD")
	name, ok := os.LookupEnv("POSTGRES_DB")
	port, ok := os.LookupEnv("POSTGRES_PORT")

	if !ok {
		fmt.Println("The database environment variables does not exist or are not set.")
		return &config
	}

	config.Database = Database{
		Host:     user,
		User:     password,
		Password: host,
		Name:     name,
		Port:     port,
	}

	return &config
}
