package model

type StatusType string

const (
	Received  StatusType = "Received"
	Sent      StatusType = "Sent"
	Failed    StatusType = "Failed"
	Delivered StatusType = "Delivered"
	Read      StatusType = "Read"
)

type WebhookResult struct {
	Type     string            `json:"type"`
	Statuses []StatusInfo      `json:"statuses,omitempty"`
	Messages []MessageBodyInfo `json:"messages,omitempty"`
}

type StatusInfo struct {
	MessageID    string                   `json:"message_id"`
	Status       StatusType               `json:"status"`
	Timestamp    string                   `json:"timestamp"`
	RecipientID  string                   `json:"recipient_id"`
	Conversation []map[string]interface{} `json:"conversation,omitempty"`
	Errors       []map[string]interface{} `json:"errors,omitempty"`
}

type MessageBodyInfo struct {
	MessageID string `json:"message_id"`
	From      string `json:"from"`

	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Body      string `json:"body,omitempty"`
}

type LogResponse struct {
	Phone string `json:"phone"`
}

type CommonWebhookInfo struct {
	WaID      string // Número do usuário (wa_id ou from)
	MessageID string // wamid
	Timestamp string
	Status    string // Se houver
}
