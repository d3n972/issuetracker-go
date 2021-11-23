package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//c.Layout = "layouts/app.tpl"
	c.TplName = "pages/index.tpl"
	c.Data["title"] = "asd"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

}
