package config

import (
	"log"
	"os"
)

type Config struct {
	AccessToken   string
	PhoneNumberID string
	ApiVersion    string
}

func LoadConfig() *Config {

	accessToken := os.Getenv("WHATSAPP_TOKEN")
	phoneID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	apiVersion := os.Getenv("WHATSAPP_API_VERSION")

	if accessToken == "" || phoneID == "" || apiVersion == "" {
		log.Fatal("Variáveis de ambiente WHATSAPP_TOKEN, WHATSAPP_PHONE_NUMBER_ID ou WHATSAPP_API_VERSION não definidas.")
	}

	return &Config{
		AccessToken:   accessToken,
		PhoneNumberID: phoneID,
	}

}
