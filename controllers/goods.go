package controllers

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) ShowIndex() {
	username := this.GetSession("username").(string)
	this.Data["username"] = username
	this.TplName = "index.html"
}
