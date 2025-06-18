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

	accessToken := os.Getenv("WHATSAPP_TOKEN")
	phoneID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")

	if accessToken == "" || phoneID == "" {
		log.Fatal("Variáveis de ambiente WHATSAPP_TOKEN ou WHATSAPP_PHONE_NUMBER_ID não definidas.")
	}

	return &Config{
		AccessToken:   accessToken,
		PhoneNumberID: phoneID,
	}

}
