package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GithubToken string
}

var config *Config

func GetConfig() *Config {
	return config
}

func Init(location string) error {
	err := godotenv.Load(location)
	if err != nil {
		log.Println("[Init] error while load environment variable")
		return err
	}

	config = &Config{
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}

	if config.GithubToken == "" {
		log.Println("[Init] GithubToken must not empty")
		return errors.New("GithubToken Empty")
	}

	return nil
}
