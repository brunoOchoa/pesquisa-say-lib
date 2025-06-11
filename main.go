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

	to := "5521997921747"
	msg := "Olá, mensagem de teste!"

	err := client.SendTextMessage(to, msg)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	fmt.Println("✅ Mensagem enviada com sucesso!")
}
