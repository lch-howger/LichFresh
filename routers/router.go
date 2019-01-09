package routers

import (
	"LichFresh/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "get:ShowLogin")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/active", &controllers.UserController{}, "get:HandleActive")
	beego.Router("/logout",&controllers.UserController{},"*:HandleLogout")

	beego.Router("/index", &controllers.GoodsController{}, "get:ShowIndex")
	beego.Router("/userCenterInfo", &controllers.GoodsController{}, "get:ShowUserCenterInfo")
	beego.Router("/userCenterOrder", &controllers.GoodsController{}, "get:ShowUserCenterOrder")
	beego.Router("/userCenterSite", &controllers.GoodsController{}, "get:ShowUserCenterSite;post:HandleSite")

}
