package setting

import "time"

//section.go 用于声明配置属性的结构体，并编写读取区段配置的配置方法

// ServerSettingS Server配置
type ServerSettingS struct {
	RunModel     string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS 应用配置
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}
type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

//ReadSection 读取配置文件
//k 配置文件中相应的名字，如configs/config.yaml 中的服务配置 k: Server
//v 模型如：&ServerSettingS
func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
