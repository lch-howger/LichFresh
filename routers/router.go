package routers

import (
	"LichFresh/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
}
