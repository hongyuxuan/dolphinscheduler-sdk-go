package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/assert"
)

var clientTenant *ds.ClientV2

func init() {
	clientTenant = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func TestListTenant(t *testing.T) {
	res, err := clientTenant.Tenant().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10), option.WithSearchVal(""))
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}
