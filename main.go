package main

import (
	"fmt"
	"log"

	"lib.com/config"
	"lib.com/pkg/whatsapp"
)

func main() {
	cfg := config.LoadConfig()

	client := whatsapp.NewClient(cfg.AccessToken, cfg.PhoneNumberID)

	to := []string{"5521985421711", "5521997921747"}
	msg := "Olá, mensagem depois de ter respondido a mensagem anterior!"

	err := client.SendTextMessage(to, msg)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	sendT := client.SendTemplateMessage(to, "hello_world", "en_US")
	if sendT != nil {
		log.Fatalf("Erro ao enviar template: %v", sendT)
	}

	fmt.Println("✅ Mensagem enviada com sucesso!")
}
