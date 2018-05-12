//在线成果登记
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/tealeg/xlsx"
	"github.com/3xxx/meritms/models"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	// "github.com/astaxie/beego/logs"
	// "sort"
	// "strconv"
	// "strings"
	// "time"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Test() {
	// users, err := models.GetUser1(1)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["json"] = users

	// profile, err := models.Getprofile(1)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["json"] = profile

	posts, err := models.GetPost(1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = posts

	// user, err := models.GetUser("The Title")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["json"] = user

	c.ServeJSON()
}
