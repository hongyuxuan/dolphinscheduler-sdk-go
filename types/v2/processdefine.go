package v2

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ExportProcessDef []*ProcessDef

type ProcessDef struct {
	ProcessDefinition       *ProcessDefinition `json:"processDefinition"`
	ProcessTaskRelationList []*TaskRelation    `json:"processTaskRelationList"`
	TaskDefinitionList      []*TaskDefinition  `json:"taskDefinitionList"`
	Schedule                *Schedule          `json:"schedule"`
}

type ProcessDefinition struct {
	Id                   int64             `json:"id"`
	Code                 int64             `json:"code"`
	Name                 string            `json:"name"`
	Version              int64             `json:"version"`
	ReleaseState         string            `json:"releaseState"`
	ProjectCode          int64             `json:"projectCode"`
	Description          string            `json:"description"`
	GlobalParams         string            `json:"globalParams"`
	GlobalParamList      []Params          `json:"globalParamList"`
	GlobalParamMap       map[string]string `json:"globalParamMap"`
	CreateTime           string            `json:"createTime"`
	UpdateTime           string            `json:"updateTime"`
	Flag                 string            `json:"flag"`
	UserId               int64             `json:"userId"`
	UserName             *string           `json:"userName"`
	ProjectName          *string           `json:"projectName"`
	Locations            *string           `json:"locations"`
	ScheduleReleaseState *string           `json:"scheduleReleaseState"`
	Timeout              int64             `json:"timeout"`
	TenantId             int64             `json:"tenantId"`
	TenantCode           *string           `json:"tenantCode"`
	ModifyBy             *string           `json:"modifyBy"`
	WarningGroupId       int64             `json:"warningGroupId"`
}

type Params struct {
	Prop   string `json:"prop"`
	Direct string `json:"direct"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

type TaskRelation struct {
	Id                       int64  `json:"id"`
	Name                     string `json:"name"`
	ProcessDefinitionVersion int64  `json:"processDefinitionVersion"`
	ProjectCode              int64  `json:"projectCode"`
	ProcessDefinitionCode    int64  `json:"processDefinitionCode"`
	PreTaskCode              int64  `json:"preTaskCode"`
	PreTaskVersion           int64  `json:"preTaskVersion"`
	PostTaskCode             int64  `json:"postTaskCode"`
	PostTaskVersion          int64  `json:"postTaskVersion"`
	ConditionType            string `json:"conditionType"`
	ConditionParams          map[string]string
	CreateTime               string `json:"createTime"`
	UpdateTime               string `json:"updateTime"`
	Operator                 int64  `json:"operator"`
	OperateTime              string `json:"operateTime"`
}

type TaskDefinition struct {
	Id                    int64             `json:"id"`
	Code                  int64             `json:"code"`
	Name                  string            `json:"name"`
	Version               int64             `json:"version"`
	Description           string            `json:"description"`
	ProjectCode           int64             `json:"projectCode"`
	UserId                int64             `json:"userId"`
	TaskType              string            `json:"taskType"`
	TaskParams            *TaskParams       `json:"taskParams"`
	TaskParamList         []Params          `json:"taskParamList"`
	TaskParamMap          map[string]string `json:"taskParamMap"`
	Flag                  string            `json:"flag"`
	TaskPriority          string            `json:"taskPriority"`
	UserName              *string           `json:"userName"`
	ProjectName           *string           `json:"projectName"`
	WorkerGroup           string            `json:"workerGroup"`
	EnvironmentCode       int64             `json:"environmentCode"`
	FailRetryTimes        int64             `json:"failRetryTimes"`
	FailRetryInterval     int64             `json:"failRetryInterval"`
	TimeoutFlag           string            `json:"timeoutFlag"`
	TimeoutNotifyStrategy string            `json:"timeoutNotifyStrategy"`
	Timeout               int64             `json:"timeout"`
	DelayTime             int64             `json:"delayTime"`
	ResourceIds           string            `json:"resourceIds"`
	CreateTime            string            `json:"createTime"`
	UpdateTime            string            `json:"updateTime"`
	ModifyBy              *string           `json:"modifyBy"`
	Operator              int64             `json:"operator"`
	OperateTime           string            `json:"operateTime"`
}

type TaskParams struct {
	ProcessDefinitionCode int64 `json:"processDefinitionCode"`
	ResourceList          []struct {
		Id int64 `json:"id"`
	} `json:"resourceList,omitempty"`
	LocalParams        []Params          `json:"localParams"`
	HttpParams         map[string]string `json:"httpParams,omitempty"`
	Url                string            `json:"url,omitempty"`
	RawScript          string            `json:"rawScript,omitempty"`
	HttpMethod         string            `json:"httpMethod,omitempty"`
	HttpCheckCondition string            `json:"httpCheckCondition,omitempty"`
	Condition          string            `json:"condition,omitempty"`
	ConnectTimeout     int64             `json:"connectTimeout,omitempty"`
	SocketTimeout      int64             `json:"socketTimeout,omitempty"`
	Dependence         struct {
		DependTaskList []*DependTask `json:"dependTaskList"`
		Relation       string        `json:"relation"`
	} `json:"dependence"`
	ConditionResult struct {
		SuccessNode []map[string]interface{} `json:"successNode"`
		FailedNode  []map[string]interface{} `json:"failedNode"`
	} `json:"conditionResult"`
	WaitStartTimeout map[string]interface{} `json:"waitStartTimeout"`
	SwitchResult     *SwitchResult          `json:"switchResult"`
}

type HttpParams struct {
	Prop               string `json:"prop"`
	HttpParametersType string `json:"httpParametersType"`
	Value              string `json:"value"`
}

type DependTask struct {
	Condition string `json:"condition"`
	NextNode  int64  `json:"nextNode"`
}

type SwitchResult struct {
	DependTaskList []*DependTask `json:"dependTaskList"`
	NextNode       int64         `json:"nextNode"`
}

type Schedule struct {
	Id                      int64   `json:"id"`
	ProcessDefinitionCode   int64   `json:"processDefinitionCode"`
	ProcessDefinitionName   *string `json:"processDefinitionName"`
	ProjectName             *string `json:"projectName"`
	DefinitionDescription   *string `json:"definitionDescription"`
	StartTime               string  `json:"startTime"`
	EndTime                 string  `json:"endTime"`
	TimezoneId              string  `json:"timezoneId"`
	Crontab                 string  `json:"crontab"`
	FailureStrategy         string  `json:"failureStrategy"`
	WarningType             string  `json:"warningType"`
	CreateTime              string  `json:"createTime"`
	UpdateTime              string  `json:"updateTime"`
	UserId                  int64   `json:"userId"`
	UserName                *string `json:"userName"`
	ReleaseState            string  `json:"releaseState"`
	WarningGroupId          int64   `json:"warningGroupId"`
	ProcessInstancePriority string  `json:"processInstancePriority"`
	WorkerGroup             string  `json:"workerGroup"`
	EnvironmentCode         int64   `json:"environmentCode"`
}

type ListProcessDefinitionResponse struct {
	types.CommonResponse
	Data ProcessDefinitionList `json:"data"`
}

type ProcessDefinitionList struct {
	Total     int                 `json:"total"`
	TotalList []ProcessDefinition `json:"totalList"`
}

type ProcessDefinitionResponse struct {
	types.CommonResponse
	Data ProcessDef `json:"data"`
}

type ListScheduleResponse struct {
	types.CommonResponse
	Data ScheduleList `json:"data"`
}

type ScheduleList struct {
	Total     int        `json:"total"`
	TotalList []Schedule `json:"totalList"`
}

func (e ExportProcessDef) ToJsonString() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *ExportProcessDef) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(e, "", "  ")
	return string(b)
}

func (p ProcessDefinition) ToJsonString() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (p *ProcessDefinition) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}

func (p ProcessDef) ToJsonString() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (p *ProcessDef) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}

type ProcessDefinitionOption map[string]string
