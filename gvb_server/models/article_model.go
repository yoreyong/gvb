package models

import (
	"gorm.io/gorm"
	"gvb_server/models/ctype"
)

type ArticleModel struct {
	gorm.Model
	Title          string         `gorm:"size:32" json:"title"`
	Abstract       string         `json:"abstract"`
	Content        string         `json:"content"`
	ViewCount      int            `json:"view_count"`
	CommentCount   int            `json:"comment_count"`
	LikeCount      int            `json:"like_count"`
	SubscribeCount int            `json:"subscribe_count"`
	TagModels      []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"`
	CommentModel   []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`
	UserModel      UserModel      `gorm:"foreignKey:UserID" json:"-"`
	UserID         uint           `json:"user_id"`
	Category       string         `gorm:"size:20" json:"category"`
	Source         string         `json:"source"`
	Link           string         `json:"link"`
	Banner         BannerModel    `gorm:"foreignKey:BannerID" json:"-"`
	BannerID       uint           `json:"banner_id"`
	NickName       string         `gorm:"size:42" json:"nick_name"`
	BannerPath     string         `json:"banner_path"`
	Tags           ctype.Array    `gorm:"type:string;size:64" json:"tags"`
}
