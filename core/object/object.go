package object

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/bkcmdb-sdk-go/config"
	"github.com/hongyuxuan/bkcmdb-sdk-go/constant"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/classification"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Object struct {
	httpclient *req.Client
	BkObjId    string
	config     *config.Config
}

func New(c *config.Config, bkObjId string) *Object {
	return &Object{
		httpclient: c.Httpclient,
		BkObjId:    bkObjId,
		config:     c,
	}
}

func (o *Object) ObjectAttrGroup() *AttrGroup {
	return NewAttrGroup(o.config, o.BkObjId)
}

func (o *Object) ObjectAttr() *Attr {
	return NewAttr(o.config, o.BkObjId)
}

func (o *Object) ObjectAssociation() *ObjectAssociation {
	return NewObjectAssociation(o.config, o.BkObjId)
}

func (o *Object) Create(ctx context.Context, body *types.Object) (resp map[string]interface{}, err error) {
	body.BkObjId = o.BkObjId
	var res types.PostResponse
	if err = o.httpclient.Post("/api/v3/create/object").SetBody(body).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Data, nil
}

func (o *Object) Delete(ctx context.Context) (err error) {
	csif := classification.New(o.config)
	resp, csiferr := csif.ListObject(ctx)
	if csiferr != nil {
		return
	}
	var id int
	for _, i := range resp {
		for _, object := range i.BkObjects {
			if o.BkObjId == object.BkObjId {
				id = object.Id
				break
			}
		}
	}
	var res types.PostResponse
	if err = o.httpclient.Delete(fmt.Sprintf("/api/v3/delete/object/%d", id)).SetSuccessResult(&res).Do().Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return fmt.Errorf(res.Message)
	}
	return
}
