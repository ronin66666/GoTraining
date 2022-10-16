package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

// 全局变量配置
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
