package option

import (
	"fmt"

	"github.com/hongyuxuan/dolphinscheduler-sdk-go/types"
)

type StartInstanceOptionFunc func(*types.StartInstanceOption)

func WithProcessName(processName string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["processName"] = processName
	}
}

func WithProcessDefinitionCode(processCode int64) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["processDefinitionCode"] = fmt.Sprintf("%d", processCode)
	}
}

func WithFailureStrategy(failureStrategy string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["failureStrategy"] = failureStrategy
	}
}

func WithWarningType(warningType string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["warningType"] = warningType
	}
}

func WithWarningGroupId(warningGroupId int64) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["warningGroupId"] = fmt.Sprintf("%d", warningGroupId)
	}
}

func WithExecType(execType string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["execType"] = execType
	}
}

func WithStartNodeList(startNodeList string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["startNodeList"] = startNodeList
	}
}

func WithTaskDependType(taskDependType string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["taskDependType"] = taskDependType
	}
}

func WithDependentMode(dependentMode string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["complementDependentMode"] = dependentMode
	}
}

func WithRunMode(runMode string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["runMode"] = runMode
	}
}

func WithInstancePriority(instancePriority string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["processInstancePriority"] = instancePriority
	}
}

func WithWorkerGroup(workerGroup string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["workerGroup"] = workerGroup
	}
}

func WithTenantCode(tenantCode string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["tenantCode"] = tenantCode
	}
}

func WithEnvironmentCode(environmentCode string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["environmentCode"] = environmentCode
	}
}

func WithStartParams(startParams string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["startParams"] = startParams
	}
}

func WithParallelism(expectedParallelismNumber string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["expectedParallelismNumber"] = expectedParallelismNumber
	}
}

func WithDryRun(dryRun string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["dryRun"] = dryRun
	}
}

func WithTestFlag(testFlag string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["testFlag"] = testFlag
	}
}

func WithVersion(version string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["version"] = version
	}
}

func WithAllLevelDependent(allLevelDependent *string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["allLevelDependent"] = *allLevelDependent
	}
}

func WithExecutionOrder(executionOrder string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["executionOrder"] = executionOrder
	}
}

func WithScheduleTime(scheduleTime string) StartInstanceOptionFunc {
	return func(l *types.StartInstanceOption) {
		(*l)["scheduleTime"] = scheduleTime
	}
}
