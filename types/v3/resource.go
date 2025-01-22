package v3

import (
	"encoding/json"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ResourceFile struct {
	Id          int64            `json:"id"`
	Pid         string           `json:"pid"`
	Type        string           `json:"type"`
	Name        string           `json:"name"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Directory   bool             `json:"dirctory"`
	IdValue     string           `json:"idValue"`
	Children    ResourceFileList `json:"children"`
}

type ListResourceFileResponse struct {
	types.CommonResponse
	Data ResourceFileList `json:"data"`
}

type ResourceFileList []ResourceFile

func (d *ResourceFileList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *ResourceFileList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
