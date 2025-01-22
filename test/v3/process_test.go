package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	typesv3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v3"
	"github.com/stretchr/testify/suite"
)

type SuiteTestProcessDefinition struct {
	suite.Suite
	client           *ds.ClientV3
	filepath         string
	projectCode      int64
	exportProcessDef *typesv3.ExportProcessDef
	tenantMap        map[int64]string
}

func (s *SuiteTestProcessDefinition) SetupSuite() {
	s.client = ds.NewClientV3(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"))
	s.filepath = "/root/workflow.json"
	s.projectCode = 128869683708960
	s.tenantMap = make(map[int64]string)
}

func (s *SuiteTestProcessDefinition) Test1ParseJsonFile() {
	var err error
	s.exportProcessDef, err = s.client.Project(nil).ProcessDefinition(nil).ParseJsonFile(s.filepath)
	fmt.Println(s.exportProcessDef.ToJsonStringPretty())
	s.Nil(err)
}

func (s *SuiteTestProcessDefinition) Test2ImportBytes() {
	for _, def := range *s.exportProcessDef {
		def.ProcessDefinition.ReleaseState = constant.PROCESS_RELEASE_STATE_OFFLINE
	}
	b, _ := json.Marshal(&s.exportProcessDef)
	err := s.client.Project(&s.projectCode).ProcessDefinition(nil).ImportBytes(context.Background(), b)
	s.Nil(err)
}

func TestSuiteProcessDefinition(t *testing.T) {
	suite.Run(t, new(SuiteTestProcessDefinition))
}
