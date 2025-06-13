package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"lib.com/pkg/model"
)

func (c *Client) SendTextMessage(to []string, message string) error {
	for _, phone := range to {
		payload := model.MessagePayload{
			MessagingProduct: "whatsapp",
			To:               phone, // Agora é string
			Type:             "text",
			Text: model.TextContent{
				Body: message,
			},
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.PhoneNumberID)

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

// func (c *Client) SendTemplateMessage(to, template, language string) error {

// 	payload := model.TemplatePayload{
// 		MessagingProduct: "whatsapp",
// 		To:               to,
// 		Type:             "template",
// 	}
// 	payload.Template.Name = "hello_world"
// 	payload.Template.Language.Code = "en_US"
// 	// Se quiser passar parâmetros para o template, adicione aqui:
// 	// payload.Template.Components = []model.TemplateComponent{ ... }

// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}

// 	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.PhoneNumberID)

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.HttpClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	var respBody bytes.Buffer
// 	respBody.ReadFrom(resp.Body)

// 	if resp.StatusCode >= 300 {
// 		return fmt.Errorf("erro: status %d, resposta: %s", resp.StatusCode, respBody.String())
// 	}

// 	fmt.Println("Resposta da API:", respBody.String())

// 	return nil
// }

func (c *Client) SendTemplateMessage(to []string, template, language string) error {
	for _, phone := range to {
		payload := model.TemplatePayload{
			MessagingProduct: "whatsapp",
			To:               phone,
			Type:             "template",
		}
		payload.Template.Name = template
		payload.Template.Language.Code = language
		// Se quiser passar parâmetros para o template, adicione aqui:
		// payload.Template.Components = []model.TemplateComponent{ ... }

		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.PhoneNumberID)

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
