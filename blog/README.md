# blog-service

## 流程

### 项目目录
- configs: 配置文件
- docs: 文档集合
- global：全局变量
- internal：内部模块
  - dao：数据访层，所有数据相关的操作都会在dao层进行，例如：MySQL、Elasticsearch等
  - middleware：HTTP中间件
  - model：模型层，用于存放model
  - routers：路由相关逻辑
  - services：项目核心业务逻辑
- pkg：项目相关的模块包
- storage：项目生成的临时文件
- scripts：各类构建、安装、分析等操作的脚本
- third_party：第三方资源工具，如Swagger UI

### 创建数据库、表结构
创建语句： `blog/scripts/create.sql`
### 创建模型
在`internal/model`目录中创建相应的模型
### 创建路由
在 `internal/routers`目录中创建router
### 错误标准化
`pkg/errcode`
与客户端保持统一的错误码
### 配置管理：

在应用程序的运行生命周期中，应用的配置读取和更新可以直接改变应用程序。
- 在启动时：可以进行一些初始化，如配置基础一应用属性，连接第三方实例（MySql）等
- 在运行中：可以通过监听文件或变更其他储存载体来实现热更新配置效果

使用Viper开源库来完成配置的读取
`go get -u github.com/spf13/viper`
配置文件: `configs/config.yaml`
### 链接数据库
使用三方`gorm`
`go get -u github.com/jinzhu/gorm`
### 日志组件
在一个项目中，日志需要标准化地记录一些公共信息，如代码调用堆栈，请求链路ID、公共业务属性字段等
使用lumberjack三方日志库
`go get -u gopkg.in/natefinch/lumberjack.v2`
核心功能：把日志写入到滚动文件中，该库允许我们设置单日志文件的最大占用空间、最大生存周期、可保留的最多就文件数等，如果有出现超出设置项的情况，则对日志文件进行滚动处理
这个库可以减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上