package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db *DbConfig
	Gql *GqlConfig
}
 
type DbConfig struct {
	Driver 		string
	Host 		string
	Port 		string
	DbName 		string
	User     	string
	Password  	string
}

type GqlConfig struct {
	Port string
}
 
func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
    	log.Fatal("Error loading .env file")
  	}

	return &Config{
		Db: &DbConfig{
			Driver: 	os.Getenv("DB_DRIVER"),
			Host: 		os.Getenv("DB_HOST"),
			Port: 		os.Getenv("DB_PORT"),
			DbName: 	os.Getenv("DB_NAME"),
			User:     	os.Getenv("DB_USER"),
			Password:  	os.Getenv("DB_PASSWORD"),
		},
		Gql: &GqlConfig{
			Port: os.Getenv("GQL_PORT"),
		},
	}
}
