package service

import "github.com/brunoOchoa/whatsapp-lib/pkg/model"

type WhatsAppService interface {
	SendTextMessage(to []string, message string) error
	SendTemplateMessage(to []string, templateName string, language string) error
	GetMessageStatus(webhookJSON []byte) ([]model.StatusInfo, error)
}

type whatsApp_service struct {
	Client WhatsAppService
}

func NewWhatsAppService(client WhatsAppService) *whatsApp_service {
	return &whatsApp_service{Client: client}
}
