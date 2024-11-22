package controllers

import (
	"gin-vue-blog/utils/res"
	"github.com/gin-gonic/gin"
)

type SettingController struct{}

// 查询页面
func (SettingController) SettingsInfoView(c *gin.Context) {
	res.OkWith(c)
}
