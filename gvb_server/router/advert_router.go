package router

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi

	router.GET("adverts", advertApi.AdvertListWithPageView)
	router.POST("adverts", advertApi.AdvertCreateView)
	router.PUT("adverts/:id", advertApi.AdvertUpdateView)
	router.DELETE("adverts", advertApi.AdvertRemoveView)
}
