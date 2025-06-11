package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"lib.com/pkg/model"
)

func (c *Client) SendTextMessage(to, message string) error {
	payload := model.MessagePayload{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "text",
		Text: model.TextContent{
			Body: message,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", c.PhoneNumberID)

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

	if resp.StatusCode >= 300 {
		return fmt.Errorf("erro ao enviar mensagem: status %d", resp.StatusCode)
	}

	return nil
}
