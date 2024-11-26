package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Title    string    `gorm:"size:32" json:"title"`           // 标签名
	Articles []Article `gorm:"many2many:article_tag" json:"_"` // 标签对应的文章
}
