package lib

import (
	"encoding/json"
	"fmt"

	"github.com/brunoOchoa/pesquisa-say-lib/pkg/model"
)

// ParseBodyFromWebhook extrai os corpos das mensagens do webhook do WhatsApp.
func ParseBodyFromWebhook(webhook *model.Webhook) ([]model.MessageBodyInfo, error) {
	var infos []model.MessageBodyInfo

	if webhook == nil || len(webhook.Entry) == 0 {
		return nil, fmt.Errorf("webhook vazio ou inválido")
	}

	for _, entry := range webhook.Entry {
		for _, change := range entry.Changes {
			for _, msg := range change.Value.Messages {
				info := model.MessageBodyInfo{
					MessageID: msg.ID,
					From:      msg.From,
					Timestamp: msg.Timestamp,
					Type:      msg.Type,
				}
				if msg.Text != nil {
					info.Body = msg.Text.Body
				}
				infos = append(infos, info)
			}
		}
	}
	return infos, nil
}

// GetBody recebe o JSON do webhook, faz o parse e retorna os corpos das mensagens
func (c *Client) GetBody(webhookJSON []byte) ([]model.MessageBodyInfo, error) {
	var webhookObj model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhookObj); err != nil {
		return nil, fmt.Errorf("erro ao decodificar webhook: %w", err)
	}

	infos, err := ParseBodyFromWebhook(&webhookObj)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
