package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListObjectAssociation(t *testing.T) {
	client := bkcmdb.NewClient(
		bkcmdb.WithBkUser("admin"),
		bkcmdb.WithSupplier("0"),
		bkcmdb.WithBaseUrl("http://bkcmdb_host:8080"))

	res, err := client.Object("postgresql").ObjectAssociation().List(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestCreateObjectAssociation(t *testing.T) {
	client := bkcmdb.NewClient(
		bkcmdb.WithBkUser("admin"),
		bkcmdb.WithSupplier("0"),
		bkcmdb.WithBaseUrl("http://bkcmdb_host:8080"))

	body := types.Objectassociation{
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
