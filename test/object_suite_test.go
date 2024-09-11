package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type SuiteTestObject struct {
	suite.Suite
	client           *bkcmdb.Client
	classificationId string
	bkObjId          string
	bkGroupId        string
	bkPropertyId     string
	bkPropertyName   string
}

func (s *SuiteTestObject) SetupSuite() {
	s.client = bkcmdb.NewClient(
		bkcmdb.WithBkUser("admin"),
		bkcmdb.WithSupplier("0"),
		// bkcmdb.WithDebug(true),
		bkcmdb.WithBaseUrl("http://10.50.219.26:8080"))

	s.bkObjId = "testobject"
	s.classificationId = "database"
	s.bkGroupId = "基本信息"
	s.bkPropertyId = "password"
	s.bkPropertyName = "密码"
}

func (s *SuiteTestObject) Test1CreateObject() {
	body := types.Object{
		Account:            "0",
		BkObjId:            s.bkObjId,
		BkObjName:          s.bkObjId,
		BkClassificationId: s.classificationId,
		BkObjIcon:          "icon-cc-default",
		Username:           "admin",
	}
	res, err := s.client.Object(s.bkObjId).Create(context.Background(), &body)
	s.Nil(err)
	if s.NotNil(res) {
		s.Equal(body.BkObjId, res["bk_obj_id"], fmt.Sprintf("create object failed with bk_obj_id=%s", body.BkObjId))
	}
}

func (s *SuiteTestObject) Test2CreateObjectAttrGroup() {
	body := types.ObjectAttrGroup{
		BkGroupId:   s.bkGroupId,
		BkGroupName: s.bkGroupId,
		BkObjId:     s.bkObjId,
		Account:     "0",
		IsCollapse:  false,
	}
	res, err := s.client.Object(s.bkObjId).ObjectAttrGroup().Create(context.Background(), &body)
	s.Nil(err)
	if s.NotNil(res) {
		s.Equal(body.BkGroupId, res["bk_group_id"], fmt.Sprintf("create object attr group failed with bk_group_id=%s", body.BkGroupId))
	}
}

func (s *SuiteTestObject) Test3ListObjectAttrGroup() {
	res, err := s.client.Object(s.bkObjId).ObjectAttrGroup().List(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonString())
		s.Equal(2, len(res), "new object attr group number should be 2: default and \"基本信息\"")
	}
}

func (s *SuiteTestObject) Test4CreateObjectAttr() {
	body := types.ObjectAttr{
		BkPropertyName:  s.bkPropertyName,
		BkPropertyId:    s.bkPropertyId,
		BkPropertyGroup: s.bkGroupId,
		Unit:            "",
		Placeholder:     "",
		BkPropertyType:  "enum",
		Editable:        true,
		Isrequired:      false,
		IsMultiple:      false,
		Option: []map[string]interface{}{
			{
				"id":         "dev",
				"is_default": true,
				"name":       "dev",
				"type":       "text",
			},
		},
		Creator: "admin",
		BkObjId: s.bkObjId,
		Account: "0",
	}
	res, err := s.client.Object(s.bkObjId).ObjectAttr().Create(context.Background(), &body)
	s.Nil(err)
	if s.NotNil(res) {
		s.Equal(body.BkPropertyId, res["bk_property_id"], fmt.Sprintf("create object attr failed with bk_property_id=%s", body.BkPropertyId))
	}
}

func (s *SuiteTestObject) Test5ListObjectAttr() {
	res, err := s.client.Object(s.bkObjId).ObjectAttr().List(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonString())
		find := false
		for _, attr := range res {
			if attr.BkPropertyId == s.bkPropertyId {
				find = true
			}
		}
		s.Equal(find, true, "object attr should have \"%s\"", s.bkPropertyId)
	}
}

func (s *SuiteTestObject) Test6DeleteObjectAttr() {
	err := s.client.Object(s.bkObjId).ObjectAttr().Delete(context.Background(), s.bkPropertyId)
	s.Nil(err)
	res, err := s.client.Object(s.bkObjId).ObjectAttr().List(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		_, ok := lo.Find(res, func(item types.ObjectAttr) bool {
			return item.BkPropertyId == s.bkPropertyId
		})
		s.Equal(false, ok, "object should not find attr \"%s\"", s.bkPropertyId)
	}
}

func (s *SuiteTestObject) Test7DeleteObjectAttrGroup() {
	err := s.client.Object(s.bkObjId).ObjectAttrGroup().Delete(context.Background(), s.bkGroupId)
	s.Nil(err)
	res, err := s.client.Object(s.bkObjId).ObjectAttrGroup().List(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		_, ok := lo.Find(res, func(item types.ObjectAttrGroup) bool {
			return item.BkGroupId == s.bkGroupId
		})
		s.Equal(false, ok, "object should not find attr group \"%s\"", s.bkGroupId)
	}
}

func (s *SuiteTestObject) TearDownSuite() {
	// delete object
	err := s.client.Object(s.bkObjId).Delete(context.Background())
	s.Nil(err)
}

func TestSuiteTestObject(t *testing.T) {
	suite.Run(t, new(SuiteTestObject))
}
