package model

type WebhookResult struct {
	Type     string            `json:"type"`
	Statuses []StatusInfo      `json:"statuses,omitempty"`
	Messages []MessageBodyInfo `json:"messages,omitempty"`
}

type StatusInfo struct {
	MessageID   string                   `json:"message_id"`
	Status      string                   `json:"status"`
	Timestamp   string                   `json:"timestamp"`
	RecipientID string                   `json:"recipient_id"`
	Errors      []map[string]interface{} `json:"errors,omitempty"`
}

type MessageBodyInfo struct {
	MessageID string `json:"message_id"`
	From      string `json:"from"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Body      string `json:"body,omitempty"`
}
