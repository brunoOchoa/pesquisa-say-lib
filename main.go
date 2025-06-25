package main

import (
	"fmt"
	"log"

	"github.com/brunoOchoa/whatsapp-lib/pkg/whatsapp"
)

func main() {
	// cfg := config.LoadConfig()

	// client := whatsapp.NewClient(cfg)

	// to := []string{"5521985421711"} // Substitua pelo número de telefone do destinatário
	// msg := "Olá, mensagem depois de ter respondido a mensagem anterior!"

	// err := client.SendTextMessage(to, msg)
	// if err != nil {
	// 	log.Fatalf("Erro ao enviar mensagem: %v", err)
	// }

	// sendT := client.SendTemplateMessage(to, "hello_world", "en_US")
	// if sendT != nil {
	// 	log.Fatalf("Erro ao enviar template: %v", sendT)
	// }

	// fmt.Println("✅ Mensagem enviada com sucesso!")

	// Exemplo de uso do GetMessageStatus
	// Substitua pelo ID real retornado ao enviar uma mensagem

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
	infos, err := whatsapp.ParseStatusFromWebhook(webhookPayload)
	if err != nil {
		log.Println("Erro ao processar webhook:", err)
	}
	for _, info := range infos {
		fmt.Printf("Mensagem %s para %s está com status: %s\n", info.MessageID, info.RecipientID, info.Status)
		if len(info.Errors) > 0 {
			fmt.Printf("Erros: %+v\n", info.Errors)
		}
	}
}
