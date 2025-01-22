package v2

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	processv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/process/v2"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Project struct {
	httpclient  *req.Client
	projectCode *int64
	config      *config.Config
}

func New(c *config.Config, projectCode *int64) *Project {
	return &Project{
		httpclient:  c.Httpclient,
		projectCode: projectCode,
		config:      c,
	}
}

func (p *Project) List(ctx context.Context, opts ...option.ListOptionFunc) (resp *types.ProjectList, err error) {
	option := make(types.ListOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *types.ListProjectResponse
	if err = p.httpclient.Get("/projects").SetQueryParams(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (p *Project) ProcessDefinition(processCode *int64) *processv2.ProcessDefinition {
	return processv2.NewProcessDefinition(p.config, p.projectCode, processCode)
}

func (p *Project) ProcessInstance(processCode *int64) *processv2.ProcessInstance {
	return processv2.NewProcessInstance(p.config, p.projectCode, processCode)
}
