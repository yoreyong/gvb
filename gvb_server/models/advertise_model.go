package models

import "gorm.io/gorm"

// AdvertiseModel 广告表
type AdvertiseModel struct {
	gorm.Model
	Title   string `gorm:"size:32" json:"title"`
	Href    string `json:"href"`
	Images  string `json:"images"`
	IsShown string `json:"is_shown"`
}
