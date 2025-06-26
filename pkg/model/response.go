package model

// Usado pela funcao GetMessageStatus para retornar o status de uma mensagem
type StatusInfo struct {
	MessageID   string
	Status      string
	Timestamp   string
	RecipientID string
	Errors      []map[string]interface{}
}
