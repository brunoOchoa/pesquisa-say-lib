package config

import (
	"log"
	"os"
)

type Config struct {
	AccessToken   string
	PhoneNumberID string
}

func LoadConfig() *Config {

	accessToken := os.Getenv("WA_ACCESS_TOKEN")
	phoneID := os.Getenv("WA_PHONE_ID")

	if accessToken == "" || phoneID == "" {
		log.Fatal("Variáveis de ambiente WA_ACCESS_TOKEN ou WA_PHONE_ID não definidas.")
	}

	return &Config{
		AccessToken:   accessToken,
		PhoneNumberID: phoneID,
	}

}
