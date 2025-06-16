package v3

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	resourceoption "github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option/resource"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v3"
	"github.com/imroc/req/v3"
)

type Resource struct {
	httpclient *req.Client
	config     *config.Config
}

func New(c *config.Config) *Resource {
	return &Resource{
		httpclient: c.Httpclient,
		config:     c,
	}
}

func (r *Resource) ListFile(ctx context.Context) (resp *typesv3.ResourceFileList, err error) {
	option := make(types.ListOption)
	option["type"] = "FILE"
	option["fullName"] = ""
	var res *typesv3.ListResourceFileResponse
	if err = r.httpclient.Get("/resources/list").SetQueryParams(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (r *Resource) ViewFile(ctx context.Context, opts ...resourceoption.ResourceOptionFunc) (resp *typesv3.FileData, err error) {
	option := make(types.ResourceOption)
	for _, opt := range opts {
		opt(&option)
	}
	option["skipLineNum"] = "0"
	var res *typesv3.ViewFileResponse
	if err = r.httpclient.Get("/resources/view").SetQueryParams(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (r *Resource) ModifyFile(ctx context.Context, opts ...resourceoption.ResourceOptionFunc) (resp *types.CommonResponse, err error) {
	option := make(types.ResourceOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *types.CommonResponse
	if err = r.httpclient.Put("/resources/update-content").SetFormData(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, nil)
	}
	return
}
