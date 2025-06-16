package types

import "encoding/json"

type Datasource struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	ConnectionParams string `json:"connectionParams"`
	Note             string `json:"note"`
	Type             string `json:"type"`
	UserId           int64  `json:"userId"`
	UserName         string `json:"userName"`
	CreateTime       string `json:"createTime"`
	UpdateTime       string `json:"updateTime"`
}

type ListDatasourceResponse struct {
	CommonResponse
	Data DatasourceList `json:"data"`
}

type DatasourceList struct {
	Total     int          `json:"total"`
	TotalList []Datasource `json:"totalList"`
}

func (d *DatasourceList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *DatasourceList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
