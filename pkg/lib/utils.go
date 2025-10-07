package lib

import (
	"encoding/json"
	"fmt"

	"github.com/brunoOchoa/pesquisa-say-lib/pkg/model"
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

func (c *Client) ExtractCommonInfo(webhookJSON []byte) ([]model.CommonWebhookInfo, error) {
	var webhook model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhook); err != nil {
		return nil, err
	}
	var infos []model.CommonWebhookInfo
	exists := make(map[string]bool)

	for _, entry := range webhook.Entry {
		for _, change := range entry.Changes {
			// Mensagens
			for _, msg := range change.Value.Messages {
				key := msg.From + msg.ID
				info := model.CommonWebhookInfo{
					WABA:      entry.ID,
					MessageID: msg.ID,
					Timestamp: msg.Timestamp,
					WaID:      msg.From,
				}
				infos = append(infos, info)
				exists[key] = true
				exists[msg.From] = true
			}
			// Status
			for _, status := range change.Value.Statuses {
				key := status.RecipientID + status.ID
				info := model.CommonWebhookInfo{
					WABA:      entry.ID,
					MessageID: status.ID,
					Timestamp: status.Timestamp,
					Status:    string(status.Status),
					WaID:      status.RecipientID,
				}
				infos = append(infos, info)
				exists[key] = true
				exists[status.RecipientID] = true
			}
			// Contatos
			for _, contact := range change.Value.Contacts {
				if !exists[contact.WAID] {
					info := model.CommonWebhookInfo{
						WABA: entry.ID,
						WaID: contact.WAID,
					}
					infos = append(infos, info)
					exists[contact.WAID] = true
				}
			}
		}
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("nenhum dado comum encontrado")
	}
	return infos, nil
}
