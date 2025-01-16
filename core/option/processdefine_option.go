package option

import (
	typesv2 "github.com/hongyuxuan/dolphinscheduler-sdk-go/types/v2"
)

type ProcessDefinitionOptionFunc func(*typesv2.ProcessDefinitionOption)

func WithLocations(locations *string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["locations"] = *locations
	}
}

func WithProcessName(name string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["name"] = name
	}
}

func WithTaskDefinitionJson(taskDefinitionJson string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["taskDefinitionJson"] = taskDefinitionJson
	}
}

func WithTaskRelationJson(taskRelationJson string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["taskRelationJson"] = taskRelationJson
	}
}

func WithTenantCode(tenantCode string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["tenantCode"] = tenantCode
	}
}

func WithGlobalParams(globalParams string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["globalParams"] = globalParams
	}
}

func WithReleaseState(releaseState string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["releaseState"] = releaseState
	}
}

func WithSchedule(scheduleJson string) ProcessDefinitionOptionFunc {
	return func(l *typesv2.ProcessDefinitionOption) {
		(*l)["scheduleJson"] = scheduleJson
	}
}
