package v3

import (
	"context"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	"github.com/imroc/req/v3"
)

type Datasource struct {
	httpclient *req.Client
	config     *config.Config
}

func New(c *config.Config) *Datasource {
	return &Datasource{
		httpclient: c.Httpclient,
		config:     c,
	}
}

func (w *Datasource) List(ctx context.Context, opts ...option.ListOptionFunc) (resp *types.DatasourceList, err error) {
	option := make(types.ListOption)
	for _, opt := range opts {
		opt(&option)
	}
	var res *types.ListDatasourceResponse
	if err = w.httpclient.Get("/datasources").
		SetQueryParams(option).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	return &res.Data, nil
}
