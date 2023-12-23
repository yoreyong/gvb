package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsQiNiuView GET
func (SettingsApi) SettingsQiNiuView(c *gin.Context) {
	res.OkWithData(global.Config.QiNiu, c)
}

// SettingsQiNiuUpdateView UPDATE
func (SettingsApi) SettingsQiNiuUpdateView(c *gin.Context) {
	var cr config.QiNiu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.QiNiu = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.Ok(c)
}
