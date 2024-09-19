package option

import (
	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
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

func WithBkUser(username string) ClientOptionFunc {
	return func(c *config.Config) {
		c.Header["BK_USER"] = username
	}
}

func WithSupplier(supplier string) ClientOptionFunc {
	return func(c *config.Config) {
		c.Header["HTTP_BLUEKING_SUPPLIER_ID"] = supplier
	}
}
