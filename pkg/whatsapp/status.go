package whatsapp

import (
	"fmt"
)

// StatusInfo representa as informações extraídas do webhook
type StatusInfo struct {
	MessageID   string
	Status      string
	Timestamp   string
	RecipientID string
	Errors      []map[string]interface{}
}

// ParseStatusFromWebhook extrai status do objeto recebido no webhook
func ParseStatusFromWebhook(webhook map[string]interface{}) ([]StatusInfo, error) {
	var infos []StatusInfo

	entries, ok := webhook["entry"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("campo 'entry' não encontrado ou inválido")
	}

	for _, entry := range entries {
		entryMap, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}
		changes, ok := entryMap["changes"].([]interface{})
		if !ok {
			continue
		}
		for _, change := range changes {
			changeMap, ok := change.(map[string]interface{})
			if !ok {
				continue
			}
			value, ok := changeMap["value"].(map[string]interface{})
			if !ok {
				continue
			}
			statuses, ok := value["statuses"].([]interface{})
			if !ok {
				continue
			}
			for _, status := range statuses {
				statusMap, ok := status.(map[string]interface{})
				if !ok {
					continue
				}
				info := StatusInfo{
					MessageID:   fmt.Sprintf("%v", statusMap["id"]),
					Status:      fmt.Sprintf("%v", statusMap["status"]),
					Timestamp:   fmt.Sprintf("%v", statusMap["timestamp"]),
					RecipientID: fmt.Sprintf("%v", statusMap["recipient_id"]),
				}
				if errs, ok := statusMap["errors"].([]interface{}); ok {
					for _, e := range errs {
						if errMap, ok := e.(map[string]interface{}); ok {
							info.Errors = append(info.Errors, errMap)
						}
					}
				}
				infos = append(infos, info)
			}
		}
	}
	return infos, nil
}

func (c *Client) GetMessageStatus(messageID string) (string, error) {
	endpoint := fmt.Sprintf("%s/%s/messages/%s", API_BASE_URL, c.ApiVersion, messageID)
	var result map[string]interface{}
	err := c.sendRequest("GET", endpoint, nil, &result)
	if err != nil {
		return "", err
	}

	// Ajuste conforme o formato real da resposta da API
	status, _ := result["status"].(string)
	return status, nil
}
