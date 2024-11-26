package models

import "gorm.io/gorm"

// 广告表
type Advertisement struct {
	gorm.Model
	Title  string `gorm:"size:32" json:"title"`         // 广告标题 唯一
	Href   string `gorm:"size:64" json:"href"`          // 广告链接
	Images string `son:"images"`                        // 图片
	IsShow *bool  `gorm:"default:false" json:"is_show"` // 是否显示
}
