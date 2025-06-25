package whatsapp

import (
	"fmt"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

// StatusInfo representa as informações extraídas do webhook
type StatusInfo struct {
	MessageID   string
	Status      string
	Timestamp   string
	RecipientID string
	Errors      []map[string]interface{}
}

// ParseStatusFromWebhook extrai status do objeto recebido no webhook
func ParseStatusFromWebhook(webhook *model.Webhook) ([]StatusInfo, error) {
	var infos []StatusInfo

	if webhook == nil || len(webhook.Entry) == 0 {
		return nil, fmt.Errorf("webhook vazio ou inválido")
	}

	for _, entry := range webhook.Entry {
		for _, change := range entry.Changes {
			for _, status := range change.Value.Statuses {
				info := StatusInfo{
					MessageID:   status.ID,
					Status:      status.Status,
					Timestamp:   status.Timestamp,
					RecipientID: status.RecipientID,
				}
				for _, errObj := range status.Errors {
					info.Errors = append(info.Errors, map[string]interface{}{
						"code":      errObj.Code,
						"title":     errObj.Title,
						"message":   errObj.Message,
						"errorData": errObj.ErrorData,
						"href":      errObj.Href,
					})
				}
				infos = append(infos, info)
			}
		}
	}
	return infos, nil
}

func (c *Client) GetMessageStatus(messageID string) (string, error) {
	endpoint := fmt.Sprintf("%s/%s/messages/%s", API_BASE_URL, c.ApiVersion, messageID)
	var result map[string]interface{}
	err := c.sendRequest("GET", endpoint, nil, &result)
	if err != nil {
		return "", err
	}

	// Ajuste conforme o formato real da resposta da API
	status, _ := result["status"].(string)
	return status, nil
}
