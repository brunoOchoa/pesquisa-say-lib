package whatsapp

import (
	"encoding/json"
	"fmt"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

func (c *Client) IdentifyWebhookType(webhookJSON []byte) (*model.WebhookResult, error) {
	var webhook model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhook); err != nil {
		return nil, err
	}
	if len(webhook.Entry) == 0 || len(webhook.Entry[0].Changes) == 0 {
		return &model.WebhookResult{Type: "unknown"}, nil
	}

	value := webhook.Entry[0].Changes[0].Value
	switch {
	case len(value.Statuses) > 0:
		infos, err := c.GetStatuses(webhookJSON)
		return &model.WebhookResult{Type: "statuses", Statuses: infos}, err
	case len(value.Messages) > 0:
		infos, err := c.GetBody(webhookJSON)
		return &model.WebhookResult{Type: "messages", Messages: infos}, err
	default:
		return &model.WebhookResult{Type: "unknown"}, nil
	}
}

func (c *Client) ExtractCommonWebhookInfo(webhookJSON []byte) ([]model.CommonWebhookInfo, error) {
	var webhook model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhook); err != nil {
		return nil, err
	}
	var infos []model.CommonWebhookInfo

	for _, entry := range webhook.Entry {
		for _, change := range entry.Changes {
			// Se houver mensagens
			for _, msg := range change.Value.Messages {
				info := model.CommonWebhookInfo{
					MessageID: msg.ID,
					Timestamp: msg.Timestamp,
				}
				if msg.From != "" {
					info.WaID = msg.From
				}
				infos = append(infos, info)
			}
			// Se houver contatos (ex: em mensagens recebidas)
			for _, contact := range change.Value.Contacts {
				info := model.CommonWebhookInfo{
					WaID: contact.WAID,
				}
				infos = append(infos, info)
			}
			// Se houver status
			for _, status := range change.Value.Statuses {
				info := model.CommonWebhookInfo{
					MessageID: status.ID,
					Timestamp: status.Timestamp,
					Status:    status.Status,
					WaID:      status.RecipientID,
				}
				infos = append(infos, info)
			}
		}
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("nenhum dado comum encontrado")
	}
	return infos, nil
}
