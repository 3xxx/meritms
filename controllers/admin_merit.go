package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/meritms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "net"
	// "net/http"
	// "net/url"
	// "path"
	"strconv"
	"strings"
	// "time"
	"github.com/tealeg/xlsx"
	"os"
)

// CMSADMIN API
type AdminMeritController struct {
	beego.Controller
}

//**********价值***********
//取得所有价值分类，或没有下级的价值
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminMeritController) Merit() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id)
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	beego.Info(id)
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	merits, err := models.GetAdminMerit(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = merits
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据科室id得到价值分类，填充table
func (c *AdminMeritController) SecofficeMerit() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	merits, err := models.GetSecofficeMerit(idNum)
	if err != nil {
		beego.Error(err)
	}
	// meritcate := make([]*models.AdminMerit, 0)
	meritcate := make([]MeritCategory, 0)

	allmerits, err := models.GetAdminMerit(0)
	// beego.Info(allmerits)
	if err != nil {
		beego.Error(err)
	}
	var level string

	level = "2"
	for _, v1 := range allmerits {
		for _, v2 := range merits {
			if v2.MeritId == v1.Id {
				// beego.Info(v2.MeritId)
				level = "1"
				// merittitle, err := models.GetAdminMeritbyId(v2.MeritId) //因为这个数据库只是科室和分类的对应表
				// if err != nil {
				// 	beego.Error(err)
				// }
				// aa := make([]MeritCategory, 1)
				// aa[0].Id = merittitle.Id
				// aa[0].Title = merittitle.Title
				// aa[0].Level = "1"
				// meritcate = append(meritcate, aa...)
			}
		}
		aa := make([]MeritCategory, 1)
		aa[0].Id = v1.Id
		aa[0].Title = v1.Title
		aa[0].Level = level
		meritcate = append(meritcate, aa...)
		aa = make([]MeritCategory, 0)
		level = "2"
	}
	c.Data["json"] = meritcate
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//向科室id里添加价值分类
func (c *AdminMeritController) AddSecofficeMerit() {
	sid := c.GetString("sid") //secofficeid
	//id转成64位
	sidNum, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取出所有sidnum的merit
	merits, err := models.GetSecofficeMerit(sidNum)
	if err != nil {
		beego.Error(err)
	}

	ids := c.GetString("ids") //meritid
	array := strings.Split(ids, ",")
	bool := false
	for _, v1 := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v1, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		for _, v2 := range merits {
			//没有找到则插入
			if v2.MeritId == idNum {
				bool = true
			}
		}
		if bool == false {
			//存入数据库
			err = models.AddSecofficeMerit(sidNum, idNum)
			if err != nil {
				beego.Error(err)
			}
		}
		bool = false
	}

	for _, v3 := range merits {
		for _, v4 := range array {
			//id转成64位
			idNum, err := strconv.ParseInt(v4, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//没有找到则删除
			if v3.MeritId == idNum {
				bool = true
			}
		}

		if bool == false {
			//存入数据库
			err = models.DeleteSecofficeMerit(sidNum, v3.MeritId)
			if err != nil {
				beego.Error(err)
			}
		}
		bool = false
	}
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//添加价值结构
func (c *AdminMeritController) AddMerit() {
	pid := c.Input().Get("pid")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}
	title := c.Input().Get("title")
	mark := c.Input().Get("mark")
	// list := c.Input().Get("list")
	// listmark := c.Input().Get("listmark")
	var markint int
	if mark != "" {
		markint, err = strconv.Atoi(mark)
		if err != nil {
			beego.Error(err)
		}
	} else {
		markint = 0
	}
	//存入数据库
	_, err = models.AddAdminMerit(pidNum, title, markint)
	if err != nil {
		beego.Error(err)
	} else {
		data := title
		c.Ctx.WriteString(data)
	}
}

//修改
func (c *AdminMeritController) UpdateMerit() {
	title := c.Input().Get("title")
	mark := c.Input().Get("mark")
	var err error
	var markint int
	if mark != "" {
		markint, err = strconv.Atoi(mark)
		if err != nil {
			beego.Error(err)
		}
	} else {
		markint = 0
	}
	// list := c.Input().Get("list")
	// listmark := c.Input().Get("listmark")
	mid := c.Input().Get("mid")
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminMerit(midNum, title, markint)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改ratio" + mid)
		logs.Close()
	}
}

//删除
func (c *AdminMeritController) DeleteMerit() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	if ids == "" || len(array) == 0 {
		return
	}
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		if idNum == 0 {
			return
		}
		//查询下级，即分级
		categories, err := models.GetAdminMerit(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminMerit(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminMerit(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
			logs := logs.NewLogger(1000)
			logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
			logs.EnableFuncCallDepth(true)
			logs.Info(c.Ctx.Input.IP() + " " + "删除价值" + ids)
			logs.Close()
		}
	}
}

// @Title post import merittopic excel
// @Description post import merit excel
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 file not found
// @router /importmerittopics [post]
// 上传excel文件，格式：第一行，序号-名称-编号，都打上格子
func (c *AdminMeritController) ImportMeritTopics() {
	//获取上传的文件
	//获取上传的文件
	_, h, err := c.GetFile("usersexcel")
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(h.path)
	// var attachment string
	var path string

	// var filesize int64
	if h != nil {
		//保存附件
		path = "./attachment/" + h.Filename    // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("usersexcel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
			c.Data["json"] = "err保存文件失败"
			c.ServeJSON()
		} else {
			var user models.User
			//读出excel内容写入数据库
			xlFile, err := xlsx.OpenFile(path) //
			if err != nil {
				beego.Error(err)
			}

			for _, sheet := range xlFile.Sheets {
				for i, row := range sheet.Rows {
					if i != 0 {
						// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
						j := 1
						usernickname := row.Cells[j].String()
						// if err != nil {
						// 	beego.Error(err)
						// }
						user = models.GetUserByNickname(usernickname)

						meritcate_1 := row.Cells[j+1].String()
						// if err != nil {
						// 	beego.Error(err)
						// }
						meritcate_2 := row.Cells[j+2].String()
						meritcate_3 := row.Cells[j+3].String()
						merit, err := models.GetMeritIdbyTitles(meritcate_1, meritcate_2, meritcate_3)
						if err != nil {
							beego.Error(err)
						}

						title := row.Cells[j+4].String()
						content := row.Cells[j+5].String()
						active := row.Cells[j+6].String()
						// beego.Info(active)
						var activebool bool
						if active == "true" || active == "TRUE" {
							activebool = true
						} else {
							activebool = false
						}
						_, err = models.AddMerit(merit.Id, user.Id, title, content, activebool)
						if err != nil {
							beego.Error(err)
						}
					}
				}
			}
			//删除附件
			err = os.Remove(path)
			if err != nil {
				beego.Error(err)
			}
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}
