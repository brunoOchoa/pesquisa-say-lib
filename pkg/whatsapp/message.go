package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

const API_BASE_URL = "https://graph.facebook.com"

func (c *Client) sendRequest(body interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("erro ao serializar JSON: %w", err)
	}

	url := fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, c.ApiVersion, c.PhoneNumberID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro na requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	var respBody bytes.Buffer
	respBody.ReadFrom(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("erro na API: status %d, resposta: %s", resp.StatusCode, respBody.String())
	}

	fmt.Printf("✅ Resposta da API: %s\n", respBody.String())
	return nil
}

func (c *Client) SendTextMessage(to []string, message string) error {
	for _, phone := range to {
		payload := model.MessagePayload{
			MessagingProduct: "whatsapp",
			To:               phone,
			Type:             "text",
			Text: model.TextContent{
				Body: message,
			},
		}

		if err := c.sendRequest(payload); err != nil {
			return fmt.Errorf("erro ao enviar mensagem para %s: %w", phone, err)
		}
	}
	return nil
}

func (c *Client) SendTemplateMessage(to []string, template, language string) error {
	for _, phone := range to {
		payload := model.TemplatePayload{
			MessagingProduct: "whatsapp",
			To:               phone,
			Type:             "template",
		}
		payload.Template.Name = template
		payload.Template.Language.Code = language

		if err := c.sendRequest(payload); err != nil {
			return fmt.Errorf("erro ao enviar template para %s: %w", phone, err)
		}
	}
	return nil
}
