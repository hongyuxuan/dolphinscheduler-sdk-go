package v2

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
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

func (p *Resource) ListFile(ctx context.Context) (resp *typesv2.ResourceFileList, err error) {
	option := make(types.ListOption)
	option["type"] = "FILE"
	var res *typesv2.ListResourceFileResponse
	if err = p.httpclient.Get("/resources/list").SetQueryParams(option).SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}

func (p *Resource) ListFileFlatMap(ctx context.Context) (resp map[int64]string, err error) {
	var res *typesv2.ResourceFileList
	if res, err = p.ListFile(ctx); err != nil {
		return
	}
	resp = make(map[int64]string)
	addResourceFileFlatMap(*res, resp)
	return
}

func addResourceFileFlatMap(children typesv2.ResourceFileList, resourceIdNameMap map[int64]string) {
	for _, resource := range children {
		if !resource.Directory {
			resourceIdNameMap[resource.Id] = resource.FullName
		} else {
			addResourceFileFlatMap(resource.Children, resourceIdNameMap)
		}
	}
}
