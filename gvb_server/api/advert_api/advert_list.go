package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (AdvertApi) AdvertListWithPageView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	advertList, count, err := common.ComListPage(models.AdvertiseModel{IsShown: true}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(advertList, count, c)
}
