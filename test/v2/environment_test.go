package test_v2

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
	client *ds.ClientV2
}

func (s *SuiteTestEnvironment) SetupSuite() {
	s.client = ds.NewClientV2(
		option.WithDebug(false),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func (s *SuiteTestEnvironment) Test1ListEnvironment() {
	res, err := s.client.Environment().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteEnvironment(t *testing.T) {
	suite.Run(t, new(SuiteTestEnvironment))
}
