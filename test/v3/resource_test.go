package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/stretchr/testify/suite"
)

type SuiteTestResource struct {
	suite.Suite
	client *ds.ClientV3
}

func (s *SuiteTestResource) SetupSuite() {
	s.client = ds.NewClientV3(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
}

func (s *SuiteTestResource) TestListResourceFile() {
	res, err := s.client.Resource().ListFile(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteResource(t *testing.T) {
	suite.Run(t, new(SuiteTestResource))
}
