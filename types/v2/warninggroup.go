package v2

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type WarningGroup struct {
	Id               int64  `json:"id"`
	GroupName        string `json:"groupName"`
	AlertInstanceIds string `json:"alertInstanceIds"`
	Description      string `json:"description"`
	CreateTime       string `json:"createTime"`
	UpdateTime       string `json:"updateTime"`
	CreateUserId     int64  `json:"createUserId"`
}

type ListWarningGroupResponse struct {
	types.CommonResponse
	Data WarningGroupList `json:"data"`
}

type WarningGroupList struct {
	Total     int            `json:"total"`
	TotalList []WarningGroup `json:"totalList"`
}

func (d *WarningGroupList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *WarningGroupList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
