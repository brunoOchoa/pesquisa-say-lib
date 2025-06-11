package whatsapp

import (
	"net/http"
)

type Client struct {
	AccessToken   string
	PhoneNumberID string
	HttpClient    *http.Client
}

func NewClient(accessToken, phoneNumberID string) *Client {
	return &Client{
		AccessToken:   accessToken,
		PhoneNumberID: phoneNumberID,
		HttpClient:    &http.Client{},
	}
}
