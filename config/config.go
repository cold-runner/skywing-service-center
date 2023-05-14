package config

import (
	"github.com/spf13/viper"
)

type ServerConf struct {
	Mysql  `mapstructure:"mysql"`
	Redis  `mapstructure:"redis"`
	Jwt    `mapstructure:"jwt"`
	Log    `mapstructure:"log"`
	Server `mapstructure:"server"`
}

func (sc *ServerConf) init() error {
	viper.SetConfigName("config") // 配置文件名为 config.yaml
	viper.AddConfigPath(".")      // 配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err = viper.Unmarshal(sc); err != nil {
		return err
	}
	return nil
}

func NewConfig() *ServerConf {
	conf := &ServerConf{}
	err := conf.init()
	if err != nil {
		panic(err)
	}
	return conf
}
