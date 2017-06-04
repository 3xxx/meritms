//接受ecms传过来的成果清单
package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"meritms/models"
)

type EcmsController struct {
	beego.Controller
}

func (c *EcmsController) GetEcmsPost() {
	var ob models.Catalog
	// beego.Info(c.Ctx.Input.RequestBody)
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// objectid := models.SaveCatalog(ob)
	// beego.Info(ob)
	ob.Id = 0
	// ob.State = 2 //说明单独改变其中一个字段是可以的。
	_, err, news := models.SaveCatalog(ob)
	if err != nil {
		beego.Error(err)
	} else {
		data := news
		c.Ctx.WriteString(data)
	}
	// c.Data["json"] = ""
	// c.ServeJSON()
}
