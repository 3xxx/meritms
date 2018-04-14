package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/utils/pagination"
	"meritms/models"
	"path"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

type progresslist struct {
	Id      string `PK,json:"id"` //这个id必须是pk，select2才能选中，否则无法选中！！
	Name1   string `json:"name1"` //这些第一个字母要大写，否则不出结果
	Content string `json:"content"`
}

type progresslist1 struct {
	Id        string `PK,json:"id"`     //这个id必须是pk，select2才能选中，否则无法选中！！
	TaskTitle string `json:"tasktitle"` //任务名称
	Name1     string `json:"name1"`     //这些第一个字母要大写，否则不出结果
	Content   string `json:"content"`
	Name2     string `json:"name2"`
	Name3     string `json:"name3"`
	Path1     string `json:"path1"`
	Path2     string `json:"path2"`
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

func (c *MainController) ZSJ() {
	c.TplName = "zsj.tpl"
}

func (c *MainController) Progress() {
	c.TplName = "cms/progress.tpl"
}

func (c *MainController) Shower() {
	c.TplName = "cms/shower.tpl"
}

func (c *MainController) Pdf() {
	c.TplName = "web/viewer.html"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	// Directory:github.com/astaxie/beego/context     Pakage in Source:context
	// func (input *BeegoInput) IP() string {}
	//c是TopicController,TopicController是beego.controller，即beego.controller.ctx.input.ip

	// beego.Info(c.Ctx.Input.IP())

	//取得附件的id——成果的id——目录的id——查询目录下所有pdf返回数量
	// id := c.Input().Get("id")
	// if id != "" {
	// 	//id转成64为
	// 	idNum, err = strconv.ParseInt(id, 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	//取得某个目录下所有pdf
	// id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	// c.Data["Id"] = id
	var idNum int64
	// var err error
	var Url string
	// if id != "" {
	//id转成64为
	idNum, err := strconv.ParseInt("7", 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//由成果id（后台传过来的行id）取得侧栏目录id
	prod, err := models.GetProd(idNum)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得url
	Url, _, err = GetUrlPath(prod.ProjectId)
	if err != nil {
		beego.Error(err)
	}

	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
			linkarr := make([]AttachmentLink, 1)
			linkarr[0].Id = v.Id
			linkarr[0].Title = v.FileName
			linkarr[0].FileSize = v.FileSize
			linkarr[0].Downloads = v.Downloads
			linkarr[0].Created = v.Created
			linkarr[0].Updated = v.Updated
			linkarr[0].Link = Url + "/" + v.FileName
			link = append(link, linkarr...)
		}
	}
	// c.Data["json"] = link

	// pdfs, err := models.GetAllPdfs(idNum, false)
	// if err != nil {
	// 	beego.Error(err)
	// }
	count := len(Attachments)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// cnt, err := o.QueryTable("user").Count()
	// if err != nil {
	// 	beego.Error(err)
	// }

	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 1
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
	// fmt.Println(*c.Ctx)
	// beego.Info(c.Ctx)
	// beego.Info(paginator.Offset())   0
	// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
	// beego.Info(p.Offset())   0
	// fetch the next 20 posts
	// pdfs, err = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//这里根据上面取得的分页topics，取出这页的成果对应的所有项目
	// slice1 := make([]models.Category, 0)
	// for _, v := range topics {
	// 	tid := strconv.FormatInt(v.Id, 10)
	// 	category, err := models.GetTopicProj(tid)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// beego.Info(category.Title)
	// 	aa := make([]models.Category, 1)
	// 	aa[0].Author = category.Title//这句出错，不知何故runtime error: invalid memory address or nil pointer dereference
	// 	slice1 = append(slice1, aa...)
	// }
	// c.Data["Categories"] = slice1

	c.Data["Pdfs"] = Attachments

	c.Data["paginator"] = paginator
	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/test.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "ListAllTopic")
	// logs.Close()
	// count, _ := models.M("logoperation").Alias(`op`).Field(`count(op.id) as count`).Where(where).Count()
	// if count > 0 {
	// 	pagesize := 10
	// 	p := tools.NewPaginator(this.Ctx.Request, pagesize, count)
	// 	log, _ := models.M("logoperation").Alias(`op`).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
	// 	this.Data["data"] = log
	// 	this.Data["paginator"] = p
	// }
}

//这个是数据库里任务的真实数据
//这里的id数据确实有问题。
func (c *MainController) GetProgress() {
	data := []progresslist1{
		{"0", "枢纽布置图1/2", "0", "未启动", "0", "0", "dashed", "dashed"},
		{"0.125", "枢纽布置图2/2", "0.125", "完成1/8", "0", "0", "dashed", "dashed"},
		{"0.25", "结构布置图1/2", "0.25", "完成1/4", "0", "0", "dashed", "dashed"},
		{"1", "结构布置图2/2", "1", "任务完成", "0.125", "0", "solid", "dashed"},
		{"1", "厂房钢筋图1/2", "1", "任务完成", "1", "0.125", "solid", "solid"},
		{"0.625", "厂房钢筋图2/2", "0.625", "完成5/8", "0", "0", "dashed", "dashed"},
		{"0.75", "横断面图1/2", "0.75", "完成3/4", "0", "0", "dashed", "dashed"},
		{"0.875", "横断面图2/2", "0.875", "完成7/8", "0", "0", "dashed", "dashed"},
		{"1", "系统示意图", "1", "任务完成", "0.875", "0", "solid", "dashed"},
	}
	c.Data["json"] = data //string(b)
	c.ServeJSON()
}

//这个是下拉列表的数据
func (c *MainController) GetSelect() {
	data := []progresslist{
		{"0", "0", "未启动"},
		{"0.125", "0.125", "完成1/8"},
		{"0.25", "0.25", "完成1/4"},
		{"0.375", "0.375", "完成3/8"},
		{"0.5", "0.5", "完成1/2"},
		{"0.625", "0.625", "完成5/8"},
		{"0.75", "0.75", "完成3/4"},
		{"0.875", "0.875", "完成7/8"},
		{"1", "1", "任务完成"},
	}
	c.Data["json"] = data //string(b)
	c.ServeJSON()
}

func (c *MainController) GetProgress1() {
	name := c.Input().Get("query")
	var data progresslist
	switch name {
	case "0":
		data = progresslist{"0", "0", "未启动"}
	case "0.125":
		data = progresslist{"0.125", "0.125", "完成1/8"}
	case "0.25":
		data = progresslist{"0.25", "0.25", "完成1/4"}
	case "0.375":
		data = progresslist{"0.375", "0.375", "完成3/8"}
	case "0.5":
		data = progresslist{"0.5", "0.5", "完成1/2"}
	case "0.625":
		data = progresslist{"0.625", "0.625", "完成5/8"}
	case "0.75":
		data = progresslist{"0.75", "0.75", "完成3/4"}
	case "0.875":
		data = progresslist{"0.875", "0.875", "完成7/8"}
	case "1":
		data = progresslist{"1", "1", "任务完成"}
	}

	c.Data["json"] = data //string(b)
	c.ServeJSON()
}

func (c *MainController) ModifyProgress() {
	// c.Data["json"] = "ok" //string(b)
	// c.ServeJSON()
	// name := c.Input().Get("name")
	value := c.Input().Get("value")
	// pk := c.Input().Get("pk")
	// data := "未启动"
	// c.Ctx.WriteString(data)
	var data progresslist
	switch value {
	case "0":
		data = progresslist{"0", "0", "未启动"}
	case "0.125":
		data = progresslist{"0.125", "0.125", "完成1/8"}
	case "0.25":
		data = progresslist{"0.25", "0.25", "完成1/4"}
	case "0.375":
		data = progresslist{"0.375", "0.375", "完成3/8"}
	case "0.5":
		data = progresslist{"0.5", "0.5", "完成1/2"}
	case "0.625":
		data = progresslist{"0.625", "0.625", "完成5/8"}
	case "0.75":
		data = progresslist{"0.75", "0.75", "完成3/4"}
	case "0.875":
		data = progresslist{"0.875", "0.875", "完成7/8"}
	case "1":
		data = progresslist{"1", "1", "任务完成"}
	}

	c.Data["json"] = data //string(b)
	c.ServeJSON()
	// c.Ctx.WriteString(data)
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
