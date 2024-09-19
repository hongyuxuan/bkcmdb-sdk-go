package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/stretchr/testify/assert"
)

var client *bkcmdb.Client

func init() {
	client = bkcmdb.NewClient(
		option.WithBkUser("admin"),
		option.WithSupplier("0"),
		// option.WithDebug(true),
		option.WithBaseUrl("http://10.50.219.26:8080"))
}

func TestListClassification(t *testing.T) {
	res, err := client.Classification().ListObject(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonString())
	}
}
