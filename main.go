package main

import (
	_ "LichFresh/models"
	_ "LichFresh/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

