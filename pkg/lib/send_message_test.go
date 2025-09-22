package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brunoOchoa/pesquisa-say-lib/pkg/model"
)

type RedirectRoundTripper struct {
	original http.RoundTripper
	mockHost string
}

func (r *RedirectRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.Host, "graph.facebook.com") {
		req.URL.Scheme = "http"
		req.URL.Host = r.mockHost
	}
	return r.original.RoundTrip(req)
}

func TestSendTextMessage(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expectedPath := "/v22.0/test_id/messages"
		if r.URL.Path != expectedPath {
			t.Errorf("Caminho incorreto. Esperado %s, recebido %s", expectedPath, r.URL.Path)
		}

		if r.Method != http.MethodPost {
			t.Errorf("Método esperado POST, recebido %s", r.Method)
		}

		if r.Header.Get("Authorization") != "Bearer test_token" {
			t.Errorf("Authorization errado. Esperado 'Bearer test_token', recebido '%s'", r.Header.Get("Authorization"))
		}

		var payload model.MessagePayload
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &payload)
		if err != nil {
			t.Errorf("Erro ao decodificar JSON: %v", err)
		}

		if payload.To != "5511999999999" {
			t.Errorf("Número de destino incorreto. Esperado '5511999999999', recebido '%s'", payload.To)
		}

		if payload.Text.Body != "mensagem de teste" {
			t.Errorf("Conteúdo da mensagem incorreto. Esperado 'mensagem de teste', recebido '%s'", payload.Text.Body)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages":[{"id":"abc123"}]}`))
	}))
	defer mockServer.Close()

	mockHost := strings.TrimPrefix(mockServer.URL, "http://")

	client := &Client{
		PhoneNumberID: "test_id",
		AccessToken:   "test_token",
		ApiVersion:    "v22.0",
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
			Transport: &RedirectRoundTripper{
				original: http.DefaultTransport,
				mockHost: mockHost,
			},
		},
	}

	err := client.SendTextMessage([]string{"5511999999999"}, "mensagem de teste")
	if err != nil {
		t.Errorf("Erro ao enviar mensagem: %v", err)
	}
}

func TestSendTemplateMessage_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expectedPath := "/v22.0/test_id/messages"
		if r.URL.Path != expectedPath {
			t.Errorf("Caminho incorreto. Esperado %s, recebido %s", expectedPath, r.URL.Path)
		}

		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Authorization inválido. Esperado 'Bearer test-token', recebido '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type inválido. Esperado 'application/json', recebido '%s'", r.Header.Get("Content-Type"))
		}

		// Validar o payload do template
		body, _ := io.ReadAll(r.Body)
		fmt.Printf("Body recebido no mock: %s\n", string(body))

		var payload model.TemplatePayload
		err := json.Unmarshal(body, &payload)
		if err != nil {
			t.Errorf("Erro ao decodificar JSON: %v", err)
		}

		if payload.Template.Name != "template_name" {
			t.Errorf("Nome do template incorreto. Esperado 'template_name', recebido '%s'", payload.Template.Name)
		}

		if payload.Template.Language.Code != "pt_BR" {
			t.Errorf("Idioma incorreto. Esperado 'pt_BR', recebido '%s'", payload.Template.Language.Code)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages": [{"id": "gBEGkYiEB1VXAglK1ZEqA1YKPrU"}]}`))
	}))
	defer mockServer.Close()

	mockHost := strings.TrimPrefix(mockServer.URL, "http://")

	client := &Client{
		PhoneNumberID: "test_id",
		AccessToken:   "test-token",
		ApiVersion:    "v22.0",
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
			Transport: &RedirectRoundTripper{
				original: http.DefaultTransport,
				mockHost: mockHost,
			},
		},
	}

	// Adicionar o parâmetro params que estava faltando
	params := map[string]string{
		"1": "Jessica Santos",
	}

	err := client.SendTemplateMessage([]string{"5511999999999"}, "template_name", "pt_BR", params)
	if err != nil {
		t.Errorf("Erro ao enviar mensagem de template: %v", err)
	}
}
