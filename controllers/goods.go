package controllers

import (
	"github.com/astaxie/beego"
	"LichFresh/models"
	"github.com/astaxie/beego/orm"
)

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) ShowIndex() {
	username := this.GetSession("username").(string)
	this.Data["username"] = username

	this.TplName = "index.html"
}

func (this *GoodsController) ShowUserCenterInfo() {
	username := this.GetSession("username")
	var user models.User
	user.Username = username.(string)

	o := orm.NewOrm()
	o.Read(&user, "username")

	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Id", user.Id).Filter("IsDefault", true).One(&addr)

	//向视图中传递数据
	this.Data["addr"] = addr
	this.Data["username"] = username

	this.TplName = "user_center_info.html"
}

func (this *GoodsController) ShowUserCenterOrder() {
	this.TplName = "user_center_order.html"
}

func (this *GoodsController) ShowUserCenterSite() {
	this.TplName = "user_center_site.html"
}
