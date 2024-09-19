package main

import (
	"context"
	"fmt"
	"testing"

	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
	"github.com/hongyuxuan/bkcmdb-sdk-go/core/option"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = bkcmdb.NewClient(
		option.WithBkUser("admin"),
		option.WithSupplier("0"),
		// option.WithDebug(true),
		option.WithBaseUrl("http://10.50.219.26:8080"))
}

func TestListInstAssociation(t *testing.T) {
	res, err := client.Instance("postgresql").InstanceAssociation(955).List(context.Background(), "postgresql_run_iaas_vm", option.WithLimit(1))
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestCreateInstAssociation(t *testing.T) {
	res, err := client.Instance("postgresql").InstanceAssociation(955).Create(context.Background(), "postgresql_run_iaas_vm", 1723)
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res)
	}
}

func TestDeleteInstAssociation(t *testing.T) {
	err := client.Instance("postgresql").InstanceAssociation(955).Delete(context.Background(), 1425)
	assert.Nil(t, err)
}
