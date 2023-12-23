package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsEmailView GET
func (SettingsApi) SettingsEmailView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}

// SettingsEmailUpdateView UPDATE
func (SettingsApi) SettingsEmailUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.Ok(c)
}
