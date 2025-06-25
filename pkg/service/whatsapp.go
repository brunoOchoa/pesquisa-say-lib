package service

type WhatsAppService interface {
	SendTextMessage(to []string, message string) error
	SendTemplateMessage(to []string, templateName string, language string) error
	GetMessageStatus(messageID string) (string, error)
}

type whatsApp_service struct {
	Client WhatsAppService
}

func NewWhatsAppService(client WhatsAppService) *whatsApp_service {
	return &whatsApp_service{Client: client}
}
