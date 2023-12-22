package models

import (
	"gorm.io/gorm"
	"gvb_server/models/ctype"
)

type MenuModel struct {
	gorm.Model
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`
	AbstractTime *int          `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_model;joinForeignKey:MenuID;JoinReferences:BannerID" json:"-"`
	BannerTime   *int          `json:"banner_time"`
	Sort         int           `gorm:"size:10" json:"sort"`
}
