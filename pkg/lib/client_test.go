package lib

import (
	"testing"

	"github.com/brunoOchoa/pesquisa-say-lib/config"
)

func TestNewClient(t *testing.T) {

	cfg := &config.Config{
		AccessToken:   "teste_token",
		PhoneNumberID: "123456789",
	}

	client := NewClient(cfg)

	if client.AccessToken != cfg.AccessToken {
		t.Errorf("AccessToken incorreto: esperado '%s', recebido '%s'", cfg.AccessToken, client.AccessToken)
	}

	if client.PhoneNumberID != cfg.PhoneNumberID {
		t.Errorf("PhoneNumberID incorreto: esperado '%s', recebido '%s'", cfg.PhoneNumberID, client.PhoneNumberID)
	}

	if client.HttpClient == nil {
		t.Error("HttpClient não deve ser nil")
	}
}
