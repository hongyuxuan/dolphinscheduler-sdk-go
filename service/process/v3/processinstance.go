package v3

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/imroc/req/v3"
)

type ProcessInstance struct {
	httpclient   *req.Client
	projectCode  *int64
	instanceCode *int64
	config       *config.Config
}

func NewProcessInstance(c *config.Config, projectCode, instanceCode *int64) *ProcessInstance {
	return &ProcessInstance{
		httpclient:   c.Httpclient,
		projectCode:  projectCode,
		instanceCode: instanceCode,
		config:       c,
	}
}

func (p *ProcessInstance) List(ctx context.Context, opts ...option.ListOptionFunc) (err error) {
	// TODO
	return
}
