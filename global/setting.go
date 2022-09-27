package global

import "goWeb/pkg/setting"

// 对最初预估的三个区段进行配置 并 声明全局变量，
// 便于在后面步骤中相关联，让应用程序内部使用

// 注意：全局变量的初始化会随着应用程序的不断演进而不断改变，这里展示的也不一定是最终结果
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatebaseSettingS
)