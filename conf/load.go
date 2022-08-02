package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//如果配置映射成config对象

// LoadConfigFromToml 从toml文件中加载配置
func LoadConfigFromToml(filePath string)  error {
	//读取toml格式的配置
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error ,path: %s, %s", filePath, err)
	}
	return nil
}

// LoadConfigFromTEnv 从环境变量中加载配置
func LoadConfigFromTEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)

}
