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
		option.WithToken("<your_token>"))
}

func TestListResourceFile(t *testing.T) {
	res, err := client.Resource().ListFile(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}
