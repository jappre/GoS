package routers

import (
	"github.com/astaxie/beego"
	"GoS/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Get")
	beego.Router("/modify_email", &controllers.MainController{}, "post:ModifyEmail")
}
