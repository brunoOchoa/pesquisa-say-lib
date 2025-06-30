package main

import (
	"encoding/json"
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
							"contacts": []interface{}{
								map[string]interface{}{
									"profile": map[string]interface{}{
										"name": "Bruno Ochoa",
									},
									"wa_id": "5521985421711",
								},
							},
							"messages": []interface{}{
								map[string]interface{}{
									"from":      "5521985421711",
									"id":        "wamid.HBgNNTUyMTk4NTQyMTcxMRUCABIYFDNGQUU0NDM0Q0U2MUI3MTRGRUU1AA==",
									"timestamp": "1750970525",
									"text": map[string]interface{}{
										"body": "Testando a lib usando GetBody",
									},
									"type": "text",
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

	// 2. Crie o client
	cfg := config.WhasAppLibConfig()
	client := whatsapp.NewClient(cfg)

	// 3. Use IdentifyWebhookType para extrair e imprimir os dados
	infos, err := client.IdentifyWebhookType(payloadBytes)
	if err != nil {
		log.Fatalf("Erro ao identificar tipo de conteúdo: %v", err)
	}

	log.Printf("Dados extraídos do webhook: %+v", infos)
}
