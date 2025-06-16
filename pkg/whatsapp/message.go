package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

const (
	API_BASE_URL = "https://graph.facebook.com"
	API_VERSION  = "v22.0"
)

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

		jsonData, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, API_VERSION, c.PhoneNumberID)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var respBody bytes.Buffer
		respBody.ReadFrom(resp.Body)

		if resp.StatusCode >= 300 {
			return fmt.Errorf("erro ao enviar para %s: status %d, resposta: %s", phone, resp.StatusCode, respBody.String())
		}

		fmt.Printf("Mensagem enviada para %s, resposta da API: %s\n", phone, respBody.String())
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

		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, API_VERSION, c.PhoneNumberID)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var respBody bytes.Buffer
		respBody.ReadFrom(resp.Body)

		if resp.StatusCode >= 300 {
			return fmt.Errorf("erro ao enviar para %s: status %d, resposta: %s", phone, resp.StatusCode, respBody.String())
		}

		fmt.Printf("Mensagem enviada para %s, resposta da API: %s\n", phone, respBody.String())
	}
	return nil
}
