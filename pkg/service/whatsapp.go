package service

type WhatsAppServiceInterface interface {
	SendTextMessage(to []string, message string) error
	SendTemplateMessage(to []string, templateName string, language string) error
}

type WhatsApp_service struct {
	Client WhatsAppServiceInterface
}

func NewWhatsAppService(client WhatsAppServiceInterface) *WhatsApp_service {
	return &WhatsApp_service{Client: client}
}
