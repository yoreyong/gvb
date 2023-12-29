package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type AdvertRequest struct {
	Title   string `json:"title" binding:"required" msg:"Please input title"`
	Href    string `json:"href" binding:"required,url" msg:"Please input advertise href"` // binding中的url提供一种校验功能, 用于校验参数的数据内容是否是url
	Images  string `json:"images" binding:"required,url" msg:"Please input image href"`
	IsShown bool   `json:"is_shown" msg:"Please choose Shown Status"` // bool值如果使用binding:"required", 传入的值必须为True, 否则被判为不满足条件
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// check for duplicate
	var advert models.AdvertiseModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("Advertise already exists!", c)
		return
	}

	// Write DB
	err = global.DB.Create(&models.AdvertiseModel{
		Title:   cr.Title,
		Href:    cr.Href,
		Images:  cr.Images,
		IsShown: cr.IsShown,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Failed to create advertise record", c)
		return
	}
	res.OkWithMessage("Success create advertise record", c)
}
