package models

type MenuBannerModel struct {
	MenuID      uint        `gorm:"primaryKey" json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `gorm:"primaryKey" json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
