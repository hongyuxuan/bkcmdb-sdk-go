package object

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/errorx"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type ObjectAssociation struct {
	httpclient *req.Client
	bkObjId    string
}

func NewObjectAssociation(c *config.Config, bkObjId string) *ObjectAssociation {
	return &ObjectAssociation{
		httpclient: c.Httpclient,
		bkObjId:    bkObjId,
	}
}

type objAssociationResponse struct {
	types.BaseResponse
	Data types.ObjectAssociationList `json:"data"`
}

func (o *ObjectAssociation) List(ctx context.Context) (resp types.ObjectAssociationList, err error) {
	var res1 objAssociationResponse
	if err = o.httpclient.Post("/api/v3/find/objectassociation").SetBody(map[string]interface{}{
		"condition": map[string]interface{}{
			"bk_obj_id": o.bkObjId,
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
			"bk_asst_obj_id": o.bkObjId,
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
func (o *ObjectAssociation) Create(ctx context.Context, body *types.ObjectAssociation) (resp map[string]interface{}, err error) {
	var res types.PostResponse
	if err = o.httpclient.Post("/api/v3/create/objectassociation").SetBody(body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Message, res.Data)
	}
	return res.Data, nil
}

func (o *ObjectAssociation) Delete(ctx context.Context, bkObjAsstId string) (err error) {
	var associations types.ObjectAssociationList
	if associations, err = o.List(ctx); err != nil {
		return
	}
	for _, association := range associations {
		if association.BkObjAsstId == bkObjAsstId {
			var res types.PostResponse
			if err = o.httpclient.Delete(fmt.Sprintf("/api/v3/delete/objectassociation/%d", association.Id)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
				return
			}
			if res.Code != constant.ERR_OK {
				return errorx.NewError(res.Code, res.Message, res.Data)
			}
			return nil
		}
	}
	return fmt.Errorf("未找到关联关系: %s", bkObjAsstId)
}
