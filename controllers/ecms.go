//接受ecms传过来的成果清单
package controllers

import (
	"encoding/json"
	"github.com/3xxx/meritms/models"
	"github.com/astaxie/beego"
	"time"
)

type EcmsController struct {
	beego.Controller
}

//Catalog添加附件链接和设计说明、校审意见
type CatalogLinkConts struct {
	Id            int64
	ProjectNumber string    //项目编号
	ProjectName   string    //项目名称
	DesignStage   string    //阶段
	Section       string    //专业
	Tnumber       string    //成果编号
	Name          string    //成果名称
	Category      string    //成果类型
	Page          string    //成果计量单位
	Count         float64   //成果数量
	Drawn         string    //编制、绘制
	Designd       string    //设计
	Checked       string    //校核
	Examined      string    //审查
	Verified      string    //核定
	Approved      string    //批准
	Complex       float64   //难度系数
	Drawnratio    float64   //编制、绘制占比系数
	Designdratio  float64   //设计系数
	Checkedratio  float64   //校核系数
	Examinedratio float64   //审查系数
	Datestring    string    //保存字符型日期
	Date          time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now_add;type(datetime)"`
	Author        string    //上传者
	State         int
	// Catalog models.Catalog
	Link    []models.CatalogLink
	Content []models.CatalogContent
}

func (c *EcmsController) GetEcmsPost() {
	// username := c.Input().Get("username")
	// beego.Info(username)
	// password := c.Input().Get("password")
	ecmsip := c.Input().Get("ecmsip")
	// beego.Info(ecmsip)
	ecmsport := c.Input().Get("ecmsport")
	var ob []CatalogLinkConts
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	var ob1 []models.Catalog
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob1)
	for i, v := range ob {
		ob1[i].Id = 0
		ob1[i].State = 1
		// json.Unmarshal(c.Ctx.Input.RequestBody, &ob1)
		// ob.State = 2 //说明单独改变其中一个字段是可以的。
		cid, err, news := models.SaveCatalog(ob1[i])
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range v.Link {
				_, err = models.AddCatalogLink(cid, "http://"+ecmsip+":"+ecmsport+v1.Url)
				if err != nil {
					beego.Error(err)
				}
			}
			data := news
			c.Ctx.WriteString(data)
		}
	}
	// c.Data["json"] = ""
	// c.ServeJSON()
}
