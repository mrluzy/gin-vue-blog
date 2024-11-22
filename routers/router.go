package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	routerGroupApp := RouterGroup{api}

	//配置路由
	routerGroupApp.SettingRouter()
	return r
}
