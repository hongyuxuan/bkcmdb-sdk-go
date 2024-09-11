package object

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Attr struct {
	httpclient *req.Client
	BkObjId    string
	config     *config.Config
}

func NewAttr(c *config.Config, bkObjId string) *Attr {
	return &Attr{
		httpclient: c.Httpclient,
		BkObjId:    bkObjId,
		config:     c,
	}
}

type attrResponse struct {
	types.BaseResponse
	Data types.ObjectAttrList `json:"data"`
}

func (a *Attr) List(ctx context.Context) (resp types.ObjectAttrList, err error) {
	var res attrResponse
	if err = a.httpclient.Post("/api/v3/find/objectattr/web").SetBody(map[string]interface{}{
		"bk_obj_id":           a.BkObjId,
		"bk_supplier_account": "0",
	}).SetSuccessResult(&res).Do(ctx).Err; err != nil {
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
//	  "bk_property_name": "密码",
//	  "bk_property_id": "password",
//	  "bk_property_group": "其它",
//	  "unit": "",
//	  "placeholder": "",
//	  "bk_property_type": "enum",
//	  "editable": true,
//	  "isrequired": false,
//	  "ismultiple": false,
//	  "option": [
//	    {
//	      "id": "dev",
//	      "is_default": true,
//	      "name": "dev",
//	      "type": "text"
//	    }
//	  ],
//	  "creator": "admin",
//	  "bk_obj_id": "mongodb",
//	  "bk_supplier_account": "0"
//	}
func (a *Attr) Create(ctx context.Context, body *types.ObjectAttr) (resp map[string]interface{}, err error) {
	var res types.PostResponse
	if err = a.httpclient.Post("/api/v3/create/objectattr").SetBody(body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Data, nil
}

func (a *Attr) Delete(ctx context.Context, bkPropertyId string) (err error) {
	attrs, err := a.List(ctx)
	if err != nil {
		return
	}
	for _, attr := range attrs {
		if attr.BkPropertyId == bkPropertyId {
			var res types.PostResponse
			if err = a.httpclient.Delete(fmt.Sprintf("/api/v3/delete/objectattr/%d", attr.Id)).SetSuccessResult(&res).Do(ctx).Err; err != nil {
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
