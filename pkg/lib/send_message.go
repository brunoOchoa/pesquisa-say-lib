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

	fmt.Println("RequestBody:", reqBody.String()) // Log do corpo da requisição

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

func (c *Client) SendTemplateMessage(to []string, template, language string, params map[string]string) error {
	for _, phone := range to {
		payload := model.TemplatePayload{
			MessagingProduct: "whatsapp",
			To:               phone,
			Type:             "template",
		}
		payload.Template.Name = template
		payload.Template.Language.Code = language

		// Converter params para o formato esperado
		if len(params) > 0 {
			var templateParams []model.TemplateParam
			for _, value := range params {
				templateParams = append(templateParams, model.TemplateParam{
					Type:          "text",
					ParameterName: "nome",
					Text:          value,
				})
			}

			payload.Template.Components = []model.TemplateComponent{
				{
					Type:       "header",
					Parameters: templateParams,
				},
			}
		}

		if err := c.sendRequest("POST", fmt.Sprintf("%s/%s/%s/messages", API_BASE_URL, c.ApiVersion, c.PhoneNumberID), payload, nil); err != nil {
			return fmt.Errorf("erro ao enviar template para %s: %w", phone, err)
		}
	}
	return nil
}
