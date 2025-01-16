package errorx

import (
	"fmt"
	"net/http"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/core/constant"
)

type DsError struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *DsError) Error() string {
	return e.Message
}

func NewError(code int64, message string, data interface{}) error {
	return &DsError{Code: code, Message: message, Data: data}
}

func NewDefaultError(message string, a ...any) error {
	return &DsError{Code: http.StatusInternalServerError, Message: fmt.Sprintf(message, a...)}
}

func IsNoSchedulerError(err error) bool {
	defer func() {
		if er := recover(); er != nil {
			return
		}
	}()
	e := err.(*DsError)
	return e.Code == constant.ERR_NO_SCHEDULER
}

func IsSchedulerOnlineError(err error) bool {
	defer func() {
		if er := recover(); er != nil {
			return
		}
	}()
	e := err.(*DsError)
	return e.Code == constant.ERR_SCHEDULER_ONLINE
}
