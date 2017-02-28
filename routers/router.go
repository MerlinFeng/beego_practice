package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "*:Register")
	beego.Router("/doreg", &controllers.UserController{}, "*:Save")
}
