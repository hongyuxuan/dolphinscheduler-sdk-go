package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/assert"
)

var client *ds.ClientV2

func init() {
	client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"))
}

func TestListProject(t *testing.T) {
	res, err := client.Project(nil).List(context.Background(), option.WithPageNo(1), option.WithPageSize(10), option.WithSearchVal("deploy"))
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}
