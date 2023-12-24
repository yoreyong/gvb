package router

import "gvb_server/api"

func (router RouterGroup) ImagesUploadRouter() {
	imagesUploadApi := api.ApiGroupApp.ImagesApi

	router.POST("images", imagesUploadApi.ImagesUploadView)
	router.GET("images", imagesUploadApi.ImageListWithPageView)
}
