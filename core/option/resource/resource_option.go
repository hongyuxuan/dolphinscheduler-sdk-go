package resource

import (
	"fmt"

	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ResourceOptionFunc func(*types.ResourceOption)

func WithContent(content string) ResourceOptionFunc {
	return func(l *types.ResourceOption) {
		(*l)["content"] = content
	}
}

func WithFullName(fullName string) ResourceOptionFunc {
	return func(l *types.ResourceOption) {
		(*l)["fullName"] = fullName
	}
}

func WithTenantCode(tenantCode string) ResourceOptionFunc {
	return func(l *types.ResourceOption) {
		(*l)["tenantCode"] = tenantCode
	}
}

func WithLimit(limit int64) ResourceOptionFunc {
	return func(l *types.ResourceOption) {
		(*l)["limit"] = fmt.Sprintf("%d", limit)
	}
}
