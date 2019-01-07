package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string `orm:"unique;size(100)"`
	Password string `orm:"size(100)"`
	Email    string
	Power    int
	Active   int

	Receivers []*Receiver `orm:"reverse(many)"`
}

type Receiver struct {
	Id        int
	Name      string
	Postcode  string
	Address   string
	Phone     string
	IsDefault bool `orm:"default(false)"`

	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/fresh?charset=utf8")
	orm.RegisterModel(new(User), new(Receiver))
	orm.RunSyncdb("default", false, true)
}
