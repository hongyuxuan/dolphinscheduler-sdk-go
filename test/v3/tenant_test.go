package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestTenant struct {
	suite.Suite
	client *ds.ClientV3
}

func (s *SuiteTestTenant) SetupSuite() {
	s.client = ds.NewClientV3(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func (s *SuiteTestTenant) TestListTenant() {
	res, err := s.client.Tenant().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10), option.WithSearchVal(""))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteTenant(t *testing.T) {
	suite.Run(t, new(SuiteTestTenant))
}
