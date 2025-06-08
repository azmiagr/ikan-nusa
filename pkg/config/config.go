package config

import (
	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	_ = godotenv.Load("../../.env")

	return nil
}
