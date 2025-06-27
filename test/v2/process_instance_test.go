package test_v2

import (
	"context"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	instanceoption "github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option/instance"
	"github.com/stretchr/testify/suite"
)

type SuiteTestProcessInstance struct {
	suite.Suite
	client      *ds.ClientV2
	projectCode int64
}

func (s *SuiteTestProcessInstance) SetupSuite() {
	s.client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
	s.projectCode = 10457530170432
}

func (s *SuiteTestProcessInstance) Test1ListProcessInstance() {
	res, err := s.client.Project(&s.projectCode).ProcessInstance(nil, nil).List(
		context.Background(),
		instanceoption.WithPageNo(1),
		instanceoption.WithPageSize(10),
		instanceoption.WithProcessDefineCode("18033190883488"))
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func (s *SuiteTestProcessInstance) Test2GetProcessInstance() {
	var processCode int64 = 10777374199936
	var instanceId int64 = 4177948
	res, err := s.client.Project(&s.projectCode).ProcessInstance(&processCode, &instanceId).Get(context.Background())
	s.Nil(err)
	if s.NotNil(res) {
		fmt.Println(res.ToJsonStringPretty())
	}
}

func TestSuiteProcessInstance(t *testing.T) {
	suite.Run(t, new(SuiteTestProcessInstance))
}
