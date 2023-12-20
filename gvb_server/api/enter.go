package api

import "gvb_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

// ApiGroupApp 实例化结构体 ApiGroup
var ApiGroupApp = new(ApiGroup)
