package bkcmdb

import (
	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/classification"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/object"
	"github.com/imroc/req/v3"
	"github.com/samber/lo"
)

type Client struct {
	Config *config.Config
}

type OptionFunc func(*config.Config)

func WithHeaders(header map[string]string) OptionFunc {
	return func(c *config.Config) {
		c.Header = lo.Assign(c.Header, header)
	}
}

func WithBaseUrl(baseUrl string) OptionFunc {
	return func(c *config.Config) {
		c.BaseUrl = baseUrl
	}
}

func WithDebug(enable bool) OptionFunc {
	return func(c *config.Config) {
		c.EnableDebug = enable
	}
}

func WithBkUser(username string) OptionFunc {
	return func(c *config.Config) {
		c.Header["BK_USER"] = username
	}
}

func WithSupplier(supplier string) OptionFunc {
	return func(c *config.Config) {
		c.Header["HTTP_BLUEKING_SUPPLIER_ID"] = supplier
	}
}

func NewClient(opts ...OptionFunc) *Client {
	config := &config.Config{
		Header: make(map[string]string),
	}
	for _, opt := range opts {
		opt(config)
	}

	httpclient := req.C().
		OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			if req.RetryAttempt > 0 {
				return nil
			}
			req.EnableDump()
			return nil
		}).
		OnAfterResponse(func(client *req.Client, res *req.Response) (err error) {
			if res.Err != nil {
				err = res.Err
			}
			return
		})

	httpclient.SetBaseURL(config.BaseUrl)
	httpclient.SetCommonHeaders(config.Header)
	if config.EnableDebug {
		httpclient.EnableDebugLog()
		httpclient.EnableDumpAll()
	} else {
		httpclient.DisableDebugLog()
		httpclient.DisableDumpAll()
	}
	config.Httpclient = httpclient

	return &Client{
		Config: config,
	}
}

func (c *Client) Object(bkObjId string) *object.Object {
	return object.New(c.Config, bkObjId)
}

func (c *Client) Classification() *classification.Classification {
	return classification.New(c.Config)
}
