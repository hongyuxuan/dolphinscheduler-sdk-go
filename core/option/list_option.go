package option

import (
	"fmt"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ListOptionFunc func(*types.ListOption)

func WithPageNo(pageNo int) ListOptionFunc {
	return func(l *types.ListOption) {
		(*l)["pageNo"] = fmt.Sprintf("%d", pageNo)
	}
}

func WithPageSize(pageSize int) ListOptionFunc {
	return func(l *types.ListOption) {
		(*l)["pageSize"] = fmt.Sprintf("%d", pageSize)
	}
}

func WithSearchVal(searchVal string) ListOptionFunc {
	return func(l *types.ListOption) {
		(*l)["searchVal"] = searchVal
	}
}
