package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "admin/index.html"
}

func (c *MainController) Slide() {
	c.TplName = "slide.tpl"
}
