package lib

import (
	"encoding/json"
	"fmt"

	"github.com/brunoOchoa/pesquisa-say-lib/pkg/model"
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
					Status:      model.StatusType(status.Status),
					Timestamp:   status.Timestamp,
					RecipientID: status.RecipientID,
				}
				if status.Conversation.ID != "" || status.Conversation.Origin.Type != "" {
					info.Conversation = append(info.Conversation, map[string]interface{}{
						"id":     status.Conversation.ID,
						"origin": status.Conversation.Origin,
					})
				}

				if len(status.Errors) > 0 {
					errObj := status.Errors[0]
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

func (c *Client) GetStatuses(webhookJSON []byte) ([]model.StatusInfo, error) {
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
