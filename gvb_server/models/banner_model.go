package models

import "gorm.io/gorm"

type BannerModel struct {
	gorm.Model
	Path string `json:"path"`                // path of image
	Hash string `json:"hash"`                // hashcode of image, used to find duplicates
	Name string `gorm:"size:38" json:"name"` // name of image
}
