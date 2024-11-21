package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SettingController struct{}

// 查询页面
func (SettingController) SettingsInfoView(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
