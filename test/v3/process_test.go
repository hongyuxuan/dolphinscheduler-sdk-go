package test_v3

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang-module/carbon"
	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	instanceoption "github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option/instance"
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
	s.projectCode = 129136205014144
	s.tenantMap = make(map[int64]string)
}

func (s *SuiteTestProcessDefinition) Test1ParseJsonFile() {
	var err error
	s.exportProcessDef, err = s.client.Project(nil).ProcessDefinition(nil).ParseJsonFile(s.filepath)
	s.Nil(err)
	for _, processDef := range *s.exportProcessDef {
		for _, taskDef := range processDef.TaskDefinitionList {
			if taskDef.TaskType == "CONDITIONS" {
				b, _ := json.Marshal(taskDef.TaskParams["dependence"])
				var dependence typesv3.Dependence
				err := json.Unmarshal(b, &dependence)
				s.Nil(err)
				fmt.Printf("%+v\n", dependence)
			}
		}
	}
}

func (s *SuiteTestProcessDefinition) Test2ImportBytes() {
	for _, def := range *s.exportProcessDef {
		def.ProcessDefinition.ReleaseState = constant.PROCESS_RELEASE_STATE_OFFLINE
	}
	b, _ := json.Marshal(&s.exportProcessDef)
	err := s.client.Project(&s.projectCode).ProcessDefinition(nil).ImportBytes(context.Background(), b)
	s.Nil(err)
}

func (s *SuiteTestProcessDefinition) Test3ModifyProcessDefinition() {
	for _, def := range *s.exportProcessDef {
		// search processDefinition
		processDefinitionRes, err := s.client.Project(&s.projectCode).
			ProcessDefinition(nil).
			List(
				context.Background(),
				option.WithPageNo(1),
				option.WithPageSize(1),
				option.WithSearchVal(fmt.Sprintf("%v_import_%s", def.ProcessDefinition.Name, carbon.Now().Format("Ymd"))))
		s.Nil(err)
		if s.NotNil(processDefinitionRes) {
			if s.Equal(1, processDefinitionRes.Total) { // only find one
				processDefinition := processDefinitionRes.TotalList[0]
				res, err := s.client.Project(&s.projectCode).
					ProcessDefinition(&processDefinition.Code).
					Get(context.Background())
				s.Nil(err)
				if s.NotNil(res) {
					fmt.Println(res.ProcessDefinition.Name, res.ProcessDefinition.Code)
					// modify processdefine to remove _import_*
					for _, taskDef := range res.TaskDefinitionList {
						taskDef.EnvironmentCode = 128678357478624
					}
					taskDefinitionJson, _ := json.Marshal(res.TaskDefinitionList)
					taskRelationJson, _ := json.Marshal(res.ProcessTaskRelationList)
					err = s.client.Project(&s.projectCode).
						ProcessDefinition(&processDefinition.Code).
						Modify(
							context.Background(),
							option.WithGlobalParams(res.ProcessDefinition.GlobalParams),
							option.WithProcessName(def.ProcessDefinition.Name),
							option.WithLocations(res.ProcessDefinition.Locations),
							option.WithTaskDefinitionJson(string(taskDefinitionJson)),
							option.WithTaskRelationJson(string(taskRelationJson)))
					s.Nil(err)
					// process online
					err = s.client.Project(&s.projectCode).
						ProcessDefinition(&processDefinition.Code).
						Release(context.Background(), def.ProcessDefinition.Name, constant.PROCESS_RELEASE_STATE_ONLINE)
					s.Nil(err)
					// schedule online
					err = s.client.Project(&s.projectCode).
						ProcessDefinition(&processDefinition.Code).
						Schedule().Online(context.Background())
					s.Nil(err)
					// delete schedule
					err = s.client.Project(&s.projectCode).
						ProcessDefinition(&processDefinition.Code).
						Schedule().Delete(context.Background())
					s.Equal(true, errorx.IsSchedulerOnlineError(err))
				}
			}
		}
	}
}

func (s *SuiteTestProcessDefinition) Test4Export() {
	res, err := s.client.Project(&s.projectCode).ProcessDefinition(nil).BatchExport(context.Background(), "129855618575584")
	s.Nil(err)
	fmt.Println(string(res))
}

func (s *SuiteTestProcessDefinition) Test6StartInstance() {
	var processCode int64 = 129136590590176

	err := s.client.Project(&s.projectCode).ProcessDefinition(nil).
		StartInstance(
			context.Background(),
			instanceoption.WithProcessDefinitionCode(processCode),
			instanceoption.WithScheduleTime(""),
			instanceoption.WithFailureStrategy("END"),
			instanceoption.WithWarningType("NONE"),
			instanceoption.WithWarningGroupId(0),
			instanceoption.WithExecType(""),
			instanceoption.WithStartNodeList(""),
			instanceoption.WithTaskDependType("TASK_POST"),
			instanceoption.WithRunMode("RUN_MODE_SERIAL"),
			instanceoption.WithInstancePriority("MEDIUM"),
			instanceoption.WithWorkerGroup("default"),
			instanceoption.WithTenantCode("default"),
			instanceoption.WithEnvironmentCode("123636042396032"),
			instanceoption.WithStartParams(""),
			instanceoption.WithParallelism(""),
			instanceoption.WithDryRun("0"))
	s.Nil(err)
}

func TestSuiteProcessDefinition(t *testing.T) {
	suite.Run(t, new(SuiteTestProcessDefinition))
}
