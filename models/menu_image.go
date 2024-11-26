package models

// 菜单图片关联表
type MenuImage struct {
	MenuID     uint  `json:"menu_id"`
	ImageID    uint  `json:"image_id"`
	MenuModel  Menu  `gorm:"foreignKey:MenuID" json:"menu_model"`
	ImageModel Image `gorm:"foreignKey:ImageID" json:"image_model"`
	Sort       int   `gorm:"size:10" json:"sort"`
}
