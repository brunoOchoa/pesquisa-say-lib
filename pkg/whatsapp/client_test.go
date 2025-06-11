package whatsapp

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	accessToken := "teste_token"
	phoneID := "123456789"

	client := NewClient(accessToken, phoneID)

	if client.AccessToken != accessToken {
		t.Errorf("AccessToken incorreto: esperado '%s', recebido '%s'", accessToken, client.AccessToken)
	}

	if client.PhoneNumberID != phoneID {
		t.Errorf("PhoneNumberID incorreto: esperado '%s', recebido '%s'", phoneID, client.PhoneNumberID)
	}

	if client.HttpClient == nil {
		t.Error("HttpClient não deve ser nil")
	}
}
