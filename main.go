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
	file, err := os.Open("doc/object_utility.json")
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

	// 3. Testar IdentifyWebhookType
	// log.Println("\n=== Testando IdentifyWebhookType ===")
	// infos, err := client.IdentifyWebhookType(payloadBytes)
	// if err != nil {
	// 	log.Fatalf("Erro ao identificar tipo de conteúdo: %v", err)
	// }
	// log.Printf("Tipo: %s", infos.Type)
	// switch infos.Type {
	// case "statuses":
	// 	log.Printf("Total de Status: %d", len(infos.Statuses))
	// 	for i, status := range infos.Statuses {
	// 		log.Printf("  Status[%d]: MessageID=%s, Status=%s, RecipientID=%s, Timestamp=%s",
	// 			i, status.MessageID, status.Status, status.RecipientID, status.Timestamp)
	// 	}
	// case "messages":
	// 	log.Printf("Total de Mensagens: %d", len(infos.Messages))
	// 	for i, msg := range infos.Messages {
	// 		log.Printf("  Message[%d]: ID=%s, From=%s, Timestamp=%s, Type=%s",
	// 			i, msg.From, msg.Timestamp, msg.Type)
	// 	}
	// default:
	// 	log.Println("Webhook não contém mensagens nem status.")
	// }

	// 4. Testar ExtractCommonInfo
	log.Println("\n=== Testando ExtractCommonInfo ===")
	commonInfos, err := client.ExtractCommonInfo(payloadBytes)
	if err != nil {
		log.Println("Erro ao extrair informações comuns:", err)
	} else {
		log.Printf("Total de registros comuns: %d", len(commonInfos))
		for i, info := range commonInfos {
			log.Printf("  Info[%d]: WABA=%s, WaID=%s, MessageID=%s, Status=%s, Timestamp=%s",
				i, info.WABA, info.WaID, info.MessageID, info.Status, info.Timestamp)
		}
	}
}
