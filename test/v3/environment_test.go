package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestEnvironment struct {
	suite.Suite
	client *ds.ClientV3
}

func (s *SuiteTestEnvironment) SetupSuite() {
	s.client = ds.NewClientV3(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func (s *SuiteTestEnvironment) TestListEnvironment() {
	res, err := s.client.Environment().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteEnvironment(t *testing.T) {
	suite.Run(t, new(SuiteTestEnvironment))
}
