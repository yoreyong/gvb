package router

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi

	router.POST("images", imagesApi.ImagesUploadView)
	router.GET("images", imagesApi.ImageListWithPageView)
	router.DELETE("images", imagesApi.ImageRemoveView)
}
