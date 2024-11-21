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

	routerGroupApp.SettingRouter()
	return r
}
