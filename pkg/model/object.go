package model

type Webhook struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
	Statuses         []Statuses `json:"statuses"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Statuses struct {
	ID           string        `json:"id"`
	Status       string        `json:"status"`
	Timestamp    string        `json:"timestamp"`
	RecipientID  string        `json:"recipient_id"`
	Conversation Conversation  `json:"conversation,omitempty"`
	Errors       []StatusError `json:"errors,omitempty"`
}

type Conversation struct {
	ID     string `json:"id"`
	Origin Origin `json:"origin"`
}

type Origin struct {
	Type string `json:"type"`
}
type StatusError struct {
	Code      int                    `json:"code"`
	Title     string                 `json:"title"`
	Message   string                 `json:"message"`
	ErrorData map[string]interface{} `json:"error_data"`
	Href      string                 `json:"href"`
}
