package dolphinscheduler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/option"

	datasourcev3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/datasource/v3"
	projectv3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/project/v3"
	resourcev3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/resource/v3"
	securityv3 "github.com/hongyuxuan/dolphinscheduler-sdk-go/service/security/v3"
	"github.com/imroc/req/v3"
)

type ClientV3 struct {
	Config *config.Config
}

func NewClientV3(opts ...option.ClientOptionFunc) *ClientV3 {
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

	return &ClientV3{
		Config: config,
	}
}

func (c *ClientV3) Project(projectCode *int64) *projectv3.Project {
	return projectv3.New(c.Config, projectCode)
}

func (c *ClientV3) Resource() *resourcev3.Resource {
	return resourcev3.New(c.Config)
}

func (c *ClientV3) Datasource() *datasourcev3.Datasource {
	return datasourcev3.New(c.Config)
}

func (c *ClientV3) WarningGroup() *securityv3.WarningGroup {
	return securityv3.NewWG(c.Config)
}

func (c *ClientV3) Environment() *securityv3.Environment {
	return securityv3.NewEnv(c.Config)
}

func (c *ClientV3) Tenant() *securityv3.Tenant {
	return securityv3.NewTenant(c.Config)
}
