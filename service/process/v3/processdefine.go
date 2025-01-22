package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v3"
	"github.com/imroc/req/v3"
)

type ProcessDefinition struct {
	httpclient  *req.Client
	projectCode *int64
	processCode *int64
	config      *config.Config
}

func NewProcessDefinition(c *config.Config, projectCode, processCode *int64) *ProcessDefinition {
	return &ProcessDefinition{
		httpclient:  c.Httpclient,
		projectCode: projectCode,
		processCode: processCode,
		config:      c,
	}
}

func (p *ProcessDefinition) ParseJsonFile(filepath string) (resp *typesv3.ExportProcessDef, err error) {
	f, _ := os.Open(filepath)
	defer f.Close()
	content, err := io.ReadAll(f)
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, errorx.NewDefaultError("error parsing file %s to v3.ExportProcessDef: %v", filepath, err)
	}
	return
}

func (p *ProcessDefinition) ImportBytes(ctx context.Context, jsonb []byte) (err error) {
	var res *types.CommonResponse
	if err = p.httpclient.Post(fmt.Sprintf("/projects/%d/process-definition/import", *p.projectCode)).
		SetFileBytes("file", "uploadfile", jsonb).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

func (p *ProcessDefinition) ImportFile(ctx context.Context, filepath string) (err error) {
	var res *types.CommonResponse
	if err = p.httpclient.Post(fmt.Sprintf("/projects/%d/process-definition/import", *p.projectCode)).
		SetFile("file", filepath).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

func (p *ProcessDefinition) List(ctx context.Context, opts ...option.ListOptionFunc) (resp *typesv3.ProcessDefinitionList, err error) {
	option := make(types.ListOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *typesv3.ListProcessDefinitionResponse
	if err = p.httpclient.Get(fmt.Sprintf("/projects/%d/process-definition", *p.projectCode)).
		SetQueryParams(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (p *ProcessDefinition) Get(ctx context.Context) (resp *typesv3.ProcessDef, err error) {
	var res *typesv3.ProcessDefinitionResponse
	if err = p.httpclient.Get(fmt.Sprintf("/projects/%d/process-definition/%d", *p.projectCode, *p.processCode)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (p *ProcessDefinition) GetByName(ctx context.Context, processName string) (resp *typesv3.ProcessDef, err error) {
	var res *typesv3.ProcessDefinitionResponse
	if err = p.httpclient.Get(fmt.Sprintf("/projects/%d/process-definition/query-by-name", *p.projectCode)).
		SetQueryParam("name", processName).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

// opts:
// option.WithGlobalParams(string),
// option.WithProcessName(string),
// option.WithLocations(*string),
// option.WithTaskDefinitionJson(string),
// option.WithTaskRelationJson(string),
// option.WithTenantCode(*string)
func (p *ProcessDefinition) Modify(ctx context.Context, opts ...option.ProcessDefinitionOptionFunc) (err error) {
	option := make(types.ProcessDefinitionOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *types.CommonResponse
	if err = p.httpclient.Put(fmt.Sprintf("/projects/%d/process-definition/%d", *p.projectCode, *p.processCode)).
		SetFormData(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return
}

// opts:
// option.WithProcessName(string),
// option.WithTenantCode(*string),
// option.WithSchedule(string),
// option.WithGlobalParams(string)
func (p *ProcessDefinition) ModifyBasicInfo(ctx context.Context, opts ...option.ProcessDefinitionOptionFunc) (err error) {
	option := make(types.ProcessDefinitionOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *types.CommonResponse
	if err = p.httpclient.Put(fmt.Sprintf("/projects/%d/process-definition/%d/basic-info", *p.projectCode, *p.processCode)).
		SetFormData(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return
}

func (p *ProcessDefinition) Release(ctx context.Context, processName, releaseState string) (err error) {
	var res *types.CommonResponse
	if err = p.httpclient.Post(fmt.Sprintf("/projects/%d/process-definition/%d/release", *p.projectCode, *p.processCode)).
		SetFormData(map[string]string{
			"name":         processName,
			"releaseState": releaseState,
		}).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

func (p *ProcessDefinition) Delete(ctx context.Context) (err error) {
	var res *types.CommonResponse
	if err = p.httpclient.Delete(fmt.Sprintf("/projects/%d/process-definition/%d", *p.projectCode, *p.processCode)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

// func (p *ProcessDefinition) Schedule() *Schedule {
// 	return NewSchedule(p.config, p.projectCode, p.processCode)
// }
