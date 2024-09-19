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
	"github.com/samber/lo"
)

type InstanceAssociation struct {
	httpclient *req.Client
	bkObjId    string
	bkInstId   int64
	config     *config.Config
}

func NewInstanceAssociation(c *config.Config, bkObjId string, bkInstId int64) *InstanceAssociation {
	return &InstanceAssociation{
		httpclient: c.Httpclient,
		bkObjId:    bkObjId,
		bkInstId:   bkInstId,
		config:     c,
	}
}

type instAssociationResponse struct {
	types.BaseResponse
	Data types.InstAssociationList `json:"data"`
}

func (i *InstanceAssociation) List(ctx context.Context, bkObjAsstId string, opts ...option.ListOptionFunc) (resp *types.ListResponse, err error) {
	var objId string

	// contains
	var res1 instAssociationResponse
	req1 := types.ListInstAssociationRequest{
		Condition: types.InstAssociationCondition{
			BkObjId:  i.bkObjId,
			BkInstId: i.bkInstId,
		},
		BkObjId: i.bkObjId,
	}
	if bkObjAsstId != "" {
		req1.Condition.BkObjAsstId = bkObjAsstId
	}
	if err = i.httpclient.Post("/api/v3/find/instassociation").SetBody(&req1).SetSuccessResult(&res1).Do(ctx).Err; err != nil {
		return
	}
	if res1.Code != constant.ERR_OK {
		return nil, errorx.NewError(res1.Code, res1.Message, res1.Data)
	}
	ids := lo.Map(res1.Data, func(item types.InstAssociation, _ int) int64 {
		objId = item.BkAsstObjId
		return item.BkAsstInstId
	})
	if len(ids) > 0 {
		return listInstance(ctx, New(i.config, objId), ids, opts...)
	}

	// belong
	var res2 instAssociationResponse
	req2 := types.ListInstAssociationRequest{
		Condition: types.InstAssociationCondition{
			BkAsstObjId:  i.bkObjId,
			BkAsstInstId: i.bkInstId,
		},
		BkObjId: i.bkObjId,
	}
	if bkObjAsstId != "" {
		req2.Condition.BkObjAsstId = bkObjAsstId
	}
	if err = i.httpclient.Post("/api/v3/find/instassociation").SetBody(&req2).SetSuccessResult(&res2).Do(ctx).Err; err != nil {
		return
	}
	if res2.Code != constant.ERR_OK {
		return nil, errorx.NewError(res2.Code, res2.Message, res2.Data)
	}
	ids = lo.Map(res2.Data, func(item types.InstAssociation, _ int) int64 {
		objId = item.BkObjId
		return item.BkInstId
	})
	if len(ids) > 0 {
		return listInstance(ctx, New(i.config, objId), ids, opts...)
	}
	return
}

func listInstance(ctx context.Context, instance *Instance, ids []int64, opts ...option.ListOptionFunc) (resp *types.ListResponse, err error) {
	opts = append(opts, option.WithCondition(&types.Conditions{
		Condition: "AND",
		Rules: []types.Rule{
			{
				Field:    "bk_inst_id",
				Operator: "in",
				Value:    ids,
			},
		},
	}))
	return instance.List(ctx, opts...)
}

func (i *InstanceAssociation) Create(ctx context.Context, bkObjAsstId string, bkAsstInstId int64) (resp map[string]interface{}, err error) {
	var res types.PostResponse
	if err = i.httpclient.Post("/api/v3/create/instassociation").SetBody(&types.InstAssociationCondition{
		BkAsstInstId: bkAsstInstId,
		BkInstId:     i.bkInstId,
		BkObjAsstId:  bkObjAsstId,
	}).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Message, res.Data)
	}
	return res.Data, nil
}

func (i *InstanceAssociation) Delete(ctx context.Context, id int64) (err error) {
	var res types.PostResponse
	if err = i.httpclient.Delete(fmt.Sprintf("/api/v3/delete/instassociation/%s/%d", i.bkObjId, id)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Message, res.Data)
	}
	return
}
