package service

type WhatsAppServiceInterface interface {
	SendTextMessage(to []string, message string) error
	SendTemplateMessage(to []string, templateName string, language string) error
}

type whatsApp_service struct {
	Client WhatsAppServiceInterface
}

func NewWhatsAppService(client WhatsAppServiceInterface) *whatsApp_service {
	return &whatsApp_service{Client: client}
}
