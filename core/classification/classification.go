package classification

import (
	"context"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Classification struct {
	httpclient *req.Client
	config     *config.Config
}

type classificationResponse struct {
	types.BaseResponse
	Data types.ClassificationList `json:"data"`
}

func New(c *config.Config) *Classification {
	return &Classification{
		httpclient: c.Httpclient,
		config:     c,
	}
}

func (c *Classification) ListObject(ctx context.Context) (resp types.ClassificationList, err error) {
	var res classificationResponse
	if err = c.httpclient.Post("/api/v3/find/classificationobject").SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	return res.Data, nil
}
