package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// 编写NewSetting方法，用于初始化本项目配置的基础属性 =》
//
//	设定配置文件名称为config
//	配置类型为yaml
//	设置配置路径为行对路径 configs/
//
// =》 以上措施都是为了能够成功启动编写组件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
