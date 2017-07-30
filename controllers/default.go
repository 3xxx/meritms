package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	username, role := checkprodRole(c.Ctx)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "admin/index.html"
}

func (c *MainController) Slide() {
	c.TplName = "slide.tpl"
}

func (c *MainController) IsSubmitAgain(token string) bool {
	cotoken := c.Ctx.GetCookie("token")
	if token == "" || len(token) == 0 || token != cotoken || strings.Compare(cotoken, token) != 0 {
		return true
	}
	return false
}

func (c *MainController) Register() {
	// flash := beego.NewFlash()
	token := c.Input().Get("token")
	//是否重复提交
	if c.IsSubmitAgain(token) {
		c.Redirect("/registerpage", 302)
		return
	}

}
func Postdata() {
	b := httplib.Post("http://127.0.0.1:80/postdata")
	b.Param("username", "astaxie")
	b.Param("password", "123456")
	b.PostFile("uploadfile", ".\\database\\meritms.db") //./static/
	// b.PostFile("uploadfile2", "httplib.txt")PostFile 第一个参数是 form 表单的字段名,第二个是需要发送的文件名或者文件路径
	str, err := b.String()
	if err != nil {
		beego.Error(str)
	}
}
