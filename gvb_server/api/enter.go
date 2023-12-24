package api

import (
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

// ApiGroupApp 实例化结构体 ApiGroup
var ApiGroupApp = new(ApiGroup)
