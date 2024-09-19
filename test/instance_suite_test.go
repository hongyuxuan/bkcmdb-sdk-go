package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func init() {
	client = bkcmdb.NewClient(
		option.WithBkUser("admin"),
		option.WithSupplier("0"),
		// option.WithDebug(true),
		option.WithBaseUrl("http://10.50.219.26:8080"))
}

type SuiteTestInstance struct {
	suite.Suite
	client     *bkcmdb.Client
	bkObjId    string
	bkInstName string
	bkInstId   int64
}

func (s *SuiteTestInstance) SetupSuite() {
	s.client = bkcmdb.NewClient(
		option.WithBkUser("admin"),
		option.WithSupplier("0"),
		// option.WithDebug(true),
		option.WithBaseUrl("http://10.50.219.26:8080"))

	s.bkObjId = "project"
	s.bkInstName = "testinst"
}

func (s *SuiteTestInstance) Test1CreateInstance() {
	body := map[string]interface{}{"bk_inst_name": s.bkInstName, "en_name": "bb", "project_name": "cc", "phecda_id": "dd", "po": "ee", "department": "ff", "project_number": "gg", "budget_number": "hh"}
	res, err := s.client.Instance(s.bkObjId).Create(context.Background(), body)
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res)
	}
}

func (s *SuiteTestInstance) Test2ListInstance() {
	res, err := s.client.Instance(s.bkObjId).List(context.Background(),
		option.WithLimit(5),
		option.WithStart(0),
		option.WithSort("bk_inst_id"),
		option.WithFields([]string{"bk_inst_name", "bk_inst_id"}),
		option.WithCondition(&types.Conditions{
			Condition: "AND",
			Rules: []types.Rule{
				{
					Field:    "bk_inst_name",
					Operator: "equal",
					Value:    s.bkInstName,
				},
			},
		}))
	s.Nil(err)
	if s.NotNil(res) {
		s.Equal(1, len(res.Info), "%s should have a instance with bk_inst_name=%s", s.bkObjId, s.bkInstName)
		inst := res.Info[0].(map[string]interface{})
		s.bkInstId = int64(inst["bk_inst_id"].(float64))
	}
}

func (s *SuiteTestInstance) Test3GetInstance() {
	_, err := s.client.Instance(s.bkObjId).Get(context.Background(), s.bkInstId)
	s.Nil(err)
}

func (s *SuiteTestInstance) Test4DeleteInstance() {
	err := s.client.Instance(s.bkObjId).Delete(context.Background(), s.bkInstId)
	s.Nil(err)
}

func TestSuiteTestInstance(t *testing.T) {
	suite.Run(t, new(SuiteTestInstance))
}

func TestListInstance(t *testing.T) {
	res, err := client.Instance("project").List(context.Background(),
		option.WithLimit(5),
		option.WithStart(0),
		option.WithSort("bk_inst_id"),
		option.WithFields([]string{"bk_inst_name", "bk_inst_id"}),
		option.WithCondition(&types.Conditions{
			Condition: "AND",
			Rules: []types.Rule{
				{
					Field:    "bk_inst_name",
					Operator: "contains",
					Value:    "ficc",
				},
			},
		}))
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestUpdateInstance(t *testing.T) {
	err := client.Instance("postgresql").Update(context.Background(), 951, map[string]interface{}{
		"memo": "helloworld",
	})
	assert.Nil(t, err)
}
