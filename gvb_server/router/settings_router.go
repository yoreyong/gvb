package router

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// info APIs
	router.GET("settings/info", settingsApi.SettingsInfoView)
	router.PUT("settings/info", settingsApi.SettingsInfoUpdateView)

	// email APIs
	router.GET("settings/email", settingsApi.SettingsEmailView)
	router.PUT("settings/email", settingsApi.SettingsEmailUpdateView)

	// jwt APIs
	router.GET("settings/jwt", settingsApi.SettingsJwtView)
	router.PUT("settings/jwt", settingsApi.SettingsJwtUpdateView)

	// qiniu APIs
	router.GET("settings/qiniu", settingsApi.SettingsQiNiuView)
	router.PUT("settings/qiniu", settingsApi.SettingsQiNiuUpdateView)

	// qq APIs
	router.GET("settings/qq", settingsApi.SettingsQQView)
	router.PUT("settings/qq", settingsApi.SettingsQQUpdateView)
}
