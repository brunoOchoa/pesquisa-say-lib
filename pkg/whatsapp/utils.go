package whatsapp

import (
	"encoding/json"

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
