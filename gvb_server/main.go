package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// Read config file
	core.InitConf()
	fmt.Println("global config structure:", global.Config)

	// Init logger
	global.Log = core.InitLogger()
	global.Log.Warnln("xixixi")
	global.Log.Error("xixixi")
	global.Log.Infof("xixixi")

	// Connect Mysql
	global.DB = core.InitGrom()
	fmt.Println(global.DB)
}
