package option

import (
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/samber/lo"
)

type ClientOptionFunc func(*config.Config)

func WithHeaders(header map[string]string) ClientOptionFunc {
	return func(c *config.Config) {
		c.Header = lo.Assign(c.Header, header)
	}
}

func WithBaseUrl(baseUrl string) ClientOptionFunc {
	return func(c *config.Config) {
		c.BaseUrl = baseUrl
	}
}

func WithDebug(enable bool) ClientOptionFunc {
	return func(c *config.Config) {
		c.EnableDebug = enable
	}
}

func WithToken(token string) ClientOptionFunc {
	return func(c *config.Config) {
		c.Header["token"] = token
	}
}

func WithAdminToken(token string) ClientOptionFunc {
	return func(c *config.Config) {
		c.AdminToken = token
	}
}

func WithTimeout(timeout int64) ClientOptionFunc {
	return func(c *config.Config) {
		c.Timeout = timeout
	}
}
