package object

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type ObjectAssociation struct {
	httpclient *req.Client
	BkObjId    string
}

func NewObjectAssociation(c *config.Config, bkObjId string) *ObjectAssociation {
	return &ObjectAssociation{
		httpclient: c.Httpclient,
		BkObjId:    bkObjId,
	}
}

type objAssociationResponse struct {
	types.BaseResponse
	Data types.ObjectassociationList `json:"data"`
}

func (o *ObjectAssociation) List(ctx context.Context) (resp types.ObjectassociationList, err error) {
	var res1 objAssociationResponse
	if err = o.httpclient.Post("/api/v3/find/objectassociation").SetBody(map[string]interface{}{
		"condition": map[string]interface{}{
			"bk_obj_id": o.BkObjId,
		},
	}).SetSuccessResult(&res1).Do(ctx).Err; err != nil {
		return
	}
	if res1.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res1.Message)
	}

	var res2 objAssociationResponse
	if err = o.httpclient.Post("/api/v3/find/objectassociation").SetBody(map[string]interface{}{
		"condition": map[string]interface{}{
			"bk_asst_obj_id": o.BkObjId,
		},
	}).SetSuccessResult(&res2).Do(ctx).Err; err != nil {
		return
	}
	if res2.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res2.Message)
	}

	return append(res1.Data, res2.Data...), nil
}

// Example
//
//	{
//	  "bk_obj_asst_id": "project_contain_clickhouse",
//	  "bk_obj_asst_name": "",
//	  "bk_obj_id": "project",
//	  "bk_asst_obj_id": "clickhouse",
//	  "bk_asst_id": "contain",
//	  "mapping": "1:n"
//	}
func (c *ObjectAssociation) Create(ctx context.Context, body *types.Objectassociation) (resp map[string]interface{}, err error) {
	var res types.PostResponse
	if err = c.httpclient.Post("/api/v3/create/objectassociation").SetBody(body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Data, nil
}
