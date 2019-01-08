package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"github.com/astaxie/beego/orm"
	"LichFresh/models"
	"github.com/astaxie/beego/utils"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowLogin() {
	username := this.Ctx.GetCookie("username")

	if username != "" {
		this.Data["username"] = username
		this.Data["checked"] = "checked"
	} else {
		this.Data["username"] = ""
		this.Data["checked"] = ""
	}

	this.TplName = "login.html"
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

func (this *UserController) HandleLogin() {
	username := this.GetString("username")
	password := this.GetString("pwd")
	remember := this.GetString("remember")

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

	if user.Active != 1 {
		this.Data["err"] = "用户还未激活"
		this.TplName = "login.html"
		return
	}

	if remember == "on" {
		this.Ctx.SetCookie("username", username, time.Second*60*60*24)
	} else {
		this.Ctx.SetCookie("username", username, -1)
	}

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

	//发送邮件
	config := `{"username":"360000521@qq.com","password":"lnvobxmwmfpibghj","host":"smtp.qq.com","port":587}`
	temail := utils.NewEMail(config)
	temail.To = []string{user.Email}
	temail.From = "360000521@qq.com"
	temail.Subject = "用户激活"

	temail.HTML = "复制该连接到浏览器中激活：127.0.0.1:8080/active?id=" + strconv.Itoa(user.Id)

	err = temail.Send()
	if err != nil {
		this.Data["errmsg"] = "发送激活邮件失败，请重新注册！"
		this.TplName = "register.html"
		return
	}

	this.Ctx.WriteString("注册成功，请前往邮箱激活!")
}

func (this *UserController) HandleActive() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["err"] = "激活失败了,链接错误"
		this.TplName = "login.html"
		return
	}

	o := orm.NewOrm()
	var user models.User
	user.Id = id
	err = o.Read(&user, "Id")
	if err != nil {
		this.Data["err"] = "激活失败了,没有这个用户"
		this.TplName = "login.html"
		return
	}

	user.Active = 1
	_, err = o.Update(&user)
	if err != nil {
		this.Data["err"] = "激活失败了"
		this.TplName = "login.html"
		return
	}

	this.TplName = "login.html"
}
