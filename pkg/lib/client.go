package lib

import (
	"net/http"

	"github.com/brunoOchoa/pesquisa-say-lib/config"
)

type Client struct {
	AccessToken   string
	PhoneNumberID string
	ApiVersion    string
	HttpClient    *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		AccessToken:   cfg.AccessToken,
		PhoneNumberID: cfg.PhoneNumberID,
		ApiVersion:    cfg.ApiVersion,
		HttpClient:    &http.Client{},
	}
}
