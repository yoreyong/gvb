package models

import "gorm.io/gorm"

type FeedbackModel struct {
	gorm.Model
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ReplyContent string `gorm:"size:128" json:"reply_content"`
	IsReply      bool   `json:"is_reply"`
}
