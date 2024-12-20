package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Email        string `gorm:"size:128" json:"email"`
	Content      string `gorm:"type:longtext" json:"content"`
	ApplyContent string `gorm:"type:longtext" json:"apply_content"` // 回复的内容
	IsApply      bool   `gorm:"default:false" json:"is_apply"`      // 是否回复
}
