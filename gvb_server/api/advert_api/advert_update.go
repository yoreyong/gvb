package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println(err)
		res.FailWithError(err, &cr, c)
		return
	}

	// check for duplicate
	var advert models.AdvertiseModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("Advertise does not exist", c)
		return
	}

	// Write DB
	// 用Updates时, 只能传入 struct 或 map, 详情请看官方文档
	err = global.DB.Model(&advert).Updates(map[string]interface{}{
		"title":    cr.Title,
		"href":     cr.Href,
		"images":   cr.Images,
		"is_shown": cr.IsShown,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Failed to update advertise record", c)
		return
	}
	res.OkWithMessage("Success update advertise record", c)
}
