package option

import (
	"fmt"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type InstanceListOptionFunc func(*types.InstanceListOption)

func WithPageNo(pageNo int) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["pageNo"] = fmt.Sprintf("%d", pageNo)
	}
}

func WithPageSize(pageSize int) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["pageSize"] = fmt.Sprintf("%d", pageSize)
	}
}

func WithSearchVal(searchVal string) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["searchVal"] = searchVal
	}
}

func WithProcessDefineCode(processDefineCode string) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["processDefineCode"] = processDefineCode
	}
}

func WithStateType(stateType string) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["stateType"] = stateType
	}
}

func WithStartDate(startDate string) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["startDate"] = startDate
	}
}

func WithEndDate(endDate string) InstanceListOptionFunc {
	return func(l *types.InstanceListOption) {
		(*l)["endDate"] = endDate
	}
}
