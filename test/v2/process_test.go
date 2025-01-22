package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
	"github.com/stretchr/testify/suite"
)

type SuiteTestProcessDefinition struct {
	suite.Suite
	client           *ds.ClientV2
	filepath         string
	projectCode      int64
	exportProcessDef *typesv2.ExportProcessDef
	tenantMap        map[int64]string
}

func (s *SuiteTestProcessDefinition) SetupSuite() {
	s.client = ds.NewClientV2(
		// option.WithDebug(true),
		option.WithBaseUrl("http://<dolphinscheduler_host>/dolphinscheduler"),
		option.WithToken("<your_token>"),
		option.WithAdminToken("<your_admin_token>"))
	s.filepath = "/root/workflow.json"
	s.projectCode = 10457530170432
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

// func (s *SuiteTestProcessDefinition) Test3ListTenant() {
// 	res, err := s.client.Tenant().List(context.Background(), option.WithPageNo(1), option.WithPageSize(500))
// 	s.Nil(err)
// 	for _, t := range res.TotalList {
// 		s.tenantMap[t.Id] = t.TenantCode
// 	}
// }

// func (s *SuiteTestProcessDefinition) Test4ModifyProcessDefinition() {
// 	for _, def := range *s.exportProcessDef {
// 		// search processDefinition
// 		processDefinitionRes, err := s.client.Project(&s.projectCode).
// 			ProcessDefinition(nil).
// 			List(
// 				context.Background(),
// 				option.WithPageNo(1),
// 				option.WithPageSize(1),
// 				option.WithSearchVal(fmt.Sprintf("%v_import_%s", def.ProcessDefinition.Name, carbon.Now().Format("Ymd"))))
// 		s.Nil(err)
// 		if s.NotNil(processDefinitionRes) {
// 			if s.Equal(1, processDefinitionRes.Total) { // only find one
// 				processDefinition := processDefinitionRes.TotalList[0]
// 				res, err := s.client.Project(&s.projectCode).
// 					ProcessDefinition(&processDefinition.Code).
// 					Get(context.Background())
// 				s.Nil(err)
// 				if s.NotNil(res) {
// 					fmt.Println(processDefinition.Name)
// 					// modify processdefine & taskdefine to remote _import_*
// 					for _, taskDefine := range res.TaskDefinitionList {
// 						re := regexp.MustCompile(fmt.Sprintf(`(.+)_import_%s(\d+)`, carbon.Now().Format("Ymd")))
// 						arr := re.FindStringSubmatch(taskDefine.Name)
// 						if len(arr) > 0 {
// 							taskDefine.Name = arr[1]
// 						}
// 					}
// 					taskDefinitionJson, _ := json.Marshal(res.TaskDefinitionList)
// 					taskRelationJson, _ := json.Marshal(res.ProcessTaskRelationList)
// 					tenantCode, ok := s.tenantMap[def.ProcessDefinition.TenantId]
// 					s.Equal(true, ok)
// 					err = s.client.Project(&s.projectCode).
// 						ProcessDefinition(&processDefinition.Code).
// 						Modify(
// 							context.Background(),
// 							option.WithGlobalParams(res.ProcessDefinition.GlobalParams),
// 							option.WithProcessName(def.ProcessDefinition.Name),
// 							option.WithLocations(res.ProcessDefinition.Locations),
// 							option.WithTaskDefinitionJson(string(taskDefinitionJson)),
// 							option.WithTaskRelationJson(string(taskRelationJson)),
// 							option.WithTenantCode(tenantCode))
// 					s.Nil(err)
// 					// if has schedule, modify WarningGroupId=1
// 					if def.Schedule != nil {
// 						def.Schedule.WarningGroupId = 1
// 						scheduleJson, _ := json.Marshal(def.Schedule)
// 						err = s.client.Project(&s.projectCode).
// 							ProcessDefinition(&processDefinition.Code).
// 							ModifyBasicInfo(
// 								context.Background(),
// 								option.WithProcessName(def.ProcessDefinition.Name),
// 								option.WithTenantCode(tenantCode),
// 								option.WithSchedule(string(scheduleJson)),
// 								option.WithGlobalParams(res.ProcessDefinition.GlobalParams))
// 						s.Nil(err)
// 						scheduleRes, err := s.client.Project(&s.projectCode).
// 							ProcessDefinition(&processDefinition.Code).
// 							Schedule().Get(context.Background())
// 						s.Nil(err)
// 						s.Equal(def.Schedule.WarningGroupId, scheduleRes.WarningGroupId) // test if WarningGroupId modified success
// 						// process online
// 						err = s.client.Project(&s.projectCode).
// 							ProcessDefinition(&processDefinition.Code).
// 							Release(context.Background(), def.ProcessDefinition.Name, constant.PROCESS_RELEASE_STATE_ONLINE)
// 						s.Nil(err)
// 						// schedule online
// 						err = s.client.Project(&s.projectCode).
// 							ProcessDefinition(&processDefinition.Code).
// 							Schedule().Online(context.Background())
// 						s.Nil(err)
// 						// delete schedule
// 						err = s.client.Project(&s.projectCode).
// 							ProcessDefinition(&processDefinition.Code).
// 							Schedule().Delete(context.Background())
// 						s.Equal(true, errorx.IsSchedulerOnlineError(err))
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func (s *SuiteTestProcessDefinition) Test5DeleteProcessDefinition() {
// 	for _, def := range *s.exportProcessDef {
// 		// get processDefinition by name
// 		res, err := s.client.Project(&s.projectCode).ProcessDefinition(nil).GetByName(context.Background(), def.ProcessDefinition.Name)
// 		s.Nil(err)
// 		if s.NotNil(res) {
// 			// delete
// 			fmt.Printf("%s will be delete\n", def.ProcessDefinition.Name)
// 			err = s.client.Project(&s.projectCode).
// 				ProcessDefinition(&res.ProcessDefinition.Code).
// 				Delete(context.Background())
// 			s.Nil(err)
// 		}
// 	}
// }

func TestSuiteProcessDefinition(t *testing.T) {
	suite.Run(t, new(SuiteTestProcessDefinition))
}
