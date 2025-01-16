package v2

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type Tenant struct {
	Id          int64  `json:"id"`
	TenantCode  string `json:"tenantCode"`
	QueueName   string `json:"queueName"`
	QueueId     int64  `json:"queueId"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type ListTenantResponse struct {
	types.CommonResponse
	Data TenantList `json:"data"`
}

type TenantList struct {
	Total     int      `json:"total"`
	TotalList []Tenant `json:"totalList"`
}

func (d *TenantList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *TenantList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
