package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("365adefd2fac542d43cfde4103f52cdd"),
		option.WithAdminToken("d7a4f25936bc25f24afe21705ecba337"))
}

func TestListWarningGroup(t *testing.T) {
	res, err := client.WarningGroup().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10))
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}
