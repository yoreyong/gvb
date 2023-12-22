package models

import (
	"gorm.io/gorm"
	"gvb_server/models/ctype"
)

type UserModel struct {
	gorm.Model
	//MODEL
	NickName         string           `gorm:"size:36" json:"nick_name"`
	UserName         string           `gorm:"size:36" json:"user_name"`
	Password         string           `gorm:"size:128" json:"password"`
	Avatar           string           `gorm:"size:256" json:"avatar"`
	Email            string           `gorm:"size:128" json:"email"`
	Tel              string           `gorm:"size:18" json:"tel"`
	Addr             string           `gorm:"size:64" json:"addr"`
	Token            string           `gorm:"size:64" json:"token"` // 其他平台唯一ID
	IP               string           `gorm:"size:20" json:"ip"`
	Role             ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus       ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
	ArticleModels    []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`
	SubscribesModels []ArticleModel   `gorm:"many2many:user_subscribe_model;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"`
}
