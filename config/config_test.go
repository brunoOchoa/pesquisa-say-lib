package config

import (
	"os"
	"testing"
)

func TestLoadConfig_Valid(t *testing.T) {
	os.Setenv("WA_ACCESS_TOKEN", "teste_token")
	os.Setenv("WA_PHONE_ID", "teste_id")

	cfg := LoadConfig()

	if cfg.AccessToken != "teste_token" {
		t.Errorf("Esperado AccessToken 'teste_token', recebido: %s", cfg.AccessToken)
	}
	if cfg.PhoneNumberID != "teste_id" {
		t.Errorf("Esperado PhoneNumberID 'teste_id', recebido: %s", cfg.PhoneNumberID)
	}
}
