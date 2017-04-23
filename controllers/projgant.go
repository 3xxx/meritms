//项目进度控制器，具体任务控制器另外做吧
package controllers

import (
	// "encoding/json"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	// "github.com/astaxie/beego/utils/pagination"
	"meritms/models"
	"os"
	// "path"
	// "path/filepath"
	"strconv"
	"strings"
	"time"
)

type ProjGantController struct {
	beego.Controller
}

type Projectgant struct {
	Id     int64       `json:"id",form:"-"`
	Code   string      `json:"code",orm:"null"` //编号
	Name   string      `json:"name",orm:"null"` //项目-阶段合并一起
	Desc   string      `json:"desc",orm:"null"` //专业
	Values []Gantvalue `json:"values"`
}

type Gantvalue struct {
	Id          int64    `json:"id"`
	Label       string   `json:"label"` //标签
	Desc        string   `json:"desc"`
	CustomClass string   `json:"customClass"`
	DataObj     []string `json:"dataObj"` //['ha','ha2']
	Starttime   string   `json:"from"`
	Endtime     string   `json:"to"`
}

//项目列表页面
func (c *ProjGantController) Get() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	c.Data["IsProjectgant"] = true
	// beego.Info(c.Ctx.Input.IP())
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "cms/projects_gant.tpl"
}

//提供给项目列表页的table中json数据，扩展后按标签显示
func (c *ProjGantController) GetProjGants() {
	projgants, err := models.GetProjGants()
	if err != nil {
		beego.Error(err)
	}
	projectgant := make([]Projectgant, 0)
	// gantvalue := make([]Gantvalue, 0)//这个可以利用，需要循环，暂时没用
	var slice1 []string
	const lll = "2006-01-02"
	for i1, v1 := range projgants {
		aa := make([]Projectgant, 1)
		aa[0].Id = v1.Id
		aa[0].Code = v1.Code
		aa[0].Name = strconv.Itoa(i1+1) + " " + v1.Title + "-" + v1.DesignStage
		aa[0].Desc = v1.Section

		bb := make([]Gantvalue, 1)
		bb[0].Label = v1.Label
		bb[0].Desc = "<b>" + v1.Desc + "Task #</b>" + strconv.Itoa(i1+1) + "<br><b>Data</b>: [" + v1.Starttime.Format(lll) + "～" + v1.Endtime.Format(lll) + "]"
		bb[0].CustomClass = v1.CustomClass
		array := strings.Split(v1.DataObj, ",")
		for _, v2 := range array {
			cc := make([]string, 1)
			cc[0] = v2
			slice1 = append(slice1, cc...)
			cc = make([]string, 0)
		}
		bb[0].DataObj = slice1
		t1 := (v1.Starttime).UnixNano() / 1e6
		t2 := (v1.Endtime).UnixNano() / 1e6
		bb[0].Starttime = "/Date(" + strconv.FormatInt(t1, 10) + ")/"
		bb[0].Endtime = "/Date(" + strconv.FormatInt(t2, 10) + ")/"

		aa[0].Values = bb
		projectgant = append(projectgant, aa...)
		slice1 = make([]string, 0)
	}
	// {
	//     id:2,
	//     name: "珠三角",
	//     desc: "可研",

	//     values: [{
	//         from: "/Date(1492790400000)/",
	//         to: "/Date(1501257600000)/",
	//         desc: '<b>Task #</b>3<br><b>Data</b>: [2011-02-01 15:30:00 - 2011-02-01 16:00:00]',
	//         label: "label是啥",
	//         customClass: "ganttRed",
	//         dataObj: ['ha','ha2']
	//     }]
	// }
	c.Data["json"] = projectgant
	c.ServeJSON()
}

//根据id查看项目，查出项目目录
func (c *ProjGantController) GetProjectGant() {
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
	c.Data["IsProject"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role

	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据id取出下级
	cates := getsons(idNum, categories)
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)
	//递归生成目录json
	root := FileNode{category.Id, category.Title, "", []*FileNode{}}
	// walk(category.Id, &root)
	maketreejson(cates, categories, &root)
	// beego.Info(root)
	// data, _ := json.Marshal(root)
	c.Data["json"] = root //data
	// c.ServeJSON()
	c.Data["Category"] = category
	c.TplName = "cms/project.tpl"
}

//根据项目侧栏id查看这个id下的成果，不含子目录中的成果
//任何一级目录下都可以放成果
//这个作废——以product中的GetProducts
func (c *ProjGantController) GetProjGant() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	// c.Data["json"] = root
	// c.ServeJSON()
	c.TplName = "cms/project_products.tpl"
}

//添加项目和项目目录、文件夹
func (c *ProjGantController) AddProjGant() {
	// iprole := Getiprole(c.Ctx.Input.IP())
	// if iprole != 1 {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
	// rows := c.Input().Get("rows2[0][0]")
	// beego.Info(rows)
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	designstage := c.Input().Get("designstage")
	section := c.Input().Get("section")
	label := c.Input().Get("label")
	desc := c.Input().Get("desc")
	customclass := c.Input().Get("customclass")
	dataobj := c.Input().Get("dataobj")
	// datefilter := c.Input().Get("ddatefilter")
	daterange := c.Input().Get("datefilter")
	// beego.Info(daterange)
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	hours := 0
	var t1, t2 time.Time
	if len(daterange) > 19 {
		array := strings.Split(daterange, " - ")
		starttime1 := array[0]
		endtime1 := array[1]
		const lll = "2006-01-02"
		starttime, _ := time.Parse(lll, starttime1)
		endtime, _ := time.Parse(lll, endtime1)
		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
		// beego.Info(t1)：2016-08-19 00:00:00 +0000 UTC
		t2 = endtime.Add(-time.Duration(hours) * time.Hour)
		// beego.Info(t2)
	} else {
		t2 = time.Now()
		// beego.Info(t1):2016-08-19 23:27:29.7463081 +0800 CST
		// starttime, _ := time.Parse("2006-01-02", starttime1)
		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
		// beego.Info(t2)
	}

	_, err := models.AddProjGant(code, title, designstage, section, label, desc, customclass, dataobj, t1, t2)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = "ok"
	c.ServeJSON()
}

//导入甘特数据
//上传excel文件，导入到数据库
func (c *ProjGantController) ImportProjGant() {
	//获取上传的文件
	_, h, err := c.GetFile("gantsexcel")
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(h.path)
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		path = ".\\attachment\\" + h.Filename  // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("gantsexcel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}

	const lll = "2006-01-02"
	var convdate string
	var date time.Time
	var code, title, designstage, section, label, desc, customclass, dataobj string
	var t1, t2 time.Time
	//读出excel内容写入数据库
	xlFile, err := xlsx.OpenFile(path) //
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i != 0 { //忽略第一行标题
				// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
				j := 1

				if len(row.Cells) >= 2 { //总列数，从1开始
					code, err = row.Cells[j].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 3 {
					title, err = row.Cells[j+1].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 4 {
					designstage, err = row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 5 {
					section, err = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 6 {
					label, err = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}

				}
				if len(row.Cells) >= 7 {
					desc, err = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 8 {
					customclass, err = row.Cells[j+6].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 9 {
					dataobj, err = row.Cells[j+7].String()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 10 {
					if row.Cells[j+8].Value != "" {
						endtime2, err := row.Cells[j+8].Float()
						if err != nil {
							beego.Error(err)
						} else {
							date = xlsx.TimeFromExcelTime(endtime2, false)
						}
					} else {
						date = time.Now()
					}
					convdate = date.Format(lll)

					date, err = time.Parse(lll, convdate)
					if err != nil {
						beego.Error(err)
					}
					t1 = date
				}
				if len(row.Cells) >= 11 {
					if row.Cells[j+9].Value != "" {
						endtime2, err := row.Cells[j+9].Float()
						if err != nil {
							beego.Error(err)
						} else {
							date = xlsx.TimeFromExcelTime(endtime2, false)
						}
					} else {
						date = time.Now()
					}
					convdate = date.Format(lll)

					date, err = time.Parse(lll, convdate)
					if err != nil {
						beego.Error(err)
					}
					t2 = date
				}
				_, err := models.AddProjGant(code, title, designstage, section, label, desc, customclass, dataobj, t1, t2)
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

//修改项目名称、负责人等，
func (c *ProjGantController) UpdateProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	var err error
	projcode := c.Input().Get("code")
	projname := c.Input().Get("name")
	projlabe := c.Input().Get("label")
	principal := c.Input().Get("principal")
	pid := c.GetString("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateProject(idNum, projcode, projname, projlabe, principal)
	if err != nil {
		beego.Error(err)
	}

	if err != nil {
		c.Data["json"] = "no"
		c.ServeJSON()
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	// c.Data["json"] = "ok"
	// c.ServeJSON()
}

//根据id删除proj
//后台删除目录，
func (c *ProjGantController) DeleteProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// var err error
	//查所有子孙项目，循环删除
	ids := c.GetString("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//循环项目id
	for _, v := range array {
		//id转成64位
		projid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据项目id取得所有子孙id
		projs, err := models.GetProjectsbyPid(projid)
		if err != nil {
			beego.Error(err)
		}
		//循环子孙项目
		for _, w := range projs {
			//取得子孙项目的成果列表
			//根据项目id取得所有成果
			products, err := models.GetProducts(w.Id)
			if err != nil {
				beego.Error(err)
			}
			for _, x := range products {
				//删除子孙成果表
				//循环删除成果
				//根据成果id取得所有附件
				attachments, err := models.GetAttachments(x.Id)
				if err != nil {
					beego.Error(err)
				}
				//删除附件表
				for _, y := range attachments {
					//删除附件数据表
					err = models.DeleteAttachment(y.Id)
					if err != nil {
						beego.Error(err)
					}
				}

				//删除子孙文章表
				//取得成果id下所有文章
				articles, err := models.GetArticles(x.Id)
				if err != nil {
					beego.Error(err)
				}
				//删除文章表
				for _, z := range articles {
					//删除文章数据表
					err = models.DeleteArticle(z.Id)
					if err != nil {
						beego.Error(err)
					}
				}
				//删除成果表自身
				err = models.DeleteProduct(x.Id) //删除成果数据表
				if err != nil {
					beego.Error(err)
				}
			}
			//删除子孙proj数据表
			err = models.DeleteProject(w.Id)
			if err != nil {
				beego.Error(err)
			}
			//删除子孙文章图片文件夹（下面已经全部删除了）
		}
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(projid)
		if err != nil {
			beego.Error(err)
		} else {
			// beego.Info(DiskDirectory)
			path := DiskDirectory
			//直接删除这个文件夹，remove删除文件
			err = os.RemoveAll(path)
			if err != nil {
				beego.Error(err)
			}
			//删除项目自身数据表
			err = models.DeleteProject(projid)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	// if err != nil {
	// 	c.Data["json"] = "no"
	// 	c.ServeJSON()
	// } else {
	c.Data["json"] = "ok"
	c.ServeJSON()
	// }
}

//关闭项目进度
func (c *ProjGantController) CloseProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// var err error
	//查所有子孙项目，循环删除
	ids := c.GetString("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//循环项目id
	for _, v := range array {
		//id转成64位
		projid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据项目id取得所有子孙id
		projs, err := models.GetProjectsbyPid(projid)
		if err != nil {
			beego.Error(err)
		}
		//循环子孙项目
		for _, w := range projs {
			//取得子孙项目的成果列表
			//根据项目id取得所有成果
			products, err := models.GetProducts(w.Id)
			if err != nil {
				beego.Error(err)
			}
			for _, x := range products {
				//删除子孙成果表
				//循环删除成果
				//根据成果id取得所有附件
				attachments, err := models.GetAttachments(x.Id)
				if err != nil {
					beego.Error(err)
				}
				//删除附件表
				for _, y := range attachments {
					//删除附件数据表
					err = models.DeleteAttachment(y.Id)
					if err != nil {
						beego.Error(err)
					}
				}

				//删除子孙文章表
				//取得成果id下所有文章
				articles, err := models.GetArticles(x.Id)
				if err != nil {
					beego.Error(err)
				}
				//删除文章表
				for _, z := range articles {
					//删除文章数据表
					err = models.DeleteArticle(z.Id)
					if err != nil {
						beego.Error(err)
					}
				}
				//删除成果表自身
				err = models.DeleteProduct(x.Id) //删除成果数据表
				if err != nil {
					beego.Error(err)
				}
			}
			//删除子孙proj数据表
			err = models.DeleteProject(w.Id)
			if err != nil {
				beego.Error(err)
			}
			//删除子孙文章图片文件夹（下面已经全部删除了）
		}
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(projid)
		if err != nil {
			beego.Error(err)
		} else {
			// beego.Info(DiskDirectory)
			path := DiskDirectory
			//直接删除这个文件夹，remove删除文件
			err = os.RemoveAll(path)
			if err != nil {
				beego.Error(err)
			}
			//删除项目自身数据表
			err = models.DeleteProject(projid)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	// if err != nil {
	// 	c.Data["json"] = "no"
	// 	c.ServeJSON()
	// } else {
	c.Data["json"] = "ok"
	c.ServeJSON()
	// }
}
