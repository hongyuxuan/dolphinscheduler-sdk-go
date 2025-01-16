package config

import "github.com/imroc/req/v3"

type Config struct {
	Header      map[string]string
	BaseUrl     string
	EnableDebug bool
	Httpclient  *req.Client
	AdminToken  string
	Timeout     int64
}
