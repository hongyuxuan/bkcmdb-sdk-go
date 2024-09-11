package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestListClassification(t *testing.T) {
	client := bkcmdb.NewClient(
		bkcmdb.WithBkUser("admin"),
		bkcmdb.WithSupplier("0"),
		bkcmdb.WithBaseUrl("http://10.50.219.26:8080"))

	res, err := client.Classification().ListObject(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonString())
	}
}
