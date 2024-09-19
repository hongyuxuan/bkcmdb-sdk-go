package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = bkcmdb.NewClient(
		option.WithBkUser("admin"),
		option.WithSupplier("0"),
		// option.WithDebug(true),
		option.WithBaseUrl("http://10.50.219.26:8080"))
}

func TestListObjectAssociation(t *testing.T) {
	res, err := client.Object("postgresql").ObjectAssociation().List(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestCreateObjectAssociation(t *testing.T) {
	body := types.ObjectAssociation{
		BkObjAsstId:   "project_contain_clickhouse",
		BkObjAsstName: "",
		BkObjId:       "project",
		BkAsstObjId:   "clickhouse",
		BkAsstId:      "contain",
		Mapping:       "1:n",
	}
	_, err := client.Object("project").ObjectAssociation().Create(context.Background(), &body)
	assert.Nil(t, err)
}

func TestDeleteObjectAssociation(t *testing.T) {
	err := client.Object("test").ObjectAssociation().Delete(context.Background(), "test_belong_project")
	assert.Nil(t, err)
}
