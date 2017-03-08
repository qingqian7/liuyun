package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["true"] = true
	c.Data["false"] = false
	c.TplName = "index.tpl"
	type user struct {
		Name   string
		Age    int
		Gender string
	}
	User := &user{
		Name:   "lijun",
		Age:    21,
		Gender: "Male",
	}
	c.Data["user"] = User
	Num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c.Data["num"] = Num
	c.Data["html"] = "<div>hello beego</div>"
}
