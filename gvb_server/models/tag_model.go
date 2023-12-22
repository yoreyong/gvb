package models

import "gorm.io/gorm"

const (
	TagArticle = 1
	TagSite    = 2
)

type TagModel struct {
	gorm.Model
	Title   string         `gorm:"size:16" json:"title"`
	Article []ArticleModel `gorm:"many2many:article_tag_models" json:"-"`
	TagType int            `gorm:"size:1; default:1" json:"tagType"`
}
