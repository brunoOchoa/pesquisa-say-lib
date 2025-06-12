package model

type MessagePayload struct {
	MessagingProduct string      `json:"messaging_product"`
	To               string      `json:"to"`
	Type             string      `json:"type"`
	Text             TextContent `json:"text"`
}

type TextContent struct {
	Body string `json:"body"`
}

type TemplateComponent struct {
	Type       string `json:"type"`
	Parameters []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"parameters"`
}

type MessageRequest struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

type TemplatePayload struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Template         struct {
		Name     string `json:"name"`
		Language struct {
			Code string `json:"code"`
		} `json:"language"`
		Components []TemplateComponent `json:"components,omitempty"`
	} `json:"template"`
}
