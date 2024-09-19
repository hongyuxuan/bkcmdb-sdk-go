package instance

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/errorx"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Instance struct {
	httpclient *req.Client
	bkObjId    string
	config     *config.Config
}

func New(c *config.Config, bkObjId string) *Instance {
	return &Instance{
		httpclient: c.Httpclient,
		bkObjId:    bkObjId,
		config:     c,
	}
}

func (i *Instance) InstanceAssociation(bkInstId int64) *InstanceAssociation {
	return NewInstanceAssociation(i.config, i.bkObjId, bkInstId)
}

func (i *Instance) List(ctx context.Context, opts ...option.ListOptionFunc) (resp *types.ListResponse, err error) {
	option := &types.ListOption{Page: &types.Page{Limit: 20}}
	for _, opt := range opts {
		opt(option)
	}
	var res1 types.ListInfoResponse
	if err = i.httpclient.Post(fmt.Sprintf("/api/v3/search/instances/object/%s", i.bkObjId)).SetBody(&option).SetSuccessResult(&res1).Do(ctx).Err; err != nil {
		return
	}
	if res1.Code != constant.ERR_OK {
		return nil, errorx.NewError(res1.Code, res1.Message, res1.Data)
	}
	var res2 types.ListCountResponse
	if err = i.httpclient.Post(fmt.Sprintf("/api/v3/count/instances/object/%s", i.bkObjId)).SetBody(&option).SetSuccessResult(&res2).Do(ctx).Err; err != nil {
		return
	}
	if res2.Code != constant.ERR_OK {
		return nil, errorx.NewError(res2.Code, res2.Message, res2.Data)
	}
	return &types.ListResponse{
		ListDataInfo:  res1.Data,
		ListDataCount: res2.Data,
	}, nil
}

func (i *Instance) Get(ctx context.Context, id int64) (resp interface{}, err error) {
	var res types.ListInfoResponse
	if err = i.httpclient.Post(fmt.Sprintf("/api/v3/search/instances/object/%s", i.bkObjId)).SetBody(&types.ListOption{
		Page:   &types.Page{Start: 0, Limit: 1},
		Fields: []string{},
		Conditions: &types.Conditions{
			Condition: "AND",
			Rules: []types.Rule{
				{
					Field:    "bk_inst_id",
					Operator: "equal",
					Value:    id,
				},
			},
		},
	}).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Message, res.Data)
	}
	if res.Data.Info != nil {
		return res.Data.Info[0], nil
	}
	return nil, fmt.Errorf("未找到 id = %d 的实例", id)
}

func (i *Instance) Create(ctx context.Context, body interface{}) (resp interface{}, err error) {
	var res types.PostResponse
	if err = i.httpclient.Post(fmt.Sprintf("/api/v3/create/instance/object/%s", i.bkObjId)).SetBody(&body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Message, res.Data)
	}
	return res.Data, nil
}

func (i *Instance) Delete(ctx context.Context, id int64) (err error) {
	var res types.PostResponse
	if err = i.httpclient.Delete(fmt.Sprintf("/api/v3/delete/instance/object/%s/inst/%d", i.bkObjId, id)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Message, res.Data)
	}
	return
}

func (i *Instance) Update(ctx context.Context, id int64, v map[string]interface{}) (err error) {
	var res types.PostResponse
	if err = i.httpclient.Put(fmt.Sprintf("/api/v3/update/instance/object/%s/inst/%d", i.bkObjId, id)).SetBody(v).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Message, res.Data)
	}
	return
}
