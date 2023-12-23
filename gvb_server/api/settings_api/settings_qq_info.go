package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsQQView GET
func (SettingsApi) SettingsQQView(c *gin.Context) {
	res.OkWithData(global.Config.QQ, c)
}

// SettingsQQUpdateView UPDATE
func (SettingsApi) SettingsQQUpdateView(c *gin.Context) {
	var cr config.QQ
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.QQ = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.Ok(c)
}
