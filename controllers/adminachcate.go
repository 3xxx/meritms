//成果类型和折标系数表
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/tealeg/xlsx"
	m "github.com/3xxx/meritms/models"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"github.com/astaxie/beego/logs"
	// "merit/models"
	// "sort"
	"strconv"
	// "strings"
	"time"
)

//进入编辑ratio页面
func (c *Achievement) Achievcategory() {
	//1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	return
	// }
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// if v != nil {
	// 	c.Data["Uname"] = v.(string)
	// }
	// //4.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	// if role > 2 {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	return
	// }
	ratio, _ := m.GetAchievcategories()
	c.Data["json"] = ratio
	c.ServeJSON()
	// c.TplName = "ratio.tpl"
}

func (c *Achievement) AddAchievcategory() {
	var ratio m.AdminAchievcategory
	ratio.Category = c.Input().Get("category")
	// ratio.Unit = c.Input().Get("Unit")
	Ismaterial := c.Input().Get("ismaterial")
	if Ismaterial == "true" {
		ratio.Ismaterial = true
	} else {
		ratio.Ismaterial = false
	}

	ratio1, err := strconv.ParseFloat(c.Input().Get("rationum"), 64)
	if err != nil {
		beego.Error(err)
	}
	ratio.Rationum = ratio1
	ratio.Created = time.Now()
	ratio.Updated = time.Now()
	cid, err := m.SaveAchievcategory(ratio)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok"
		c.Ctx.WriteString(data)
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "添加ratio" + strconv.FormatInt(cid, 10))
		logs.Close()
	}
}

func (c *Achievement) UpdateAchievcategory() {
	var ratio m.AdminAchievcategory
	ratio.Category = c.Input().Get("category")
	// ratio.Unit = c.Input().Get("Unit")
	Ismaterial := c.Input().Get("ismaterial")
	if Ismaterial == "true" {
		ratio.Ismaterial = true
	} else {
		ratio.Ismaterial = false
	}
	ratio1, err := strconv.ParseFloat(c.Input().Get("rationum"), 64)
	if err != nil {
		beego.Error(err)
	}
	ratio.Rationum = ratio1
	// ratio.Created = time.Now()
	ratio.Updated = time.Now()

	cid := c.Input().Get("cid")
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// var id string
	// if cid != "" {
	// 	id = string(cid[3:len(cid)])
	// 	// beego.Info(id)
	// }
	err = m.UpdateAchievcategory(cidNum, ratio)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改ratio" + cid)
		logs.Close()
	}
}

//删除一条类型记录
func (c *Achievement) DeleteAchievcategory() {
	cid := c.Input().Get("cid")
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = m.DeleteAchievcategory(cidNum)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "删除记录" + cid)
		logs.Close()
	}
}
