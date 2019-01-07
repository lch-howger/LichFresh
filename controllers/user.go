package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"github.com/astaxie/beego/orm"
	"LichFresh/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowLogin() {
	this.TplName = "login.html"
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

func (this *UserController) HandleLogin() {
	username := this.GetString("username")
	password := this.GetString("pwd")

	if username == "" || password == "" {
		this.Data["err"] = "用户名或密码不能为空"
		this.TplName = "login.html"
		return
	}

	o := orm.NewOrm()
	var user models.User
	user.Username = username

	err := o.Read(&user, "Username")

	if err != nil || user.Password != password {
		this.Data["err"] = "用户名或密码错误"
		this.TplName = "login.html"
		return
	}

	beego.Info("username是",username)

	this.SetSession("username", username)
	this.Redirect("/index", 302)

}

func (this *UserController) HandleRegister() {
	username := this.GetString("user_name")
	pwd := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")

	if username == "" || pwd == "" || cpwd == "" || email == "" {
		this.Data["err"] = "输入数据不完整"
		this.TplName = "register.html"
		return
	}

	reg, err := regexp.Compile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if err != nil {
		this.Data["err"] = "正则创建失败"
		this.TplName = "register.html"
		return
	}

	res := reg.MatchString(email)
	if res == false {
		this.Data["err"] = "邮箱格式不正确"
		this.TplName = "register.html"
		return
	}

	if pwd != cpwd {
		this.Data["err"] = "两次密码不一致"
		this.TplName = "register.html"
		return
	}

	//开始处理数据
	o := orm.NewOrm()
	var user models.User
	user.Username = username
	user.Password = pwd
	user.Email = email

	//插入
	_, err = o.Insert(&user)
	if err != nil {
		this.Data["err"] = "用户名重复"
		this.TplName = "register.html"
		return
	}

	this.Redirect("/login", 302)
}
