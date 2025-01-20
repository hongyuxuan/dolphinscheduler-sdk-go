package option

import (
	types "github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type ProcessDefinitionOptionFunc func(*types.ProcessDefinitionOption)

func WithLocations(locations *string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["locations"] = *locations
	}
}

func WithProcessName(name string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["name"] = name
	}
}

func WithTaskDefinitionJson(taskDefinitionJson string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["taskDefinitionJson"] = taskDefinitionJson
	}
}

func WithTaskRelationJson(taskRelationJson string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["taskRelationJson"] = taskRelationJson
	}
}

func WithTenantCode(tenantCode string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["tenantCode"] = tenantCode
	}
}

func WithGlobalParams(globalParams string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["globalParams"] = globalParams
	}
}

func WithReleaseState(releaseState string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["releaseState"] = releaseState
	}
}

func WithSchedule(scheduleJson string) ProcessDefinitionOptionFunc {
	return func(l *types.ProcessDefinitionOption) {
		(*l)["scheduleJson"] = scheduleJson
	}
}
