package main

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	resourceoption "github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option/resource"
	"github.com/stretchr/testify/suite"
)

type SuiteTestResource struct {
	suite.Suite
	client       *ds.ClientV3
	resourceName string
}

func (s *SuiteTestResource) SetupSuite() {
	s.client = ds.NewClientV3(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
	s.resourceName = "dolphinscheduler-tidb/artifact_url.yaml"
}

func (s *SuiteTestResource) Test1ListResourceFile() {
	res, err := s.client.Resource().ListFile(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func (s *SuiteTestResource) Test2ViewAndModifyResourceFile() {
	res, err := s.client.Resource().ViewFile(
		context.Background(),
		resourceoption.WithFullName(s.resourceName),
		resourceoption.WithLimit(1000),
		resourceoption.WithTenantCode("test-tenant"),
	)
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.Content)
	}

	// modify file
	res.Content += "\nhello: world"
	_, err = s.client.Resource().ModifyFile(
		context.Background(),
		resourceoption.WithFullName(s.resourceName),
		resourceoption.WithContent(res.Content),
		resourceoption.WithTenantCode("test-tenant"),
	)
	s.Nil(err)
}

func TestSuiteResource(t *testing.T) {
	suite.Run(t, new(SuiteTestResource))
}
