package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//如果配置映射成config对象

// LoadConfigFromToml 从toml文件中加载配置
func LoadConfigFromToml(filePath string) error {
	//读取toml格式的配置
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error ,path: %s, %s", filePath, err)
	}

	return nil
	return loadGlobal()
}

// LoadConfigFromTEnv 从环境变量中加载配置
func LoadConfigFromTEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}

	//return loadGlobal() //2种方式

	return nil

}

//加载全局实例
func loadGlobal() (err error) {
	//加载sql
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return err
	}
	return nil
}
