package whatsapp

import (
	"encoding/json"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

func (c *Client) IdentifyWebhookType(webhookJSON []byte) (string, error) {
	var webhook model.Webhook
	if err := json.Unmarshal(webhookJSON, &webhook); err != nil {
		return "unknown", err
	}
	if len(webhook.Entry) == 0 || len(webhook.Entry[0].Changes) == 0 {
		return "unknown", nil
	}

	value := webhook.Entry[0].Changes[0].Value
	switch {
	case len(value.Statuses) > 0:
		return "statuses", nil
	case len(value.Messages) > 0:
		return "messages", nil
	default:
		return "unknown", nil
	}
}
