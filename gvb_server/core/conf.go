package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"
)

// InitConf Read yaml config files
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config yamFile Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	//fmt.Println(c)
	global.Config = c // 将读取的数据指向全局变量
}
