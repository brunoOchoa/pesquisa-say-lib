package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := LoadConfig()

	if cfg.AccessToken != "teste_token" {
		t.Errorf("Esperado AccessToken 'teste_token', recebido: %s", cfg.AccessToken)
	}
	if cfg.PhoneNumberID != "teste_id" {
		t.Errorf("Esperado PhoneNumberID 'teste_id', recebido: %s", cfg.PhoneNumberID)
	}
}
