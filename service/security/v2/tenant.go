package v2

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
	"github.com/imroc/req/v3"
)

type Tenant struct {
	httpclient *req.Client
	config     *config.Config
}

func NewTenant(c *config.Config) *Tenant {
	return &Tenant{
		httpclient: c.Httpclient,
		config:     c,
	}
}

func (t *Tenant) List(ctx context.Context, opts ...option.ListOptionFunc) (resp *typesv2.TenantList, err error) {
	option := make(types.ListOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *typesv2.ListTenantResponse
	if err = t.httpclient.Get("/tenants").
		SetHeader("token", t.config.AdminToken).
		SetQueryParams(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}
