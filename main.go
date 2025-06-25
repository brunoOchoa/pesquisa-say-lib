package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brunoOchoa/whatsapp-lib/config"
	"github.com/brunoOchoa/whatsapp-lib/pkg/whatsapp"
)

func main() {
	// Exemplo de payload recebido do webhook do WhatsApp (como map)
	webhookPayload := map[string]interface{}{
		"object": "whatsapp_business_account",
		"entry": []interface{}{
			map[string]interface{}{
				"id": "1235292671536231",
				"changes": []interface{}{
					map[string]interface{}{
						"value": map[string]interface{}{
							"messaging_product": "whatsapp",
							"metadata": map[string]interface{}{
								"display_phone_number": "15556576647",
								"phone_number_id":      "650915048109117",
							},
							"statuses": []interface{}{
								map[string]interface{}{
									"id":           "wamid.HBgNNTUyMTk4NTE5MzUyNRUCABEYEjJFMTEyRkRBQjlBQzA5MjFGNwA=",
									"status":       "failed",
									"timestamp":    "1750773570",
									"recipient_id": "5521985193525",
									"errors": []interface{}{
										map[string]interface{}{
											"code":    131047,
											"title":   "Re-engagement message",
											"message": "Re-engagement message",
											"error_data": map[string]interface{}{
												"details": "Message failed to send because more than 24 hours have passed since the customer last replied to this number.",
											},
											"href": "https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/",
										},
									},
								},
							},
						},
						"field": "messages",
					},
				},
			},
		},
	}

	// 1. Converta o map para JSON
	payloadBytes, err := json.Marshal(webhookPayload)
	if err != nil {
		log.Fatalf("Erro ao serializar payload: %v", err)
	}

	// 2. Crie o client e o service
	cfg := config.WhatsAppLibpConfig()
	client := whatsapp.NewClient(cfg)
	svc := client // whatsapp.Client já implementa WhatsAppService

	// 3. Use o service.GetMessageStatus
	infos, err := svc.GetMessageStatus(payloadBytes)
	if err != nil {
		log.Println("Erro ao processar webhook:", err)
	}
	for _, info := range infos {
		fmt.Printf("Mensagem %s para %s está com status: %s\n", info.MessageID, info.RecipientID, info.Status)
		if len(info.Errors) > 0 {
			fmt.Printf("Erro da conversa: %+v\n", info.Errors)
		}
	}
}
