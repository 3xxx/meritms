package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
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

func Postdata() {
	b := httplib.Post("http://127.0.0.1:80/postdata")
	b.Param("username", "astaxie")
	b.Param("password", "123456")
	b.PostFile("uploadfile1", "d:\\1.txt")
	// b.PostFile("uploadfile2", "httplib.txt")PostFile 第一个参数是 form 表单的字段名,第二个是需要发送的文件名或者文件路径
	str, err := b.String()
	if err != nil {
		beego.Error(str)
	}
}
