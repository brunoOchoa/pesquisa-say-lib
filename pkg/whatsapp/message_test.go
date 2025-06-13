package whatsapp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"lib.com/pkg/model"
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
		if r.URL.Path != "/v22.0/test_id/messages" {
			t.Errorf("Caminho incorreto: %s", r.URL.Path)
		}

		if r.Method != http.MethodPost {
			t.Errorf("Método esperado POST, recebido %s", r.Method)
		}

		if r.Header.Get("Authorization") != "Bearer test_token" {
			t.Errorf("Authorization errado: %s", r.Header.Get("Authorization"))
		}

		var payload model.MessagePayload
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &payload)
		if err != nil {
			t.Errorf("Erro ao decodificar JSON: %v", err)
		}

		if payload.To != "5511999999999" || payload.Text.Body != "mensagem de teste" {
			t.Errorf("Payload incorreto: %+v", payload)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages":[{"id":"abc123"}]}`))
	}))
	defer mockServer.Close()

	mockHost := strings.TrimPrefix(mockServer.URL, "http://")

	client := &Client{
		PhoneNumberID: "test_id",
		AccessToken:   "test_token",
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
