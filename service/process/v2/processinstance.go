package v2

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	instanceoption "github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option/instance"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
	"github.com/imroc/req/v3"
)

type ProcessInstance struct {
	httpclient  *req.Client
	projectCode *int64
	processCode *int64
	instanceId  *int64
	config      *config.Config
}

func NewProcessInstance(c *config.Config, projectCode, processCode, instanceId *int64) *ProcessInstance {
	return &ProcessInstance{
		httpclient:  c.Httpclient,
		projectCode: projectCode,
		processCode: processCode,
		instanceId:  instanceId,
		config:      c,
	}
}

func (p *ProcessInstance) List(ctx context.Context, opts ...instanceoption.InstanceListOptionFunc) (resp *typesv2.ProcessInstanceList, err error) {
	option := make(types.InstanceListOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *typesv2.ListProcessInstanceResponse
	if err = p.httpclient.Get(fmt.Sprintf("/projects/%d/process-instances", *p.projectCode)).
		SetQueryParams(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (p *ProcessInstance) Get(ctx context.Context) (resp *typesv2.ProcessInstance, err error) {
	var res *typesv2.ProcessInstanceResponse
	if err = p.httpclient.Get(fmt.Sprintf("/projects/%d/process-instances/%d", *p.projectCode, *p.instanceId)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return res.Data, nil
}
