package v2

import (
	"context"
	"fmt"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/config"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/errorx"
	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
	"github.com/imroc/req/v3"
)

type Schedule struct {
	httpclient  *req.Client
	projectCode *int64
	processCode *int64
	config      *config.Config
}

func NewSchedule(c *config.Config, projectCode, processCode *int64) *Schedule {
	return &Schedule{
		httpclient:  c.Httpclient,
		projectCode: projectCode,
		processCode: processCode,
		config:      c,
	}
}

func (s *Schedule) Get(ctx context.Context) (resp *typesv2.Schedule, err error) {
	var res *typesv2.ListScheduleResponse
	if err = s.httpclient.Get(fmt.Sprintf("/projects/%d/schedules", *s.projectCode)).
		SetQueryParams(map[string]string{
			"processDefinitionCode": fmt.Sprintf("%d", *s.processCode),
			"pageNo":                "1",
			"pageSize":              "1",
			"searchVal":             "",
		}).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return nil, errorx.NewError(res.Code, res.Msg, res.Data)
	}
	if res.Data.Total == 0 {
		return nil, errorx.NewError(constant.ERR_NO_SCHEDULER, fmt.Sprintf("Cannot find scheduler for processDefinitionCode=%d", *s.processCode), nil)
	}
	return &res.Data.TotalList[0], nil
}

func (s *Schedule) Online(ctx context.Context) (err error) {
	var schedule *typesv2.Schedule
	if schedule, err = s.Get(ctx); err != nil {
		return
	}
	var res *types.CommonResponse
	if err = s.httpclient.Post(fmt.Sprintf("/projects/%d/schedules/%d/online", *s.projectCode, schedule.Id)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

func (s *Schedule) Offline(ctx context.Context) (err error) {
	var schedule *typesv2.Schedule
	if schedule, err = s.Get(ctx); err != nil {
		return
	}
	var res *types.CommonResponse
	if err = s.httpclient.Post(fmt.Sprintf("/projects/%d/schedules/%d/offline", *s.projectCode, schedule.Id)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}

func (s *Schedule) Delete(ctx context.Context) (err error) {
	var schedule *typesv2.Schedule
	if schedule, err = s.Get(ctx); err != nil {
		return
	}
	var res *types.CommonResponse
	if err = s.httpclient.Delete(fmt.Sprintf("/projects/%d/schedules/%d?scheduleId=%d", *s.projectCode, schedule.Id, schedule.Id)).
		SetSuccessResult(&res).Do(ctx).Err; err != nil {
		return
	}
	if res.Code != constant.ERR_OK {
		return errorx.NewError(res.Code, res.Msg, nil)
	}
	return nil
}
