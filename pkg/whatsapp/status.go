package whatsapp

import (
	"encoding/json"
	"fmt"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

// ParseStatusFromWebhook extrai informações de status de mensagens do webhook do WhatsApp.
func ParseStatusFromWebhook(webhook *model.Webhook) ([]model.StatusInfo, error) {
	var infos []model.StatusInfo

	if webhook == nil || len(webhook.Entry) == 0 {
		return nil, fmt.Errorf("webhook vazio ou inválido")
	}

	for _, entry := range webhook.Entry {
		for _, change := range entry.Changes {
			for _, status := range change.Value.Statuses {
				info := model.StatusInfo{
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

func (c *Client) GetMessageStatus(webhookJSON []byte) ([]model.StatusInfo, error) {
	var webhookObj model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhookObj); err != nil {
		return nil, fmt.Errorf("erro ao decodificar webhook: %w", err)
	}

	infos, err := ParseStatusFromWebhook(&webhookObj)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
