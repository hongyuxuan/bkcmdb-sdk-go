package object

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type AttrGroup struct {
	httpclient *req.Client
	BkObjId    string
	config     *config.Config
}

func NewAttrGroup(c *config.Config, bkObjId string) *AttrGroup {
	return &AttrGroup{
		httpclient: c.Httpclient,
		BkObjId:    bkObjId,
		config:     c,
	}
}

type attrGroupResponse struct {
	types.BaseResponse
	Data types.ObjectAttrGroupList `json:"data"`
}

func (a *AttrGroup) List(ctx context.Context) (resp types.ObjectAttrGroupList, err error) {
	var res attrGroupResponse
	if err = a.httpclient.Post(fmt.Sprintf("/api/v3/find/objectattgroup/object/%s", a.BkObjId)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Data, nil
}

// Example
//
//	{
//	  "bk_group_index": 2, // 可以不写，会自动计算
//	  "bk_group_id": "其它",
//	  "bk_group_name": "其它",
//	  "bk_obj_id": "hello",
//	  "bk_supplier_account": "0",
//	  "is_collapse": false
//	}
func (a *AttrGroup) Create(ctx context.Context, body *types.ObjectAttrGroup) (resp map[string]interface{}, err error) {
	var objAttrGroups types.ObjectAttrGroupList
	if objAttrGroups, err = a.List(ctx); err != nil {
		return
	}
	bkGroupIndex := len(objAttrGroups)
	body.BkGroupIndex = bkGroupIndex - 1 // Default group index is -1
	var res types.PostResponse
	if err = a.httpclient.Post("/api/v3/create/objectattgroup").SetBody(body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Data, nil
}

func (a *AttrGroup) Delete(ctx context.Context, attrGroupName string) (err error) {
	attrGroups, err := a.List(ctx)
	if err != nil {
		return
	}
	for _, group := range attrGroups {
		if group.BkGroupName == attrGroupName {
			var res types.PostResponse
			if err = a.httpclient.Delete(fmt.Sprintf("/api/v3/delete/objectattgroup/%d", group.Id)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
				return
			}
			if res.Code != constant.ERR_OK {
				return fmt.Errorf(res.Message)
			}
			break
		}
	}
	return
}
