# Dolphinscheduler Golang SDK
注意，本 SDK 分 v2 和 v3 两个版本，其中 v2 适配 dolphinscheduler 2.* 版本，v3 适配 dolphinscheduler 3.* 版本。

# 接口列表
- 获取项目列表
- 资源中心获取文件列表
- 安全中心获取环境列表
- 安全中心获取告警组列表
- 安全中心获取租户列表
- 获取工作流定义列表
- 根据Code获取工作流定义
- 根据Name获取工作流定义
- 解析Json文件到工作流定义
- 从文件导入工作流定义
- 从Bytes导入工作流定义
- 修改工作流定义
- 修改工作流定义基本信息（定时信息）
- 工作流定义上线/下线
- 删除工作流定义
- 获取工作流定时列表
- 工作流定定时上线/下线
- 删除工作流定时

# v2用法
1. 引入包
```go
import (
	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
)
```

2. 新建到 dolphinscheduler 连接
```go
client := ds.NewClientV2(
  option.WithDebug(true),
  option.WithBaseUrl("http://<your_dolphinscheduler_host>/dolphinscheduler"),
  option.WithToken("<your_token>"))
```

3. 实例1：获取项目列表
```go
res, err := client.Project(nil).List(
  context.Background(), 
  option.WithPageNo(1), 
  option.WithPageSize(10), 
  option.WithSearchVal("<your_search_key>"))
if err != nil {
  log.Fatal(err)
}
fmt.Println(res)
```

4. 实例2：从文件导入工作流定义
```go
if err := client.Project(<projectCode>).
  ProcessDefinition(nil).
  ImportFile(context.Background(), "<path_to_your_file>"); err != nil {
  log.Fatal(err)
}
```

更多示例详见test。

# v3用法
1. 引入包
```go
import (
	ds "github.com/hongyuxuan/dolphinscheduler-sdk-go"
)
```

2. 新建到 dolphinscheduler 连接
```go
client := ds.NewClientV3(
  option.WithDebug(true),
  option.WithBaseUrl("http://<your_dolphinscheduler_host>/dolphinscheduler"),
  option.WithToken("<your_token>"))
```

3. 实例1：获取项目列表
```go
res, err := client.Project(nil).List(
  context.Background(), 
  option.WithPageNo(1), 
  option.WithPageSize(10), 
  option.WithSearchVal("<your_search_key>"))
if err != nil {
  log.Fatal(err)
}
fmt.Println(res)
```

4. 实例2：从文件导入工作流定义
```go
if err := client.Project(<projectCode>).
  ProcessDefinition(nil).
  ImportFile(context.Background(), "<path_to_your_file>"); err != nil {
  log.Fatal(err)
}
```