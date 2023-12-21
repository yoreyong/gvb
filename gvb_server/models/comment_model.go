package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	SubComment         []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comment"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"`
	LikeCount          int             `gorm:"size:8;default:0" json:"like_count"`
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID          uint            `json:"article_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `gorm:"size:10" json:"user_id"`
}
