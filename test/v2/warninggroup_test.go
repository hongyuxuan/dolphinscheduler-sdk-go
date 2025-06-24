package test_v2

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestWarningGroup struct {
	suite.Suite
	client *ds.ClientV2
}

func (s *SuiteTestWarningGroup) SetupSuite() {
	s.client = ds.NewClientV2(
		option.WithDebug(false),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func (s *SuiteTestWarningGroup) Test1ListWarningGroup() {
	res, err := s.client.WarningGroup().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteWarningGroup(t *testing.T) {
	suite.Run(t, new(SuiteTestWarningGroup))
}
