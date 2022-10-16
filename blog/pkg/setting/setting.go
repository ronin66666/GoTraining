package setting

import "github.com/spf13/viper"

// 配置文件，使用三方viper来读取
//viper是允许设置多个配置路径的，这样可以尽可能地尝试解决路径查找问题，也就是说可以不断调用AddConfigPath方法

type Setting struct {
	vp *viper.Viper
}

//NewSetting 用于初始化本地配置的基础属性，设置配置文件的名称，路径、类型
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{
		vp: vp,
	}, nil
}
