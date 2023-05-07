package fmp

import "net/http"

type Client struct {
	Config Config
	http.Client
}

type Config interface {
	GetFMPAPIKey() string
}
