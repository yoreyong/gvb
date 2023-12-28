package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	gorm.Model
	Path      string          `json:"path"`                        // path of image
	Hash      string          `json:"hash"`                        // hashcode of image, used to find duplicates
	Name      string          `gorm:"size:38" json:"name"`         // name of image
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // Path of stored images: Local or QiNiu
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// Image is stored in local, remove local image before remove DB image path
		// 图片存储在本地,删除数据库图片路径的同时,还要删除本地的图片
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
