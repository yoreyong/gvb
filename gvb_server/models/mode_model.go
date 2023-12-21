package models

import (
	"gorm.io/gorm"
	"gvb_server/models/ctype"
)

type ModeModel struct {
	gorm.Model
	UserID    uint        `json:"user_id"`
	UserModel UserModel   `gorm:"foreignKey:UserID" json:"-"`
	Avatar    string      `json:"avatar"`
	IP        string      `gorm:"size:32" json:"ip"`
	Addr      string      `gorm:"size:64" json:"addr"`
	Content   string      `gorm:"size:256" json:"content"`
	Drawing   ctype.Array `gorm:"type:longtext" json:"drawing"`
}
