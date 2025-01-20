package v3

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
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

func (p *Resource) ListFile(ctx context.Context) (resp *typesv3.ResourceFileList, err error) {
	option := make(types.ListOption)
	option["type"] = "FILE"
	option["fullName"] = ""
	var res *typesv3.ListResourceFileResponse
	if err = p.httpclient.Get("/resources/list").SetQueryParams(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}
