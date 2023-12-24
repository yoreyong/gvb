package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

//func (ImagesApi) ImageListView(c *gin.Context) {
//	var imageList []models.BannerModel
//
//	global.DB.Find(&imageList)
//
//	res.OkWithData(imageList, c)
//
//	return
//}

// ImageListWithPageView 图片列表查询页
func (ImagesApi) ImageListWithPageView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//var imageList []models.BannerModel
	//count := global.DB.Select("id").Find(&imageList).RowsAffected
	//offset := (cr.Page - 1) * cr.Limit
	//if offset < 0 {
	//	offset = 0
	//}
	//global.DB.Limit(cr.Limit).Offset(offset).Find(&imageList)
	imageList, count, err := common.ComListPage(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	//res.OkWithData(gin.H{
	//	"count": count,
	//	"list":  imageList,
	//}, c)
	res.OkWithList(imageList, count, c)

	return
}
