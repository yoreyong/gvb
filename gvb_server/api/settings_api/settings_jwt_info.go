package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsJwtView GET
func (SettingsApi) SettingsJwtView(c *gin.Context) {
	res.OkWithData(global.Config.Jwt, c)
}

// SettingsJwtUpdateView UPDATE
func (SettingsApi) SettingsJwtUpdateView(c *gin.Context) {
	var cr config.Jwt
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Jwt = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.Ok(c)
}
