package models

import "gorm.io/gorm"

// MessageModel 记录消息
type Message struct {
	gorm.Model
	SendUserID       uint   `gorm:"primaryKey" json:"send_user_id"` // 发送者id
	SendUserModel    User   `gorm:"foreignKey:SendUserID" json:"_"`
	SendUserNickName string `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string `json:"send_user_avatar"`

	RevUserID       uint   `gorm:"primaryKey" json:"rev_user_id"` // 接收者id
	RevUserModel    User   `gorm:"foreignKey:RevUserID" json:"_"`
	RevUserNickName string `gorm:"size:42" json:"rev_user_nick_name"`
	RevUserAvatar   string `json:"rev_user_avatar"`

	IsRead  bool   `gorm:"default:false" json:"is_read"` // 接收方是否已读
	Content string `gorm:"type:longtext" json:"content"` // 消息内容
}
