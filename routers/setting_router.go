package routers

import (
	"gin-vue-blog/controllers"
)

func (r RouterGroup) SettingRouter() {

	r.GET("/settings", controllers.SettingController{}.SettingsInfoView)
	//router.PUT("settings/:name", controllers.SettingController{}.SettingsInfoUpdateView)
}