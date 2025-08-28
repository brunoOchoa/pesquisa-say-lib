package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/brunoOchoa/pesquisa-say-lib/pkg/model"
)

const API_BASE_URL = "https://graph.facebook.com"

func (c *Client) sendRequest(method, endpoint string, body interface{}, result interface{}) error {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("erro ao serializar JSON: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	} else {
		reqBody = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro na requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("erro na API: status %d, resposta: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("erro ao decodificar resposta: %w", err)
		}
	}

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

		if err := c.sendRequest("POST", fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, c.ApiVersion, c.PhoneNumberID), payload, nil); err != nil {
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

		if err := c.sendRequest("POST", fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, c.ApiVersion, c.PhoneNumberID), payload, nil); err != nil {
			return fmt.Errorf("erro ao enviar template para %s: %w", phone, err)
		}
	}
	return nil
}
