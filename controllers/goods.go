package controllers

import (
	"LichFresh/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) ShowIndex() {
	username := this.GetSession("username")
	if username != nil {
		this.Data["username"] = username.(string)
	} else {
		this.Data["username"] = ""
	}

	//获取orm对象
	o := orm.NewOrm()

	//定义容器
	var goodsTypes []models.GoodsType

	//查询所有的商品类型
	o.QueryTable("GoodsType").All(&goodsTypes)
	this.Data["goodsTypes"] = goodsTypes

	//轮播图
	var goodsLunbo []models.IndexGoodsBanner
	o.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&goodsLunbo)
	this.Data["goodsLunbo"] = goodsLunbo

	//获取促销商品
	var goodsPro []models.IndexPromotionBanner
	o.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&goodsPro)
	this.Data["goodsPro"] = goodsPro




	//获取首页分类商品展示,定义了一个切片,注意这里是一个map切片,而不是map
	var goods []map[string]interface{}

	//把所有的商品类型插入到大容器中
	//前面已经把所有的商品类型goodsType查询出来了
	//那么我们把goodsTypes遍历一下,每个value就是一个商品类型GoodType
	for _,v := range goodsTypes{

		//先定义出一个每一行的容器,每一行都是一个map,key是string,value是interface{}
		//然后给每一行的map的goodsType赋值一个GoodsType
		//所以 temp["goodsType"]=一个GoodsType对象
		temp := make(map[string]interface{})
		temp["goodsType"] = v

		//把临时map放到大容器中,所以这里就是把map放到切片里面
		//goods[0]=temp01    temp01["goodsType"]=GoodsType对象01
		//goods[1]=temp02    temp02["goodsType"]=GoodsType对象01
		goods = append(goods,temp)
	}

	//把类型对应的首页展示商品插入到大容器中
	//现在我们遍历的是goods切片,也就是map切片
	for _,v := range goods{

		//获取到类型对应的所有商品
		//我们要找到的是每一行的Banner,所以指向IndexTypeGoodsBanner表,然后我们关联一下GoodsSKU和GoodsType表
		//然后我们要筛选一下,我们根据GoodsType来筛选
		qs :=o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsSKU","GoodsType").Filter("GoodsType",v["goodsType"])

		//需要把商品放到map
		//把文字对象全部查询出来
		var goodsText []models.IndexTypeGoodsBanner
		qs.Filter("DisplayType",0).OrderBy("Index").All(&goodsText)
		//获取图片商品
		//把图片对象全都查询出来
		var goodsImage []models.IndexTypeGoodsBanner
		qs.Filter("DisplayType",1).OrderBy("Index").All(&goodsImage)

		//插入到大容器
		//goods[0]= map
		//map["goodsType"] = GoodsType对象
		//map["goodsText"] = GoodsText切片
		//map["goodsImage"] = GoodsImage切片
		v["goodsText"] = goodsText
		v["goodsImage"] = goodsImage

	}
	//goods类型是map切片
	this.Data["goods"] = goods


	this.Layout = "layout.html" //指定布局页面
	this.TplName = "index.html" //显示主要页面
}

func (this *GoodsController) ShowUserCenterInfo() {
	username := this.GetSession("username")
	var user models.User
	user.Username = username.(string)

	o := orm.NewOrm()
	o.Read(&user, "Username")

	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Id", user.Id).Filter("IsDefault", true).One(&addr)

	this.Data["addr"] = addr
	this.Data["username"] = username

	this.Layout = "layout.html"
	this.TplName = "user_center_info.html"
}

func (this *GoodsController) ShowUserCenterOrder() {
	this.TplName = "user_center_order.html"
}

func (this *GoodsController) ShowUserCenterSite() {
	username := this.GetSession("username")
	var user models.User
	user.Username = username.(string)

	o := orm.NewOrm()
	o.Read(&user, "Username")

	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Id", user.Id).Filter("IsDefault", true).One(&addr)

	var addrs []models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Id", user.Id).All(&addrs)

	this.Data["addr"] = addr
	this.Data["addrs"] = addrs
	this.Data["username"] = username

	this.TplName = "user_center_site.html"
}

func (this *GoodsController) HandleSite() {
	username := this.GetSession("username")
	var user models.User
	user.Username = username.(string)

	receiver := this.GetString("receiver")
	addr := this.GetString("addr")
	zipcode := this.GetString("zipcode")
	phone := this.GetString("phone")

	var address models.Address
	address.Receiver = receiver
	address.Addr = addr
	address.Zipcode = zipcode
	address.Phone = phone

	o := orm.NewOrm()
	o.Read(&user, "Username")

	address.User = &user

	o.Insert(&address)

	this.TplName = "user_center_site.html"
}

