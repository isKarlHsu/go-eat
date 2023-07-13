package core

import (
	"eat/config"
	"eat/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func InitConfig() {
	// todo 解决测试用例下目录结构问题
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(any("get yamlConfig error:" + err.Error()))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		panic(any("yamlConfig Unmarshal error:" + err.Error()))
	}
	log.Printf("yamlConfig load success")
	global.Config = c
}
