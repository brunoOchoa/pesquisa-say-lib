package whatsapp

import (
	"net/http"

	"github.com/brunoOchoa/whatsapp-lib/config"
)

type Client struct {
	AccessToken   string
	PhoneNumberID string
	ApiVersion    string
	HttpClient    *http.Client
}

func NewClient(cfg *config.ApiMetaConfig) *Client {
	return &Client{
		AccessToken:   cfg.AccessToken,
		PhoneNumberID: cfg.PhoneNumberID,
		ApiVersion:    cfg.ApiVersion,
		HttpClient:    &http.Client{},
	}
}
