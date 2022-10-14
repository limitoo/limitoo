package initcfg

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func ConfigInit(cfg string) error {

	c := Config{Name: cfg}

	err := c.initconfig()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) initconfig() error {

	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("configs") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("env")
	}
	viper.SetConfigType("yaml")   // 设置配置文件格式为YAML
	viper.AutomaticEnv()          // 读取匹配的环境变量
	viper.SetEnvPrefix("LIMITOO") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}
