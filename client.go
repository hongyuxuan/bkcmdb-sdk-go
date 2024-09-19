package bkcmdb

import (
	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/hongyuxuan/bkcmdb-sdk-go/service/classification"
	"github.com/hongyuxuan/bkcmdb-sdk-go/service/instance"
	"github.com/hongyuxuan/bkcmdb-sdk-go/service/object"
	"github.com/imroc/req/v3"
)

type Client struct {
	Config *config.Config
}

func NewClient(opts ...option.ClientOptionFunc) *Client {
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

func (c *Client) Instance(bkObjId string) *instance.Instance {
	return instance.New(c.Config, bkObjId)
}
