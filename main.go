package main

import (
	"io"
	"log"
	"os"

	"github.com/brunoOchoa/pesquisa-say-lib/config"
	"github.com/brunoOchoa/pesquisa-say-lib/pkg/lib"
)

func main() {
	// 1. Leia o conteúdo do arquivo object_utility.json
	file, err := os.Open("doc/object_messages.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	payloadBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// 2. Crie o client
	cfg := config.LibConfig()
	client := lib.NewClient(cfg)

	// 3. Use IdentifyWebhookType para extrair e imprimir os dados
	// infos, err := client.IdentifyWebhookType(payloadBytes)
	// if err != nil {
	// 	log.Fatalf("Erro ao identificar tipo de conteúdo: %v", err)
	// }
	// switch infos.Type {
	// case "statuses":
	// 	log.Printf("Status extraídos: %+v", infos.Statuses)
	// case "messages":
	// 	log.Printf("Mensagens extraídas: %+v", infos.Messages)
	// default:
	// 	log.Println("Webhook não contém mensagens nem status.")
	// }

	infos, err := client.ExtractCommonInfo(payloadBytes)
	if err != nil {
		log.Println("Erro ao extrair informações comuns:", err)
	}
	log.Printf("Informações comuns extraídas: %+v", infos)
}
