package types

type ListOption map[string]string

type ListResponseData struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Start       int `json:"start"`
	Total       int `json:"total"`
	TotalPage   int `json:"totalPage"`
}

type CommonResponse struct {
	Code    int64  `json:"code"`
	Failed  bool   `json:"failed"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type ProcessDefinitionOption map[string]string
