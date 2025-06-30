package whatsapp

import (
	"encoding/json"

	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

func (c *Client) IdentifyWebhookType(webhookJSON []byte) (interface{}, error) {
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
		infos, err := c.GetStatuses(webhookJSON)
		return infos, err
	case len(value.Messages) > 0:
		infos, err := c.GetBody(webhookJSON)
		return infos, err
	default:
		return "unknown", nil
	}
}
