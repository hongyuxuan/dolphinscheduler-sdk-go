package types

import (
	"encoding/json"
)

type Project struct {
	Id               int64  `json:"id"`
	Code             int64  `json:"code"`
	Name             string `json:"name"`
	DefCount         int64  `json:"defCount"`
	Description      string `json:"description"`
	InstRunningCount int64  `json:"instRunningCount"`
	Perm             int64  `json:"perm"`
	UserName         string `json:"userName"`
	UserId           int64  `json:"userId"`
	CreateTime       string `json:"createTime"`
	UpdateTime       string `json:"updateTime"`
}

type ListProjectResponse struct {
	CommonResponse
	Data ProjectList `json:"data"`
}

type ProjectList struct {
	Total     int       `json:"total"`
	TotalList []Project `json:"totalList"`
}

func (d *ProjectList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *ProjectList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
