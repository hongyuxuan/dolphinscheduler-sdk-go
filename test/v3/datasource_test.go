package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestDatasource struct {
	suite.Suite
	client *ds.ClientV2
}

func (s *SuiteTestDatasource) SetupSuite() {
	s.client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"))
}

func (s *SuiteTestDatasource) TestListDatasource() {
	res, err := s.client.Datasource().List(context.Background(), option.WithPageNo(1), option.WithPageSize(10))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteDatasource(t *testing.T) {
	suite.Run(t, new(SuiteTestDatasource))
}
