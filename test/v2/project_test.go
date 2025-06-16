package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestProject struct {
	suite.Suite
	client      *ds.ClientV2
	projectCode int64
}

func (s *SuiteTestProject) SetupSuite() {
	s.client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
	s.projectCode = 10457530170432
}

func (s *SuiteTestProject) Test1ListProject() {
	res, err := s.client.Project(nil).List(context.Background(), option.WithPageNo(1), option.WithPageSize(10), option.WithSearchVal("deploy"))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteProject(t *testing.T) {
	suite.Run(t, new(SuiteTestProject))
}
