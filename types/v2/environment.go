package v2

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type Environment struct {
	Id           int64    `json:"id"`
	Code         int64    `json:"code"`
	Name         string   `json:"name"`
	Config       string   `json:"config"`
	Description  string   `json:"description"`
	WorkerGroups []string `json:"workerGroups"`
	CreateTime   string   `json:"createTime"`
	UpdateTime   string   `json:"updateTime"`
	Operator     int64    `json:"operator"`
}

type ListEnvironmentResponse struct {
	types.CommonResponse
	Data EnvironmentList `json:"data"`
}

type EnvironmentList struct {
	Total     int           `json:"total"`
	TotalList []Environment `json:"totalList"`
}

func (d *EnvironmentList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *EnvironmentList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
