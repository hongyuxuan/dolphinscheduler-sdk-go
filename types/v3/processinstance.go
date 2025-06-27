package v3

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ProcessInstance struct {
	Id                       int64        `json:"id"`
	Blocked                  bool         `json:"blocked"`
	CmdTypeIfComplement      *string      `json:"cmdTypeIfComplement"`
	CommandParam             *string      `json:"commandParam"`
	CommandStartTime         *string      `json:"commandStartTime"`
	CommandType              *string      `json:"commandType"`
	ComplementData           bool         `json:"complementData"`
	DagData                  **ProcessDef `json:"dagData"`
	DependenceScheduleTimes  *string      `json:"dependenceScheduleTimes"`
	DryRun                   int64        `json:"dryRun"`
	Duration                 *string      `json:"duration"`
	EndTime                  *string      `json:"endTime"`
	EnvironmentCode          *int64       `json:"environmentCode"`
	ExecutorId               *int64       `json:"executorId"`
	ExecutorName             *string      `json:"executorName"`
	FailureStrategy          *string      `json:"failureStrategy"`
	GlobalParams             *string      `json:"globalParams"`
	HistoryCmd               *string      `json:"historyCmd"`
	Host                     *string      `json:"host"`
	IsSubProcess             *string      `json:"isSubProcess"`
	Locations                *string      `json:"locations"`
	MaxTryTimes              int64        `json:"maxTryTimes"`
	Name                     *string      `json:"name"`
	NextProcessInstanceId    int64        `json:"nextProcessInstanceId"`
	ProcessDefinition        *string      `json:"processDefinition"`
	ProcessDefinitionCode    int64        `json:"processDefinitionCode"`
	ProcessDefinitionVersion int64        `json:"processDefinitionVersion"`
	ProcessInstancePriority  *string      `json:"processInstancePriority"`
	ProjectCode              int64        `json:"projectCode"`
	Queue                    interface{}  `json:"queue"`
	Recovery                 *string      `json:"recovery"`
	RestartTime              *string      `json:"restartTime"`
	RunTimes                 int64        `json:"runTimes"`
	ScheduleTime             *string      `json:"scheduleTime"`
	StartTime                *string      `json:"startTime"`
	State                    *string      `json:"state"`
	StateDescList            *string      `json:"stateDescList"`
	StateHistory             *string      `json:"stateHistory"`
	TaskDependType           *string      `json:"taskDependType"`
	TenantCode               *string      `json:"tenantCode"`
	TestFlag                 int64        `json:"testFlag"`
	Timeout                  int64        `json:"timeout"`
	VarPool                  *string      `json:"varPool"`
	WarningGroupId           *int64       `json:"warningGroupId"`
	WarningType              *string      `json:"warningType"`
	WorkerGroup              *string      `json:"workerGroup"`
}

type ListProcessInstanceResponse struct {
	types.CommonResponse
	Data ProcessInstanceList `json:"data"`
}

type ProcessInstanceList struct {
	Total     int               `json:"total"`
	TotalList []ProcessInstance `json:"totalList"`
}

func (p *ProcessInstanceList) ToJsonString() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (p *ProcessInstanceList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}

type ProcessInstanceResponse struct {
	types.CommonResponse
	Data *ProcessInstance `json:"data"`
}

func (p *ProcessInstance) ToJsonString() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (p *ProcessInstance) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}
