package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertiseModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		fmt.Println(count)
		res.FailWithMessage("Files does not exist!", c)
		return
	}

	// Delete records from database
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("Totally delete %d advertise files", count), c)
}
