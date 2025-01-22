package dolphinscheduler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"
	datasourcev2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/datasource/v2"
	projectv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/project/v2"
	resourcev2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/resource/v2"
	securityv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/security/v2"
	"github.com/imroc/req/v3"
)

type ClientV2 struct {
	Config *config.Config
}

func NewClientV2(opts ...option.ClientOptionFunc) *ClientV2 {
	config := &config.Config{
		Header: make(map[string]string),
	}
	for _, opt := range opts {
		opt(config)
	}

	httpclient := req.C().
		OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			if req.RetryAttempt > 0 {
				return nil
			}
			req.EnableDump()
			return nil
		}).
		OnAfterResponse(func(client *req.Client, res *req.Response) (err error) {
			if res.Err != nil {
				return errorx.NewError(http.StatusInternalServerError, res.Err.Error(), nil)
			}
			responseCode := strconv.Itoa(res.StatusCode)
			if !strings.HasPrefix(responseCode, "2") && !strings.HasPrefix(responseCode, "3") {
				defer func() {
					if e := recover(); e != nil {
						err = errorx.NewError(int64(res.StatusCode), res.String(), nil)
					}
				}()
				ress := make(map[string]interface{})
				res.UnmarshalJson(&ress)
				err = errorx.NewError(int64(res.StatusCode), ress["message"].(string), ress["data"])
			}
			return
		})

	httpclient.SetBaseURL(config.BaseUrl).SetCommonHeaders(config.Header)
	if config.Timeout != 0 {
		httpclient.SetTimeout(time.Duration(config.Timeout) * time.Millisecond)
	}
	if config.EnableDebug {
		httpclient.EnableDebugLog()
		httpclient.EnableDumpAll()
	} else {
		httpclient.DisableDebugLog()
		httpclient.DisableDumpAll()
	}
	config.Httpclient = httpclient

	return &ClientV2{
		Config: config,
	}
}

func (c *ClientV2) Project(projectCode *int64) *projectv2.Project {
	return projectv2.New(c.Config, projectCode)
}

func (c *ClientV2) Resource() *resourcev2.Resource {
	return resourcev2.New(c.Config)
}

func (c *ClientV2) Datasource() *datasourcev2.Datasource {
	return datasourcev2.New(c.Config)
}

func (c *ClientV2) WarningGroup() *securityv2.WarningGroup {
	return securityv2.NewWG(c.Config)
}

func (c *ClientV2) Environment() *securityv2.Environment {
	return securityv2.NewEnv(c.Config)
}

func (c *ClientV2) Tenant() *securityv2.Tenant {
	return securityv2.NewTenant(c.Config)
}
