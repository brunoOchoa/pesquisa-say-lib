package service

import (
	"github.com/brunoOchoa/whatsapp-lib/pkg/model"
)

type WhatsAppService interface {
	SendTextMessage(to []string, message string) error
	SendTemplateMessage(to []string, templateName string, language string) error
	GetStatuses(webhookJSON []byte) ([]model.StatusInfo, error)
	GetBody(webhookJSON []byte) ([]model.MessageBodyInfo, error)
	IdentifyWebhookType(webhookJSON []byte) (string, error)
	ExtractCommonInfo(webhookJSON []byte) ([]model.CommonWebhookInfo, error)
}

type whatsApp_service struct {
	Client WhatsAppService
}

func NewWhatsAppService(client WhatsAppService) *whatsApp_service {
	return &whatsApp_service{Client: client}
}
