package main

import (
	"fmt"
	"log"

	"github.com/brunoOchoa/whatsapp-lib/config"
	"github.com/brunoOchoa/whatsapp-lib/pkg/whatsapp"
)

func main() {
	cfg := config.LoadConfig()

	client := whatsapp.NewClient(cfg)

	to := []string{"5521985421711", "5521997921747"}
	// msg := "Olá, mensagem depois de ter respondido a mensagem anterior!"

	// err := client.SendTextMessage(to, msg)
	// if err != nil {
	// 	log.Fatalf("Erro ao enviar mensagem: %v", err)
	// }

	sendT := client.SendTemplateMessage(to, "hello_world", "en_US")
	if sendT != nil {
		log.Fatalf("Erro ao enviar template: %v", sendT)
	}

	fmt.Println("✅ Mensagem enviada com sucesso!")

	// Exemplo de uso do GetMessageStatus
	// Substitua pelo ID real retornado ao enviar uma mensagem
	messageID := "SEU_MESSAGE_ID_AQUI"
	status, err := client.GetMessageStatus(messageID)
	if err != nil {
		log.Fatalf("Erro ao buscar status da mensagem: %v", err)
	}
	fmt.Printf("Status da mensagem %s: %s\n", messageID, status)
}
