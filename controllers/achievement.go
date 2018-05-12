//在线成果登记
package controllers

import (
	"encoding/json"
	// "fmt"
	m "github.com/3xxx/meritms/models"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"github.com/3xxx/meritms/models"
	"github.com/astaxie/beego/logs"
	"sort"
	"strconv"
	"strings"
	"time"
)

//个人参与的项目列表
type Project struct {
	Id            int64
	ProjectNumber string //项目编号
	ProjectName   string //项目名称
	DesignStage   string //阶段
	Section       string //专业
	Value         float64
	Myvalue       float64
	Percent       float64
}

//Catalog添加附件链接和设计说明、校审意见
type CatalogLinkCont struct {
	Id            int64     `json:"id"`
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

//附件链接表
type CatalogLinkEditable struct {
	Id        int64
	CatalogId int64
	Url       string `orm:"sie(500)"`
	Editable  bool
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

//设计说明表
type CatalogContentlevel struct {
	Id        int64
	Content   string
	Title     string
	Editable  bool
	CatalogId int64
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

//struct排序
type graphictopics []m.Employeeachievement

func (list graphictopics) Len() int {
	return len(list)
}

func (list graphictopics) Less(i, j int) bool {
	if list[i].Sigma > list[j].Sigma {
		return true
	} else if list[i].Sigma < list[j].Sigma {
		return false
	} else {
		return list[i].Name > list[j].Name
	}
}

func (list graphictopics) Swap(i, j int) {
	var temp m.Employeeachievement = list[i]
	list[i] = list[j]
	list[j] = temp
}

type Userselect struct {
	Id   int64  //`json:"id"`
	Ad   string `json:"id"`
	Name string `json:"text"`
}

type Select1 struct {
	Title string `json:"title"`
}

type Achievement struct {
	beego.Controller
}

type AchEmployee struct { //员工姓名
	Id       int64  `json:"Id"` //`form:"-"`
	Pid      int64  `form:"-"`
	Nickname string `json:"text"` //这个是侧栏显示的内容
	Level    string `json:"Level"`
	Href     string `json:"href"`
}

type AchSecoffice struct { //专业室：水工、施工……
	Id         int64         `json:"Id"` //`form:"-"`
	Pid        int64         `form:"-"`
	Title      string        `json:"text"`
	Tags       [1]string     `json:"tags"` //显示员工数量
	Employee   []AchEmployee `json:"nodes"`
	Level      string        `json:"Level"`
	Href       string        `json:"href"`       //点击科室，显示总体情况
	Selectable bool          `json:"selectable"` //否则点击node，没反应，即默认false点击展开收缩
}

type AchDepart struct { //分院：施工预算、水工分院……
	Id int64 `json:"Id"` //`form:"-"`
	// Pid       int64          `form:"-"`
	Title      string         `json:"text"` //这个后面json仅仅对于encode解析有用
	Secoffice  []AchSecoffice `json:"nodes"`
	Level      string         `json:"Level"`
	Tags       [1]int         `json:"tags"` //显示员工数量
	Selectable bool           `json:"selectable"`
}

type Employee struct { //职员的分院和科室属性
	Id         int64  `form:"-"`
	Name       string `json:"Name"`
	Department string `json:"Department"` //分院
	Secoffice  string `json:"Keshi"`      //科室。当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //分值
	Marks      int    //记录个数
}

//管理员登录显示侧栏结构，方便查看科室里员工总体情况，以及查看员工个人详细
func (c *Achievement) GetAchievement() {
	// username, role := checkprodRole(c.Ctx)
	// if role == 1 {
	// 	c.Data["IsAdmin"] = true
	// } else if role >= 1 && role < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsAchievement"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	//1.首先判断是否登录
	// if !checkAccount(c.Ctx) {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	return
	// }
	//2.取得客户端用户名
	// var uname string
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// }
	//3.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	if roleint > 4 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	achemployee := make([]AchEmployee, 0)
	achsecoffice := make([]AchSecoffice, 0)
	achdepart := make([]AchDepart, 0)
	//由uname取得user,获得user的分院名称
	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}
	var depcount int
	// beego.Info(username)
	// beego.Info(role)
	switch role {
	case "1": //管理员登录显示的侧栏是全部的
		category1, err := models.GetAdminDepart(0) //得到多个分院
		if err != nil {
			beego.Error(err)
		}
		for i1, _ := range category1 {
			aa := make([]AchDepart, 1)
			aa[0].Id = category1[i1].Id
			aa[0].Level = "1"
			// aa[0].Pid = category[0].Id
			aa[0].Title = category1[i1].Title //分院名称
			// beego.Info(category1[i1].Title)
			category2, err := models.GetAdminDepart(category1[i1].Id) //得到多个科室
			if err != nil {
				beego.Error(err)
			}
			//如果返回科室为空，则直接取得员工
			//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
			//应该是反过来找出所有没有科室字段的人员，把他放在部门下
			if len(category2) > 0 {
				for i2, _ := range category2 {
					bb := make([]AchSecoffice, 1)
					bb[0].Id = category2[i2].Id
					bb[0].Level = "2"
					bb[0].Pid = category1[i1].Id
					bb[0].Title = category2[i2].Title //科室名称
					// beego.Info(category2[i2].Title)
					//根据分院和科室查所有员工
					users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
					if err != nil {
						beego.Error(err)
					}
					for i3, _ := range users {
						cc := make([]AchEmployee, 1)
						cc[0].Id = users[i3].Id
						cc[0].Level = "3"
						cc[0].Pid = category2[i2].Id
						cc[0].Nickname = users[i3].Nickname //名称
						// beego.Info(users[i3].Nickname)
						// cc[0].Selectable = false
						achemployee = append(achemployee, cc...)
					}
					bb[0].Tags[0] = strconv.Itoa(count)
					bb[0].Employee = achemployee
					bb[0].Selectable = true
					achemployee = make([]AchEmployee, 0) //再把slice置0
					achsecoffice = append(achsecoffice, bb...)
					depcount = depcount + count //部门人员数=科室人员数相加
				}
				// aa[0].Secoffice = achsecoffice
				// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				// achdepart = append(achdepart, aa...)
			}
			//查出所有有这个部门但科室名为空的人员
			//根据分院查所有员工
			// beego.Info(category1[i1].Title)
			users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(users)
			for i3, _ := range users {
				dd := make([]AchSecoffice, 1)
				dd[0].Id = users[i3].Id
				dd[0].Level = "3"
				// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
				dd[0].Pid = category1[i1].Id
				dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
				dd[0].Selectable = true
				achsecoffice = append(achsecoffice, dd...)
			}
			aa[0].Tags[0] = count + depcount
			// count = 0
			depcount = 0
			aa[0].Secoffice = achsecoffice
			aa[0].Selectable = true                //默认是false点击展开
			achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			achdepart = append(achdepart, aa...)
		}
	case "2": //分院领导登录显示的侧栏是本分院的所有科室
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		// aa[0].Selectable = false
		category2, err := models.GetAdminDepartTitle(user.Department) //得到多个科室
		if err != nil {
			beego.Error(err)
		}
		if len(category2) > 0 {
			for i2, _ := range category2 {
				bb := make([]AchSecoffice, 1)
				bb[0].Id = category2[i2].Id
				bb[0].Level = "2"
				bb[0].Pid = category1.Id
				bb[0].Title = category2[i2].Title //科室名称
				//根据分院和科室查所有员工
				users, count, err := models.GetUsersbySec(category1.Title, category2[i2].Title) //得到员工姓名
				if err != nil {
					beego.Error(err)
				}
				for i3, _ := range users {
					cc := make([]AchEmployee, 1)
					cc[0].Id = users[i3].Id
					cc[0].Level = "3"
					cc[0].Pid = category2[i2].Id
					cc[0].Nickname = users[i3].Nickname //名称
					// cc[0].Selectable = false
					achemployee = append(achemployee, cc...)
				}
				bb[0].Tags[0] = strconv.Itoa(count)
				bb[0].Employee = achemployee
				achemployee = make([]AchEmployee, 0) //再把slice置0
				achsecoffice = append(achsecoffice, bb...)
				depcount = depcount + count //部门人员数=科室人员数相加
			}
			// aa[0].Secoffice = achsecoffice
			// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			// achdepart = append(achdepart, aa...)
		}
		//查出所有有这个部门但科室名为空的人员
		//根据分院查所有员工
		users, count, err := models.GetUsersbySecOnly(category1.Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			dd := make([]AchSecoffice, 1)
			dd[0].Id = users[i3].Id
			dd[0].Level = "3"
			// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
			dd[0].Pid = category1.Id
			dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
			dd[0].Selectable = true
			achsecoffice = append(achsecoffice, dd...)
		}
		aa[0].Tags[0] = count + depcount
		// count = 0
		depcount = 0
		aa[0].Secoffice = achsecoffice
		aa[0].Selectable = false               //点击展开，默认是true
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	case "3": //主任登录显示的侧栏是本科室的所有人
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		//由分院id和科室名称取得科室
		category2, err := models.GetAdminDepartbyidtitle(category1.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		bb := make([]AchSecoffice, 1)
		bb[0].Id = category2.Id
		bb[0].Level = "2"
		bb[0].Pid = category1.Id
		bb[0].Title = category2.Title //科室名称
		//根据分院和科室查所有员工
		users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			cc := make([]AchEmployee, 1)
			cc[0].Id = users[i3].Id
			cc[0].Level = "3"
			cc[0].Pid = category2.Id
			cc[0].Nickname = users[i3].Nickname //名称
			// cc[0].Selectable = false
			achemployee = append(achemployee, cc...)
		}
		bb[0].Tags[0] = strconv.Itoa(count)
		bb[0].Employee = achemployee
		achemployee = make([]AchEmployee, 0) //再把slice置0
		achsecoffice = append(achsecoffice, bb...)
		aa[0].Secoffice = achsecoffice
		aa[0].Tags[0] = count
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	case "4": //个人登录显示自己
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		//由分院id和科室名称取得科室
		category2, err := models.GetAdminDepartbyidtitle(category1.Id, user.Secoffice)
		if err == nil { //== orm.ErrNoRows { // 没有找到记录 {
			bb := make([]AchSecoffice, 1)
			bb[0].Id = category2.Id
			bb[0].Level = "2"
			bb[0].Pid = category1.Id
			bb[0].Title = category2.Title //科室名称

			cc := make([]AchEmployee, 1)
			cc[0].Id = user.Id
			cc[0].Level = "3"
			cc[0].Pid = category2.Id
			cc[0].Nickname = user.Nickname //名称

			achemployee = append(achemployee, cc...)

			bb[0].Tags[0] = "1"
			bb[0].Employee = achemployee
			achemployee = make([]AchEmployee, 0) //再把slice置0
			achsecoffice = append(achsecoffice, bb...)
			aa[0].Secoffice = achsecoffice
			aa[0].Tags[0] = 1
			achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			achdepart = append(achdepart, aa...)
		} else {
			beego.Error(err)
		}
		//查出所有有这个部门但科室名为空的人员
		//根据分院查所有员工
		users, count, err := models.GetUsersbySecOnly(category1.Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			dd := make([]AchSecoffice, 1)
			dd[0].Id = users[i3].Id
			dd[0].Level = "3"
			// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
			dd[0].Pid = category1.Id
			dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
			dd[0].Selectable = true
			achsecoffice = append(achsecoffice, dd...)
		}
		aa[0].Tags[0] = count + depcount
		// count = 0
		depcount = 0
		aa[0].Secoffice = achsecoffice
		aa[0].Selectable = false               //点击展开，默认是true
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = achdepart
	// beego.Info(achdepart)
	c.TplName = "merit/achievement.tpl"
}

//下面关于用户名的取得需要重新来改改
//上面那个是显示侧栏
//这个是显示右侧iframe框架内容——科室内人员成果情况统计
func (c *Achievement) Secofficeshow() {
	// username, role := checkprodRole(c.Ctx)
	// if role == 1 {
	// 	c.Data["IsAdmin"] = true
	// } else if role >= 1 && role < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsAchievement"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	if roleint > 4 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}

	//由uname取得user
	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}

	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	if secid == "" { //如果为空，则用登录的
		secid = strconv.FormatInt(user.Id, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	level := c.Input().Get("level")
	key := c.Input().Get("key")
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

	switch level {
	case "0": //如果是总院，则显示全部分院情况
		// c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["Starttime"] = t1
		c.Data["Endtime"] = t2
		c.TplName = "merit/institute_show.tpl"
	case "1": //如果是分院，则显示全部科室
		categoryname, err := models.GetAdminDepartbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		//权限判断，并且属于这个分院
		if role == "1" || role == "2" && user.Department == categoryname.Title { //

			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["Deptitle"] = categoryname.Title
			c.TplName = "merit/achiev_depoffice.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	case "2": //如果是科室，则显示全部人员情况
		//取得科室名称
		categoryname, err := models.GetAdminDepartbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		// 取得分院名称
		categoryname1, err := models.GetAdminDepartbyId(categoryname.ParentId)
		if err != nil {
			beego.Error(err)
		}
		//1.进行权限读取,属于这个科室，或者属于这个分院
		if role == "1" || role == "3" && user.Secoffice == categoryname.Title || role == "2" && user.Department == categoryname1.Title { //

			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Sectitle"] = categoryname.Title
			c.Data["Level"] = level

			c.TplName = "merit/achiev_secoffice.tpl"
		} else {
			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	default:
		// case "3": //如果是个人，则显示个人详细情况
		//分2部分，一部分是已经完成状态的，state是4，另一部分是状态分别是3待审查通过,2，1的
		usernickname := models.GetUserByUserId(secid1)
		//1.进行权限读取，室主任以上并且属于这个科室，或者或本人
		if role == "1" || role == "3" && user.Secoffice == usernickname.Secoffice || role == "2" && user.Department == usernickname.Department || user.Nickname == usernickname.Nickname { //
			// employeecatalog := make([]models.Catalog, 0)
			//根据员工id和成果类型查出所有成果，设计成果，校核成果，审查成果
			//1、查图纸、报告……补充时间段secid即为userid
			//这里根据成果类型表循环查找
			//取得成果类型
			var slice1 []string
			ratios, err := models.GetAchievcategories()
			for _, v := range ratios {
				aa := make([]string, 1)
				aa[0] = v.Category //名称
				// cc[0].Selectable = false
				slice1 = append(slice1, aa...)
			}
			c.Data["Select2"] = slice1

			//查自己每个月的工作量和排名
			employeevalue := make([]models.Employeeachievement, 0)
			employeerealvalue := make([]models.Employeeachievement, 0)
			//查自己所在科室
			//根据分院和科室查所有员工
			users, _, err := models.GetUsersbySec(usernickname.Department, usernickname.Secoffice) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(users)
			//现在的月份
			const MMM = "2006-01"
			// date := time.Now()
			//	fmt.Println(date)
			month1 := time.Now().Format(MMM)
			// fmt.Println(convdate)
			month2, err := time.Parse(MMM, month1)
			if err != nil {
				beego.Error(err)
			}
			//全部工作量
			var slice2 []float64
			var slice3 []int
			var slice4 []time.Month
			//实物工作量
			var slice5 []float64
			var slice6 []int
			var slice7 []time.Month
			//往前12个月循环每个月
			for i4 := 0; i4 < 12; i4++ {
				for _, v3 := range users {
					//由username查出所有编制成果总分、设计总分……合计
					employee, employeereal, err := models.Getemployeevalue(v3.Nickname, month2.AddDate(0, -11+i4, 0), month2.AddDate(0, -11+i4+1, 0))
					if err != nil {
						beego.Error(err)
					}
					employeevalue = append(employeevalue, employee...)
					employeerealvalue = append(employeerealvalue, employeereal...)
				}
				// beego.Info(employeevalue)
				//排序
				pList := graphictopics(employeevalue)
				pListreal := graphictopics(employeerealvalue)
				sort.Sort(pList)
				sort.Sort(pListreal)
				// beego.Info(pList)
				for i5, v5 := range pList {
					if v5.Name == usernickname.Nickname {
						aa := make([]float64, 1)
						bb := make([]int, 1)
						cc := make([]time.Month, 1)
						aa[0] = v5.Sigma                             //工作量
						bb[0] = i5 + 1                               //排名
						cc[0] = month2.AddDate(0, -11+i4, 0).Month() //月份
						slice2 = append(slice2, aa...)
						slice3 = append(slice3, bb...)
						slice4 = append(slice4, cc...)
						break
					}
				}
				employeevalue = make([]models.Employeeachievement, 0)
				for i6, v6 := range pListreal {
					if v6.Name == usernickname.Nickname {
						dd := make([]float64, 1)
						ee := make([]int, 1)
						ff := make([]time.Month, 1)
						dd[0] = v6.Sigma                             //工作量
						ee[0] = i6 + 1                               //排名
						ff[0] = month2.AddDate(0, -11+i4, 0).Month() //月份
						slice5 = append(slice5, dd...)
						slice6 = append(slice6, ee...)
						slice7 = append(slice7, ff...)
						break
					}
				}
				employeerealvalue = make([]models.Employeeachievement, 0)
			}
			c.Data["Value1"] = slice2     //工作量
			c.Data["Value2"] = slice3     //排名
			c.Data["Value3"] = slice4     //月份
			c.Data["realValue1"] = slice5 //实物工作量
			c.Data["realValue2"] = slice6 //实物工作量排名
			c.Data["realValue3"] = slice7 //月份
			//参与的项目列表——检索、去重

			//根据userid得到所有成果,时间段，在模板里，根据catalogs的类型与category匹配进行显示即可
			//查出所有名单，传给json结构，再传给前端修改人名时选择
			var user22 models.User //这里修改[]*models.User(uname string)
			inputs := c.Input()
			user22.Username = inputs.Get("uname")
			uname1, err := models.GetUname(user22)

			if err != nil {
				beego.Error(err)
			}

			slice11 := make([]Userselect, 0)

			for _, v := range uname1 {
				aa := make([]Userselect, 1)
				aa[0].Id = v.Id //这里用for i1,v1,然后用v1.Id一样的意思
				aa[0].Ad = v.Nickname
				aa[0].Name = v.Nickname //v.Username + " " +
				slice11 = append(slice11, aa...)
			}
			if err != nil {
				beego.Error(err)
			}
			c.Data["Userselect"] = slice11

			catalogs, err := models.Getcatalog2byuserid(secid, t1, t2)
			if err != nil {
				beego.Error(err)
			}

			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Ratio"] = ratios //定义的成果类型
			//下面这个catalogs用于employee_show.tpl
			c.Data["Catalogs"] = catalogs
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["UserNickname"] = usernickname.Nickname

			if key == "modify" { //新窗口显示处理页面
				//如果是本人，则显示
				c.TplName = "merit/achiev_employeework.tpl"
			} else { //直接查看页面
				//如果是本人，则显示带处理按钮的
				if usernickname.Nickname == user.Nickname {
					c.Data["IsMe"] = true
				} else { //别人查看，不显示处理按钮
					c.Data["IsMe"] = false
				}
				c.TplName = "merit/achiev_employee.tpl"
			}
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	}
}

//上面那个是显示右侧页面
//这个是填充数据——科室内人员成果情况统计
func (c *Achievement) SecofficeData() {
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	level := c.Input().Get("level")
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
		beego.Info(t2)
	} else {
		t2 = time.Now()
		// beego.Info(t1):2016-08-19 23:27:29.7463081 +0800 CST
		// starttime, _ := time.Parse("2006-01-02", starttime1)
		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
		// beego.Info(t2)
	}

	//取得科室名称
	categoryname, err := models.GetAdminDepartbyId(secid1)
	if err != nil {
		beego.Error(err)
	}

	employeevalue := make([]models.Employeeachievement, 0)
	employeerealvalue := make([]models.Employeeachievement, 0)
	//根据科室id查所有员工
	users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
	// beego.Info(users)
	if err != nil {
		beego.Error(err)
	}
	for _, v := range users {
		//由username查出所有编制成果总分、设计总分……合计
		employee, employeereal, err := models.Getemployeevalue(v.Nickname, t1, t2)
		if err != nil {
			beego.Error(err)
		}
		employeevalue = append(employeevalue, employee...)
		employeerealvalue = append(employeerealvalue, employeereal...)
	}
	//排序
	pList := graphictopics(employeevalue)
	pListreal := graphictopics(employeerealvalue)
	sort.Sort(pList)
	sort.Sort(pListreal)
	c.Data["Starttime"] = t1
	c.Data["Endtime"] = t2
	c.Data["Secid"] = secid
	c.Data["Sectitle"] = categoryname.Title
	c.Data["Level"] = level
	c.Data["Employee"] = pList //employeevalue
	c.Data["json"] = pList
	c.ServeJSON()
}

//这个是填充数据——分院情况统计
//未修改吧——getpids已经不对了！！！
// func (c *Achievement) DepartData() {
// 	//1.首先判断是否注册
// 	if !checkAccount(c.Ctx) {
// 		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/login?url="+route, 302)
// 		// c.Redirect("/login", 302)
// 		return
// 	}
// 	//4.取得客户端用户名
// 	var uname string
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// uname := v.(string) //ck.Value
// 	//4.取出用户的权限等级
// 	role, _ := checkRole(c.Ctx) //login里的
// 	// beego.Info(role)
// 	//5.进行逻辑分析：
// 	// rolename, err := strconv.ParseInt(role, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }
// 	if role > 4 { //
// 		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/roleerr?url="+route, 302)
// 		// c.Redirect("/roleerr", 302)
// 		return
// 	}
// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid")
// 	if secid == "" { //如果为空，则用登录的
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	secid1, err := strconv.ParseInt(secid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	level := c.Input().Get("level")
// 	key := c.Input().Get("key")
// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 0
// 	var t1, t2 time.Time
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		const lll = "2006-01-02"
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
// 		// beego.Info(t1)：2016-08-19 00:00:00 +0000 UTC
// 		t2 = endtime.Add(-time.Duration(hours) * time.Hour)
// 		// beego.Info(t2)
// 	} else {
// 		t2 = time.Now()
// 		// beego.Info(t1):2016-08-19 23:27:29.7463081 +0800 CST
// 		// starttime, _ := time.Parse("2006-01-02", starttime1)
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t2)
// 	}

// 	switch level {
// 	case "0": //如果是总院，则显示全部分院情况
// 		// c.Data["IsLogin"] = checkAccount(c.Ctx)
// 		c.Data["Starttime"] = t1
// 		c.Data["Endtime"] = t2
// 		c.TplName = "institute_show.tpl"
// 	case "1": //如果是分院，则显示全部科室
// 		categoryname, err := models.GetAdminDepartbyId(secid1)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		//权限判断，并且属于这个分院
// 		if role == 1 || role == 2 && user.Department == categoryname.Title { //
// 			//根据分院id得到科室id
// 			//循环构造分院数据，view中进行循环显示各个科室情况
// 			Secofficevalue := make([]models.Secofficeachievement, 0)
// 			Secofficerealvalue := make([]models.Secofficeachievement, 0)
// 			category, err := models.GetPids(secid1) //得到多个科室
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			employeevalue := make([]models.Employeeachievement, 0)
// 			employeerealvalue := make([]models.Employeeachievement, 0)
// 			for _, v1 := range category {
// 				aa := make([]models.Secofficeachievement, 1)
// 				realvalue := make([]models.Secofficeachievement, 1)
// 				//根据科室id查所有员工
// 				secid2 := strconv.FormatInt(v1.Id, 10)
// 				users, _, err := models.GetUsersbySecId(secid2) //得到员工姓名
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				for _, v := range users {
// 					//由username查出所有编制成果总分、设计总分……合计
// 					employee, employeereal, err := models.Getemployeevalue(v.Nickname, t1, t2)
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 					employeevalue = append(employeevalue, employee...)
// 					employeerealvalue = append(employeerealvalue, employeereal...)
// 				}
// 				//排序
// 				pList := graphictopics(employeevalue)
// 				pListreal := graphictopics(employeerealvalue)
// 				sort.Sort(pList)
// 				sort.Sort(pListreal)
// 				aa[0].Id = v1.Id //科室Id
// 				realvalue[0].Id = v1.Id
// 				aa[0].Name = v1.Title //科室名称
// 				realvalue[0].Name = v1.Title
// 				aa[0].Employee = pList //employeevalue
// 				realvalue[0].Employee = pListreal
// 				Secofficevalue = append(Secofficevalue, aa...)
// 				Secofficerealvalue = append(Secofficerealvalue, realvalue...)
// 				aa = make([]models.Secofficeachievement, 0) //再把slice置0
// 				realvalue = make([]models.Secofficeachievement, 0)
// 				employeevalue = make([]models.Employeeachievement, 0)
// 				employeerealvalue = make([]models.Employeeachievement, 0)
// 			}
// 			c.Data["Starttime"] = t1
// 			c.Data["Endtime"] = t2
// 			c.Data["Secid"] = secid
// 			c.Data["Level"] = level
// 			c.Data["Secoffice"] = Secofficevalue
// 			c.Data["Deptitle"] = categoryname.Title
// 			c.TplName = "depoffice_show.tpl"
// 		} else {
// 			route := c.Ctx.Request.URL.String()
// 			c.Data["Url"] = route
// 			c.Redirect("/roleerr?url="+route, 302)
// 			// c.Redirect("/roleerr", 302)
// 			return
// 		}
// 	case "2": //如果是科室，则显示全部人员情况
// 		//取得科室名称
// 		categoryname, err := models.GetAdminDepartbyId(secid1)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		// 取得分院名称
// 		categoryname1, err := models.GetAdminDepartbyId(categoryname.ParentId)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		//1.进行权限读取,属于这个科室，或者属于这个分院
// 		if role == 1 || role == 3 && user.Secoffice == categoryname.Title || role == 2 && user.Department == categoryname1.Title { //
// 			employeevalue := make([]models.Employeeachievement, 0)
// 			employeerealvalue := make([]models.Employeeachievement, 0)
// 			// depid := c.Input().Get("depid")
// 			//根据分院和科室查所有员工
// 			// users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
// 			//根据科室id查所有员工
// 			users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
// 			// beego.Info(users)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			for _, v := range users {
// 				//由username查出所有编制成果总分、设计总分……合计
// 				employee, employeereal, err := models.Getemployeevalue(v.Nickname, t1, t2)
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				employeevalue = append(employeevalue, employee...)
// 				employeerealvalue = append(employeerealvalue, employeereal...)
// 			}
// 			//排序
// 			pList := graphictopics(employeevalue)
// 			pListreal := graphictopics(employeerealvalue)
// 			sort.Sort(pList)
// 			sort.Sort(pListreal)
// 			c.Data["Starttime"] = t1
// 			c.Data["Endtime"] = t2
// 			c.Data["Secid"] = secid
// 			c.Data["Sectitle"] = categoryname.Title
// 			c.Data["Level"] = level
// 			c.Data["Employee"] = pList //employeevalue
// 			c.Data["json"] = pList
// 			c.ServeJSON()
// 			c.Data["Employeereal"] = pListreal
// 			c.TplName = "secoffice_show.tpl"
// 		} else {
// 			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 			route := c.Ctx.Request.URL.String()
// 			c.Data["Url"] = route
// 			c.Redirect("/roleerr?url="+route, 302)
// 			// c.Redirect("/roleerr", 302)
// 			return
// 		}
// 	default:
// 		// case "3": //如果是个人，则显示个人详细情况
// 		//分2部分，一部分是已经完成状态的，state是4，另一部分是状态分别是3待审查通过,2，1的
// 		usernickname := models.GetUserByUserId(secid1)
// 		//1.进行权限读取，室主任以上并且属于这个科室，或者或本人
// 		if role == 1 || role == 3 && user.Secoffice == usernickname.Secoffice || role == 2 && user.Department == usernickname.Department || user.Nickname == usernickname.Nickname { //
// 			// employeecatalog := make([]models.Catalog, 0)
// 			//根据员工id和成果类型查出所有成果，设计成果，校核成果，审查成果
// 			//1、查图纸、报告……补充时间段secid即为userid
// 			//这里根据成果类型表循环查找
// 			//取得成果类型
// 			var slice1 []string
// 			ratios, err := models.GetAchievcategories()
// 			for _, v := range ratios {
// 				aa := make([]string, 1)
// 				aa[0] = v.Category //名称
// 				// cc[0].Selectable = false
// 				slice1 = append(slice1, aa...)
// 			}
// 			c.Data["Select2"] = slice1

// 			//查自己每个月的工作量和排名
// 			employeevalue := make([]models.Employeeachievement, 0)
// 			employeerealvalue := make([]models.Employeeachievement, 0)
// 			//查自己所在科室
// 			//根据分院和科室查所有员工
// 			users, _, err := models.GetUsersbySec(usernickname.Department, usernickname.Secoffice) //得到员工姓名
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			// beego.Info(users)
// 			//现在的月份
// 			const MMM = "2006-01"
// 			// date := time.Now()
// 			//	fmt.Println(date)
// 			month1 := time.Now().Format(MMM)
// 			// fmt.Println(convdate)
// 			month2, err := time.Parse(MMM, month1)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			//全部工作量
// 			var slice2 []float64
// 			var slice3 []int
// 			var slice4 []time.Month
// 			//实物工作量
// 			var slice5 []float64
// 			var slice6 []int
// 			var slice7 []time.Month
// 			//往前12个月循环每个月
// 			for i4 := 0; i4 < 12; i4++ {
// 				for _, v3 := range users {
// 					//由username查出所有编制成果总分、设计总分……合计
// 					employee, employeereal, err := models.Getemployeevalue(v3.Nickname, month2.AddDate(0, -11+i4, 0), month2.AddDate(0, -11+i4+1, 0))
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 					employeevalue = append(employeevalue, employee...)
// 					employeerealvalue = append(employeerealvalue, employeereal...)
// 				}
// 				// beego.Info(employeevalue)
// 				//排序
// 				pList := graphictopics(employeevalue)
// 				pListreal := graphictopics(employeerealvalue)
// 				sort.Sort(pList)
// 				sort.Sort(pListreal)
// 				// beego.Info(pList)
// 				for i5, v5 := range pList {
// 					if v5.Name == usernickname.Nickname {
// 						aa := make([]float64, 1)
// 						bb := make([]int, 1)
// 						cc := make([]time.Month, 1)
// 						aa[0] = v5.Sigma                             //工作量
// 						bb[0] = i5 + 1                               //排名
// 						cc[0] = month2.AddDate(0, -11+i4, 0).Month() //月份
// 						slice2 = append(slice2, aa...)
// 						slice3 = append(slice3, bb...)
// 						slice4 = append(slice4, cc...)
// 						break
// 					}
// 				}
// 				employeevalue = make([]models.Employeeachievement, 0)
// 				for i6, v6 := range pListreal {
// 					if v6.Name == usernickname.Nickname {
// 						dd := make([]float64, 1)
// 						ee := make([]int, 1)
// 						ff := make([]time.Month, 1)
// 						dd[0] = v6.Sigma                             //工作量
// 						ee[0] = i6 + 1                               //排名
// 						ff[0] = month2.AddDate(0, -11+i4, 0).Month() //月份
// 						slice5 = append(slice5, dd...)
// 						slice6 = append(slice6, ee...)
// 						slice7 = append(slice7, ff...)
// 						break
// 					}
// 				}
// 				employeerealvalue = make([]models.Employeeachievement, 0)
// 			}
// 			c.Data["Value1"] = slice2     //工作量
// 			c.Data["Value2"] = slice3     //排名
// 			c.Data["Value3"] = slice4     //月份
// 			c.Data["realValue1"] = slice5 //实物工作量
// 			c.Data["realValue2"] = slice6 //实物工作量排名
// 			c.Data["realValue3"] = slice7 //月份
// 			//参与的项目列表——检索、去重

// 			//根据userid得到所有成果,时间段，在模板里，根据catalogs的类型与category匹配进行显示即可
// 			//查出所有名单，传给json结构，再传给前端修改人名时选择
// 			var user22 models.User //这里修改[]*models.User(uname string)
// 			inputs := c.Input()
// 			user22.Username = inputs.Get("uname")
// 			uname1, err := models.GetUname(user22)

// 			if err != nil {
// 				beego.Error(err)
// 			}

// 			slice11 := make([]Userselect, 0)

// 			for _, v := range uname1 {
// 				aa := make([]Userselect, 1)
// 				aa[0].Id = v.Id //这里用for i1,v1,然后用v1.Id一样的意思
// 				aa[0].Ad = v.Nickname
// 				aa[0].Name = v.Nickname //v.Username + " " +
// 				slice11 = append(slice11, aa...)
// 			}
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			c.Data["Userselect"] = slice11

// 			catalogs, err := models.Getcatalog2byuserid(secid, t1, t2)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			c.Data["Starttime"] = t1
// 			c.Data["Endtime"] = t2
// 			c.Data["Ratio"] = ratios //定义的成果类型
// 			//下面这个catalogs用于employee_show.tpl
// 			c.Data["Catalogs"] = catalogs
// 			c.Data["Secid"] = secid
// 			c.Data["Level"] = level
// 			c.Data["UserNickname"] = usernickname.Nickname

// 			if key == "modify" { //新窗口显示处理页面
// 				//如果是本人，则显示
// 				// if usernickname.Nickname == user.Nickname {
// 				c.TplName = "employeeselfmodify_show.tpl"
// 				// } else { //别人查看，只能显示结果页面，不显示处理页面
// 				// 	c.TplName = "employee_show2.tpl"
// 				// }
// 			} else { //直接查看页面
// 				//如果是本人，则显示带处理按钮的
// 				if usernickname.Nickname == user.Nickname {
// 					c.Data["IsMe"] = true
// 				} else { //别人查看，不显示处理按钮
// 					// c.TplName = "employee_show2.tpl"
// 					c.Data["IsMe"] = false
// 				}
// 				c.TplName = "employee_show.tpl"
// 			}
// 		} else {
// 			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 			route := c.Ctx.Request.URL.String()
// 			c.Data["Url"] = route
// 			c.Redirect("/roleerr?url="+route, 302)
// 			// c.Redirect("/roleerr", 302)
// 			return
// 		}
// 	}
// }

//20170608修改
//显示登录用户待提交，处于设计，校核，审查，已提交，已经完成的
//author=登录的人名，登录名所处制图-状态为1；设计-状态为2；校核-状态为3；审查-状态为4；已经完成，6；已经提交，5
func (c *Achievement) AchievementSend() {
	// username, role := checkprodRole(c.Ctx)
	// if role == 1 {
	// 	c.Data["IsAdmin"] = true
	// } else if role >= 1 && role < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsAchievement"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	if roleint > 4 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}

	id := c.Ctx.Input.Param(":id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		beego.Error(err)
	}

	//如果是主任以上权限人查看，则id代表用户名id，个人查看，id则代表价值id
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}

	//1.取得客户端用户名
	// var uname string
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// }
	//由uname取得user
	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}

	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid") //要查看的用户id

	if secid == "" { //自己登录直接显示自己
		secid = strconv.FormatInt(user.Id, 10)
	}

	daterange := c.Input().Get("datefilter")

	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	hours := 8
	var t1, t2 time.Time
	// var convdate1, convdate2 string
	const lll = "2006-01-02"
	if len(daterange) > 19 {
		array := strings.Split(daterange, " - ")
		starttime1 := array[0]
		endtime1 := array[1]
		starttime, _ := time.Parse(lll, starttime1)
		endtime, _ := time.Parse(lll, endtime1)
		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
		t2 = endtime.Add(+time.Duration(16) * time.Hour) //因为数据库存的时间晚8小时，所以整个时间段像前退8小时，但是，最后一天加24小时，所以是16
	} else {
		t2 = time.Now()
		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
	}
	//取得成果类型
	ratios, err := models.GetAchievcategories()
	if err != nil {
		beego.Error(err)
	}
	var select2 string
	for i, v := range ratios {
		if i == 0 {
			select2 = v.Category
		} else {
			select2 = select2 + "," + v.Category
		}
	}

	var catalogs []*models.Catalog
	if idint == 1 {
		//取得这个id下的所有merittopic
		catalogs, err = models.GetcatalogMyself(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 2 {
		catalogs, err = models.GetcatalogDesignd(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 3 {
		catalogs, err = models.GetcatalogChecked(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 4 {
		catalogs, err = models.GetcatalogExamined(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 5 {
		catalogs, err = models.GetcatalogRunning(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 6 {
		catalogs, err = models.GetcatalogCompleted(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
	}

	link := make([]CatalogLinkCont, 0)
	Attachslice := make([]models.CatalogLink, 0)
	Contentslice := make([]models.CatalogContent, 0)
	linkarr := make([]CatalogLinkCont, 1)
	attacharr := make([]models.CatalogLink, 1)
	contarr := make([]models.CatalogContent, 1)

	//这里循环，添加附件链接和设计说，校审意见
	for _, w := range catalogs {
		// linkarr[0].Catalog = *w
		linkarr[0].Id = w.Id //table必须有这个id，否则不能删除某行
		linkarr[0].ProjectNumber = w.ProjectNumber
		linkarr[0].ProjectName = w.ProjectName
		linkarr[0].DesignStage = w.DesignStage
		linkarr[0].Section = w.Section
		linkarr[0].Tnumber = w.Tnumber
		linkarr[0].Name = w.Name
		linkarr[0].Category = w.Category
		linkarr[0].Page = w.Page
		linkarr[0].Count = w.Count
		linkarr[0].Drawn = w.Drawn
		linkarr[0].Designd = w.Designd
		linkarr[0].Checked = w.Checked
		linkarr[0].Examined = w.Examined
		linkarr[0].Verified = w.Verified
		linkarr[0].Approved = w.Approved
		linkarr[0].Complex = w.Complex
		linkarr[0].Drawnratio = w.Drawnratio
		linkarr[0].Designdratio = w.Designdratio
		linkarr[0].Checkedratio = w.Checkedratio
		linkarr[0].Examinedratio = w.Examinedratio
		linkarr[0].Datestring = w.Datestring
		linkarr[0].Date = w.Date
		linkarr[0].Created = w.Created
		linkarr[0].Updated = w.Updated
		linkarr[0].Author = w.Author
		linkarr[0].State = w.State
		links, err := models.GetCatalogLinks(w.Id)
		if err != nil {
			beego.Error(err)
		}
		for _, v := range links {
			attacharr[0].Url = v.Url
			// beego.Info(v.Url)
			Attachslice = append(Attachslice, attacharr...)
		}

		contents, err := models.GetCatalogContents(w.Id)
		if err != nil {
			beego.Error(err)
		}
		for _, v := range contents {
			contarr[0].Content = v.Content
			Contentslice = append(Contentslice, contarr...)
		}
		linkarr[0].Link = Attachslice
		linkarr[0].Content = Contentslice

		Attachslice = make([]models.CatalogLink, 0)
		Contentslice = make([]models.CatalogContent, 0)
		link = append(link, linkarr...)
	}
	c.Data["Select2"] = select2
	c.Data["Starttime"] = t1
	c.Data["Endtime"] = t2
	c.Data["Ratio"] = ratios //定义的成果类型

	c.Data["UserNickname"] = user.Nickname
	c.Data["json"] = link //catalogs
	c.ServeJSON()
}

//列表显示成果附件
func (c *Achievement) CatalogAttachment() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	// var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	//由id取得成果状态
	catalog, err := m.GetCatalog(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	links, err := models.GetCatalogLinks(idNum)
	if err != nil {
		beego.Error(err)
	}

	Attachslice := make([]CatalogLinkEditable, 0)
	attacharr := make([]CatalogLinkEditable, 1)
	if len(links) > 0 {
		for _, v := range links {
			attacharr[0].Id = v.Id
			// linkarr[0].Title = v.FileName
			attacharr[0].Url = v.Url
			attacharr[0].CatalogId = idNum
			if catalog.State == 1 || catalog.State == 2 {
				attacharr[0].Editable = true
			} else {
				attacharr[0].Editable = false
			}
			// beego.Info(v.Url)
			Attachslice = append(Attachslice, attacharr...)
		}
		if catalog.State == 1 || catalog.State == 2 {
			attacharr[0].Url = "http://"
			attacharr[0].Id = 0
			attacharr[0].CatalogId = idNum
			attacharr[0].Editable = true
			Attachslice = append(Attachslice, attacharr...)
		}
	} else {
		if catalog.State == 1 || catalog.State == 2 {
			attacharr[0].Created = time.Now()
			attacharr[0].Updated = time.Now()
			attacharr[0].Editable = true
			attacharr[0].Url = "http://"
			attacharr[0].CatalogId = idNum
			Attachslice = attacharr
		}
	}

	c.Data["json"] = Attachslice
	c.ServeJSON()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

//列表显示成果的校审意见
func (c *Achievement) CatalogContent() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	// var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	//由id取得成果状态
	catalog, err := m.GetCatalog(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有意见
	contents, err := models.GetCatalogContents(idNum)
	if err != nil {
		beego.Error(err)
	}

	Contentslice := make([]CatalogContentlevel, 3)
	// contentarr := make([]CatalogContentlevel, 1)
	for _, v := range contents {
		// attacharr[0].Id = v.Id
		if v.Level == 1 {
			Contentslice[0].Title = "设计说明"
			Contentslice[0].Id = v.Id
			Contentslice[0].CatalogId = v.CatalogId
			Contentslice[0].Content = v.Content
			Contentslice[0].Created = v.Created
			Contentslice[0].Updated = v.Updated
			if catalog.State == 1 || catalog.State == 2 {
				Contentslice[0].Editable = true
			} else {
				Contentslice[0].Editable = false
			}
		} else if v.Level == 2 {
			Contentslice[1].Title = "校核意见"
			Contentslice[1].Id = v.Id
			Contentslice[1].CatalogId = v.CatalogId
			Contentslice[1].Content = v.Content
			Contentslice[1].Created = v.Created
			Contentslice[1].Updated = v.Updated
			if catalog.State == 3 {
				Contentslice[1].Editable = true
			} else {
				Contentslice[1].Editable = false
			}
		} else if v.Level == 3 {
			Contentslice[2].Title = "审查意见"
			Contentslice[2].Id = v.Id
			Contentslice[2].CatalogId = v.CatalogId
			Contentslice[2].Content = v.Content
			Contentslice[2].Created = v.Created
			Contentslice[2].Updated = v.Updated
			if catalog.State == 4 {
				Contentslice[2].Editable = true
			} else {
				Contentslice[2].Editable = false
			}
		}
		// Contentslice = append(Contentslice, contentarr...)
	}
	if Contentslice[0].Title != "设计说明" {
		Contentslice[0].Title = "设计说明"
		Contentslice[0].Id = 0
		Contentslice[0].CatalogId = idNum
		if catalog.State == 1 || catalog.State == 2 {
			Contentslice[0].Editable = true
		} else {
			Contentslice[0].Editable = false
		}
	}
	if Contentslice[1].Title != "校核意见" {
		Contentslice[1].Title = "校核意见"
		Contentslice[1].Id = 0
		Contentslice[1].CatalogId = idNum
		if catalog.State == 3 {
			Contentslice[1].Editable = true
		} else {
			Contentslice[1].Editable = false
		}
	}
	if Contentslice[2].Title != "审查意见" {
		Contentslice[2].Title = "审查意见"
		Contentslice[2].Id = 0
		Contentslice[2].CatalogId = idNum
		if catalog.State == 4 {
			Contentslice[2].Editable = true
		} else {
			Contentslice[2].Editable = false
		}
	}
	c.Data["json"] = Contentslice
	c.ServeJSON()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

//修改link
func (c *Achievement) ModifyLink() {
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")

	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	cid := c.Input().Get("cid") //成果id
	cidnum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//无论如何都修改当前，重复修改了
	err = m.ModifyCatalogLink(id, cidnum, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := value
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + pk)
	logs.Close()
}

//修改校审意见
func (c *Achievement) ModifyContent() {
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")   //意见id
	cid := c.Input().Get("cid") //成果id
	// var cidnum int64
	var level int
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// if cid != "" {
	cidnum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// } else {
	// 	cidnum = 0
	// }
	//根据成果state，确定意见的level，如果是1和2，level=1,3对应level=2,4对应level=3
	catalog, err := models.GetCatalog(cidnum)
	if err != nil {
		beego.Error(err)
	}
	if catalog.State == 1 || catalog.State == 2 {
		level = 1
	} else if catalog.State == 3 {
		level = 2
	} else if catalog.State == 4 {
		level = 3
	}
	//无论如何都修改当前，重复修改了
	err = m.ModifyCatalogContent(id, cidnum, name, value, level)
	if err != nil {
		beego.Error(err)
	} else {
		data := value
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + pk)
	logs.Close()
}

//自己发起的成果,已经提交
//author=登录人名，状态>登录名字所处位置，且状态<5
// func (c *Achievement) Running() {
// 	//1.首先判断是否注册

// 	//4.取得客户端用户名
// 	var uname string
// 	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := c.GetSession("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid") //用户id
// 	if secid == "" {                //自己登录直接显示自己
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	// secid1, err := strconv.ParseInt(secid, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	// level := c.Input().Get("level")

// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 8
// 	var t1, t2 time.Time
// 	// var convdate1, convdate2 string
// 	const lll = "2006-01-02"
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
// 		t2 = endtime.Add(+time.Duration(16) * time.Hour)
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	} else {
// 		t2 = time.Now()
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	}
// 	// convdate1 = t1.Format(lll)
// 	// convdate2 = t2.Format(lll)
// 	// usernickname := models.GetUserByUserId(secid1)
// 	ratios, err := models.GetAchievcategories()
// 	var select2 string
// 	for i, v := range ratios {
// 		if i == 0 {
// 			select2 = v.Category
// 		} else {
// 			select2 = select2 + "," + v.Category
// 		}
// 	}
// 	// beego.Info(select2)
// 	catalogs, err := models.GetcatalogRunning(secid, t1, t2)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["Select2"] = select2
// 	c.Data["Starttime"] = t1
// 	c.Data["Endtime"] = t2
// 	c.Data["Ratio"] = ratios //定义的成果类型
// 	// c.Data["Secid"] = secid
// 	// c.Data["Level"] = level
// 	c.Data["UserNickname"] = user.Nickname
// 	c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//自己已经完成的成果
//制图、设计、校核、审查中含有登录名字，状态为5
// func (c *Achievement) Completed() {
// 	//1.首先判断是否注册

// 	//4.取得客户端用户名
// 	var uname string
// 	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := c.GetSession("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid") //用户id
// 	if secid == "" {                //自己登录直接显示自己
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	// secid1, err := strconv.ParseInt(secid, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	// level := c.Input().Get("level")

// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 24
// 	var t1, t2 time.Time
// 	// var convdate1, convdate2 string
// 	const lll = "2006-01-02"
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime //.Add(-time.Duration(hours) * time.Hour)
// 		t2 = endtime.Add(+time.Duration(hours) * time.Hour)
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	} else {
// 		t2 = time.Now()
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	}
// 	// convdate1 = t1.Format(lll)
// 	// convdate2 = t2.Format(lll)
// 	// usernickname := models.GetUserByUserId(secid1)
// 	ratios, err := models.GetAchievcategories()
// 	var select2 string
// 	for i, v := range ratios {
// 		if i == 0 {
// 			select2 = v.Category
// 		} else {
// 			select2 = select2 + "," + v.Category
// 		}
// 	}
// 	// beego.Info(select2)
// 	catalogs, err := models.GetcatalogCompleted(secid, t1, t2)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["Select2"] = select2
// 	c.Data["Starttime"] = t1
// 	c.Data["Endtime"] = t2
// 	c.Data["Ratio"] = ratios //定义的成果类型
// 	// c.Data["Secid"] = secid
// 	// c.Data["Level"] = level
// 	c.Data["UserNickname"] = user.Nickname
// 	c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//别人传来，自己处于设计位置的展示
//设计人名=登录名，状态2，author！=登录名
// func (c *Achievement) Designd() {
// 	//1.首先判断是否注册

// 	//4.取得客户端用户名
// 	var uname string
// 	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := c.GetSession("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid") //用户id
// 	if secid == "" {                //自己登录直接显示自己
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	// secid1, err := strconv.ParseInt(secid, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	// level := c.Input().Get("level")

// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 8
// 	var t1, t2 time.Time
// 	// var convdate1, convdate2 string
// 	const lll = "2006-01-02"
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
// 		t2 = endtime.Add(+time.Duration(16) * time.Hour)
// 	} else {
// 		t2 = time.Now()
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	}
// 	// convdate1 = t1.Format(lll)
// 	// convdate2 = t2.Format(lll)
// 	// usernickname := models.GetUserByUserId(secid1)
// 	ratios, err := models.GetAchievcategories()
// 	var select2 string
// 	for i, v := range ratios {
// 		if i == 0 {
// 			select2 = v.Category
// 		} else {
// 			select2 = select2 + "," + v.Category
// 		}
// 	}
// 	// beego.Info(select2)
// 	catalogs, err := models.GetcatalogDesignd(secid, t1, t2)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	link := make([]CatalogLinkCont, 0)
// 	Attachslice := make([]models.CatalogLink, 0)
// 	Contentslice := make([]models.CatalogContent, 0)
// 	linkarr := make([]CatalogLinkCont, 1)
// 	attacharr := make([]models.CatalogLink, 1)
// 	contarr := make([]models.CatalogContent, 1)

// 	//这里循环，添加附件链接和设计说，校审意见
// 	for _, w := range catalogs {
// 		linkarr[0].Id = w.Id
// 		linkarr[0].ProjectNumber = w.ProjectNumber
// 		linkarr[0].ProjectName = w.ProjectName
// 		linkarr[0].DesignStage = w.DesignStage
// 		linkarr[0].Section = w.Section
// 		linkarr[0].Tnumber = w.Tnumber
// 		linkarr[0].Name = w.Name
// 		linkarr[0].Category = w.Category
// 		linkarr[0].Page = w.Page
// 		linkarr[0].Count = w.Count
// 		linkarr[0].Drawn = w.Drawn
// 		linkarr[0].Designd = w.Designd
// 		linkarr[0].Checked = w.Checked
// 		linkarr[0].Examined = w.Examined
// 		linkarr[0].Verified = w.Verified
// 		linkarr[0].Approved = w.Approved
// 		linkarr[0].Complex = w.Complex
// 		linkarr[0].Drawnratio = w.Drawnratio
// 		linkarr[0].Designdratio = w.Designdratio
// 		linkarr[0].Checkedratio = w.Checkedratio
// 		linkarr[0].Examinedratio = w.Examinedratio
// 		linkarr[0].Datestring = w.Datestring
// 		linkarr[0].Date = w.Date
// 		linkarr[0].Created = w.Created
// 		linkarr[0].Updated = w.Updated
// 		linkarr[0].Author = w.Author
// 		links, err := models.GetCatalogLinks(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		// beego.Info(links)
// 		if len(links) > 0 {
// 			linkarray := strings.Split(links[0].Url, ",")
// 			for _, v := range linkarray {
// 				attacharr[0].Url = v
// 				// beego.Info(v.Url)
// 				Attachslice = append(Attachslice, attacharr...)
// 			}
// 		}
// 		contents, err := models.GetCatalogContents(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		for _, v := range contents {
// 			contarr[0].Content = v.Content
// 			Contentslice = append(Contentslice, contarr...)
// 		}
// 		linkarr[0].Link = Attachslice
// 		linkarr[0].Content = Contentslice
// 		// if len(Contentslice) == 0 {
// 		// 	Contentslice = make([]models.CatalogContent, 1)
// 		// 	linkarr[0].Content = Contentslice
// 		// }
// 		Attachslice = make([]models.CatalogLink, 0)
// 		Contentslice = make([]models.CatalogContent, 0)
// 		link = append(link, linkarr...)
// 	}
// 	c.Data["json"] = link //catalogs
// 	c.Data["Select2"] = select2
// 	c.Data["Starttime"] = t1
// 	c.Data["Endtime"] = t2
// 	c.Data["Ratio"] = ratios //定义的成果类型
// 	// c.Data["Secid"] = secid
// 	// c.Data["Level"] = level
// 	c.Data["UserNickname"] = user.Nickname
// 	// c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//别人提交过来，自己处于校核位置的
//校核名=登录名，状态为3，author！=登录名
// func (c *Achievement) Checked() {
// 	//1.首先判断是否注册

// 	//4.取得客户端用户名
// 	var uname string
// 	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := c.GetSession("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid") //用户id
// 	if secid == "" {                //自己登录直接显示自己
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	// secid1, err := strconv.ParseInt(secid, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	// level := c.Input().Get("level")

// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 8
// 	var t1, t2 time.Time
// 	// var convdate1, convdate2 string
// 	const lll = "2006-01-02"
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
// 		t2 = endtime.Add(+time.Duration(16) * time.Hour)
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	} else {
// 		t2 = time.Now()
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	}
// 	// convdate1 = t1.Format(lll)
// 	// convdate2 = t2.Format(lll)
// 	// usernickname := models.GetUserByUserId(secid1)
// 	ratios, err := models.GetAchievcategories()
// 	var select2 string
// 	for i, v := range ratios {
// 		if i == 0 {
// 			select2 = v.Category
// 		} else {
// 			select2 = select2 + "," + v.Category
// 		}
// 	}
// 	// beego.Info(select2)
// 	catalogs, err := models.GetcatalogChecked(secid, t1, t2)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	link := make([]CatalogLinkCont, 0)
// 	Attachslice := make([]models.CatalogLink, 0)
// 	Contentslice := make([]models.CatalogContent, 0)
// 	linkarr := make([]CatalogLinkCont, 1)
// 	attacharr := make([]models.CatalogLink, 1)
// 	contarr := make([]models.CatalogContent, 1)

// 	//这里循环，添加附件链接和设计说，校审意见
// 	for _, w := range catalogs {
// 		linkarr[0].Id = w.Id
// 		linkarr[0].ProjectNumber = w.ProjectNumber
// 		linkarr[0].ProjectName = w.ProjectName
// 		linkarr[0].DesignStage = w.DesignStage
// 		linkarr[0].Section = w.Section
// 		linkarr[0].Tnumber = w.Tnumber
// 		linkarr[0].Name = w.Name
// 		linkarr[0].Category = w.Category
// 		linkarr[0].Page = w.Page
// 		linkarr[0].Count = w.Count
// 		linkarr[0].Drawn = w.Drawn
// 		linkarr[0].Designd = w.Designd
// 		linkarr[0].Checked = w.Checked
// 		linkarr[0].Examined = w.Examined
// 		linkarr[0].Verified = w.Verified
// 		linkarr[0].Approved = w.Approved
// 		linkarr[0].Complex = w.Complex
// 		linkarr[0].Drawnratio = w.Drawnratio
// 		linkarr[0].Designdratio = w.Designdratio
// 		linkarr[0].Checkedratio = w.Checkedratio
// 		linkarr[0].Examinedratio = w.Examinedratio
// 		linkarr[0].Datestring = w.Datestring
// 		linkarr[0].Date = w.Date
// 		linkarr[0].Created = w.Created
// 		linkarr[0].Updated = w.Updated
// 		linkarr[0].Author = w.Author
// 		links, err := models.GetCatalogLinks(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		// beego.Info(links)
// 		if len(links) > 0 {
// 			linkarray := strings.Split(links[0].Url, ",")
// 			for _, v := range linkarray {
// 				attacharr[0].Url = v
// 				// beego.Info(v.Url)
// 				Attachslice = append(Attachslice, attacharr...)
// 			}
// 		}
// 		contents, err := models.GetCatalogContents(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		for _, v := range contents {
// 			contarr[0].Content = v.Content
// 			Contentslice = append(Contentslice, contarr...)
// 		}
// 		linkarr[0].Link = Attachslice
// 		linkarr[0].Content = Contentslice
// 		// if len(Contentslice) == 0 {
// 		// 	Contentslice = make([]models.CatalogContent, 1)
// 		// 	linkarr[0].Content = Contentslice
// 		// }
// 		Attachslice = make([]models.CatalogLink, 0)
// 		Contentslice = make([]models.CatalogContent, 0)
// 		link = append(link, linkarr...)
// 	}
// 	c.Data["json"] = link //catalogs
// 	c.Data["Select2"] = select2
// 	c.Data["Starttime"] = t1
// 	c.Data["Endtime"] = t2
// 	c.Data["Ratio"] = ratios //定义的成果类型
// 	// c.Data["Secid"] = secid
// 	// c.Data["Level"] = level
// 	c.Data["UserNickname"] = user.Nickname
// 	// c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//别人提交过来，自己处于审查位置的
//审查名=登录名
// func (c *Achievement) Examined() {
// 	//1.首先判断是否注册

// 	//4.取得客户端用户名
// 	var uname string
// 	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := c.GetSession("uname")
// 	if v != nil {
// 		uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//由uname取得user
// 	user, err := models.GetUserByUsername(uname)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
// 	secid := c.Input().Get("secid") //用户id
// 	if secid == "" {                //自己登录直接显示自己
// 		secid = strconv.FormatInt(user.Id, 10)
// 	}
// 	// secid1, err := strconv.ParseInt(secid, 10, 64)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	// level := c.Input().Get("level")

// 	daterange := c.Input().Get("datefilter")
// 	// beego.Info(daterange)
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	hours := 8
// 	var t1, t2 time.Time
// 	// var convdate1, convdate2 string
// 	const lll = "2006-01-02"
// 	if len(daterange) > 19 {
// 		array := strings.Split(daterange, " - ")
// 		starttime1 := array[0]
// 		endtime1 := array[1]
// 		starttime, _ := time.Parse(lll, starttime1)
// 		endtime, _ := time.Parse(lll, endtime1)
// 		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
// 		t2 = endtime.Add(+time.Duration(16) * time.Hour)
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	} else {
// 		t2 = time.Now()
// 		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
// 		// beego.Info(t1)
// 		// beego.Info(t2)
// 	}
// 	// convdate1 = t1.Format(lll)
// 	// convdate2 = t2.Format(lll)
// 	// usernickname := models.GetUserByUserId(secid1)
// 	ratios, err := models.GetAchievcategories()
// 	var select2 string
// 	for i, v := range ratios {
// 		if i == 0 {
// 			select2 = v.Category
// 		} else {
// 			select2 = select2 + "," + v.Category
// 		}
// 	}
// 	// beego.Info(select2)
// 	catalogs, err := models.GetcatalogExamined(secid, t1, t2)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	link := make([]CatalogLinkCont, 0)
// 	Attachslice := make([]models.CatalogLink, 0)
// 	Contentslice := make([]models.CatalogContent, 0)
// 	linkarr := make([]CatalogLinkCont, 1)
// 	attacharr := make([]models.CatalogLink, 1)
// 	contarr := make([]models.CatalogContent, 1)

// 	//这里循环，添加附件链接和设计说，校审意见
// 	for _, w := range catalogs {
// 		linkarr[0].Id = w.Id
// 		linkarr[0].ProjectNumber = w.ProjectNumber
// 		linkarr[0].ProjectName = w.ProjectName
// 		linkarr[0].DesignStage = w.DesignStage
// 		linkarr[0].Section = w.Section
// 		linkarr[0].Tnumber = w.Tnumber
// 		linkarr[0].Name = w.Name
// 		linkarr[0].Category = w.Category
// 		linkarr[0].Page = w.Page
// 		linkarr[0].Count = w.Count
// 		linkarr[0].Drawn = w.Drawn
// 		linkarr[0].Designd = w.Designd
// 		linkarr[0].Checked = w.Checked
// 		linkarr[0].Examined = w.Examined
// 		linkarr[0].Verified = w.Verified
// 		linkarr[0].Approved = w.Approved
// 		linkarr[0].Complex = w.Complex
// 		linkarr[0].Drawnratio = w.Drawnratio
// 		linkarr[0].Designdratio = w.Designdratio
// 		linkarr[0].Checkedratio = w.Checkedratio
// 		linkarr[0].Examinedratio = w.Examinedratio
// 		linkarr[0].Datestring = w.Datestring
// 		linkarr[0].Date = w.Date
// 		linkarr[0].Created = w.Created
// 		linkarr[0].Updated = w.Updated
// 		linkarr[0].Author = w.Author
// 		links, err := models.GetCatalogLinks(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		// beego.Info(links)
// 		if len(links) > 0 {
// 			linkarray := strings.Split(links[0].Url, ",")
// 			for _, v := range linkarray {
// 				attacharr[0].Url = v
// 				// beego.Info(v.Url)
// 				Attachslice = append(Attachslice, attacharr...)
// 			}
// 		}
// 		contents, err := models.GetCatalogContents(w.Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		for _, v := range contents {
// 			contarr[0].Content = v.Content
// 			Contentslice = append(Contentslice, contarr...)
// 		}
// 		linkarr[0].Link = Attachslice
// 		linkarr[0].Content = Contentslice
// 		// if len(Contentslice) == 0 {
// 		// 	Contentslice = make([]models.CatalogContent, 1)
// 		// 	linkarr[0].Content = Contentslice
// 		// }
// 		Attachslice = make([]models.CatalogLink, 0)
// 		Contentslice = make([]models.CatalogContent, 0)
// 		link = append(link, linkarr...)
// 	}
// 	c.Data["json"] = link //catalogs
// 	c.Data["Select2"] = select2
// 	c.Data["Starttime"] = t1
// 	c.Data["Endtime"] = t2
// 	c.Data["Ratio"] = ratios //定义的成果类型
// 	// c.Data["Secid"] = secid
// 	// c.Data["Level"] = level
// 	c.Data["UserNickname"] = user.Nickname
// 	// c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//用户登录后获得自己所在的分院和科室，然后显示对应的菜单
//同时显示所有的成果记录
//这个没用
func (c *Achievement) GetAchievementUser() {
	//读取用户id
	//查询用户分院名称和科室名称
	//查出分院id和科室id
	//从数据库取得parentid为0的单位名称和ID
	//然后
	//过滤科室下——得到价值分类名称和id
	//查询所有pid为价值分类id——得到价值名称和id，分值
	//查询所有pid为价值id——得到选择项和分值——进行字符串分割
	//构造struct——
	//这个不用：转json数据b, err := json.Marshal(group) fmt.Println(string(b))
	var user models.User
	var err error
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	//管理员可以查看
	Uid := c.Input().Get("uid")
	if Uid == "" { //如果是技术人员自己进行查看，则Uid为空
		//1.首先判断是否注册
		// if !checkAccount(c.Ctx) {
		// 	route := c.Ctx.Request.URL.String()
		// 	c.Data["Url"] = route
		// 	c.Redirect("/login?url="+route, 302)
		// 	return
		// }
		//2.取得文章的作者
		//3.由用户id取得用户名
		//4.取得客户端用户名
		// var uname string
		// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		// defer sess.SessionRelease(c.Ctx.ResponseWriter)
		// v := c.GetSession("uname")
		// if v != nil {
		// 	uname = v.(string)
		// 	c.Data["Uname"] = v.(string)
		// }
		//4.取出用户的权限等级
		// role, _ := checkRole(c.Ctx) //login里的
		//5.进行逻辑分析：
		// rolename, err := strconv.ParseInt(role, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		if roleint > 5 { //
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
		user, err = models.GetUserByUsername(username) //得到用户的id、分院和科室等
		if err != nil {
			beego.Error(err)
		}
	} else { //如果是管理员进行查看，则uid是用户名
		userid, err := strconv.ParseInt(Uid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		user = models.GetUserByUserId(userid)
	}
	c.Data["category"] = user
}

//上传excel文件，导入成果到数据库
//管理员上传，状态是根据表格，其他人上传都只能状态是的位置
func (c *Achievement) Import_Xls_Catalog() {
	// type Duration int64
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	// hours := 8
	// const lll = "2006/01/02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24

	//解析表单
	//获取上传的文件
	_, h, err := c.GetFile("catalog")
	if err != nil {
		beego.Error(err)
	}
	var path string
	if h != nil {
		//保存附件
		path = ".\\attachment\\" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("catalog", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	//2.取得客户端用户名
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := c.GetSession("uname")
	var uname string
	if v != nil {
		uname = v.(string)
	} else {
		beego.Error(err)
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var catalog m.Catalog
	const lll = "2006-01-02"
	var convdate string
	var date time.Time
	// id1 := c.Input().Get("id")
	// cid, _ := strconv.ParseInt(id1, 10, 64)
	// catalog.ParentId = cid
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	j := 0
	// var err error
	// var news string
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows { //行数,第一行从0开始
			if i != 0 { //忽略第一行标题
				// 1ProjectNumber string    //项目编号
				// 2ProjectName   string    //项目名称
				// 3DesignStage   string    //阶段
				// 4Section       string    //专业
				// 5Tnumber       string    //成果编号
				// 6Name          string    //成果名称
				// 7Drawn         string    //编制、绘制
				// 8Designd       string    //设计
				// 9Checked       string    //校核
				// 10Examined      string    //审查
				// 11Verified      string    //核定
				// 12Approved      string    //批准
				// 13Data          time.Time `orm:"null;auto_now_add;type(datetime)"`
				// 14Created       time.Time `orm:"index;auto_now_add;type(datetime)"`
				// 15Updated       time.Time `orm:"index;auto_now_add;type(datetime)"`
				// 16Author        string    //上传者
				if len(row.Cells) >= 2 { //总列数，从1开始
					catalog.ProjectNumber = row.Cells[j+1].String() //第一列从0开始,忽略第一列序号
				}
				if len(row.Cells) >= 3 {
					catalog.ProjectName = row.Cells[j+2].String()
				}
				if len(row.Cells) >= 4 {
					catalog.DesignStage = row.Cells[j+3].String()
				}
				if len(row.Cells) >= 5 {
					catalog.Section = row.Cells[j+4].String()
				}
				if len(row.Cells) >= 6 {
					catalog.Tnumber = row.Cells[j+5].String()
				}
				if len(row.Cells) >= 7 {
					catalog.Name = row.Cells[j+6].String()
				}

				if len(row.Cells) >= 8 {
					catalog.Category = row.Cells[j+7].String()
				}
				if len(row.Cells) >= 9 {
					catalog.Page = row.Cells[j+8].String()
				}
				if len(row.Cells) >= 10 {
					if row.Cells[j+9].Value != "" {
						catalog.Count, err = row.Cells[j+9].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Count = 1.0
					}
				}

				if len(row.Cells) >= 11 {
					catalog.Drawn = row.Cells[j+10].String()
				}
				if len(row.Cells) >= 12 {
					catalog.Designd = row.Cells[j+11].String()
				}
				if len(row.Cells) >= 13 {
					catalog.Checked = row.Cells[j+12].String()
				}
				if len(row.Cells) >= 14 {
					catalog.Examined = row.Cells[j+13].String()
				}
				if len(row.Cells) >= 15 {
					catalog.Verified = row.Cells[j+14].String()
				}
				if len(row.Cells) >= 16 {
					catalog.Approved = row.Cells[j+15].String()
				}

				if len(row.Cells) >= 17 {
					if row.Cells[j+16].Value != "" {
						catalog.Complex, err = row.Cells[j+16].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Complex = 1
					}
				}
				if len(row.Cells) >= 18 {
					if row.Cells[j+17].Value != "" {
						catalog.Drawnratio, err = row.Cells[j+17].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Drawnratio = 0
					}
				}
				if len(row.Cells) >= 19 {
					if row.Cells[j+18].Value != "" {
						catalog.Designdratio, err = row.Cells[j+18].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Designdratio = 0
					}
				}
				if len(row.Cells) >= 20 {
					if row.Cells[j+19].Value != "" {
						catalog.Checkedratio, err = row.Cells[j+19].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Checkedratio = 0
					}
				}
				if len(row.Cells) >= 21 {
					if row.Cells[j+20].Value != "" {
						catalog.Examinedratio, err = row.Cells[j+20].Float()
						if err != nil {
							beego.Error(err)
						}
					} else {
						catalog.Examinedratio = 0
					}
				}
				if len(row.Cells) >= 22 {
					if row.Cells[j+21].Value != "" {
						endtime2, err := row.Cells[j+21].Float()
						if err != nil {
							beego.Error(err)

						} else {
							date = xlsx.TimeFromExcelTime(endtime2, false)
						}
					} else {
						date = time.Now()
					}
					convdate = date.Format(lll)
					catalog.Datestring = convdate
					date, err = time.Parse(lll, convdate)
					if err != nil {
						beego.Error(err)
					}
					catalog.Date = date
				}

				catalog.Created = time.Now() //.Add(+time.Duration(hours) * time.Hour)
				catalog.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
				catalog.Author = uname
				if len(row.Cells) >= 23 {
					if row.Cells[j+22].Value != "" {
						catalog.State, err = row.Cells[j+22].Int()
						if err != nil {
							beego.Error(err)
						}
					} else { //如果没填，则默认为5
						catalog.State = 5
					}
					// text1, _ := row.Cells[j+22].String()
					// beego.Info(text1)
				}
				if user.Role == "1" {
					// catalog.State = 5
					_, err, _ = m.SaveCatalog(catalog)
					if err != nil {
						beego.Error(err)
					} //else {
					// 	data := news
					// 	c.Ctx.WriteString(data)
					// }
				} else { //如果不是管理员，根据自己的位置判断状态
					if catalog.Checked == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
						// beego.Info(catalog.Checked)
						catalog.State = 3
						_, err, _ = m.SaveCatalog(catalog)
						if err != nil {
							beego.Error(err)
						} //else {
						// 	data := news
						// 	c.Ctx.WriteString(data)
						// }
					} else if catalog.Designd == user.Nickname && catalog.Checked != user.Nickname && catalog.Checked != "" || catalog.Designd == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
						// beego.Info(catalog.Designd)
						catalog.State = 2
						_, err, _ = m.SaveCatalog(catalog)
						if err != nil {
							beego.Error(err)
						} //else {
						// 	data := news
						// 	c.Ctx.WriteString(data)
						// }
					} else if catalog.Drawn == user.Nickname && catalog.Designd != user.Nickname && catalog.Designd != "" || catalog.Drawn == user.Nickname && catalog.Checked != user.Nickname && catalog.Checked != "" || catalog.Drawn == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
						// beego.Info(catalog.Drawn)
						catalog.State = 1
						_, err, _ = m.SaveCatalog(catalog)
						if err != nil {
							beego.Error(err)
						} //else {
						// 	data := news
						// 	c.Ctx.WriteString(data)
						// }
					} else {
						data := "缺少下级 或 自己与下级同名 或 不能发起审查"
						c.Ctx.WriteString(data)
					}
				}
			}
		}
	}
	// c.TplName = "catalog.tpl"
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "上传成果文件")
	logs.Close()
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok"
		c.Ctx.WriteString(data)
	}
	//上传后应该返回ok字样即可
	// c.Redirect("/admin", 302)
}

//在线添加目录，即插入一条目录
//只能填写自己是设计/绘图/编制/校核的成果,并且至少下级有一级，不能添加审查
func (c *Achievement) AddCatalog() {
	//4.取得客户端用户名
	var uname string
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := c.GetSession("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var catalog m.Catalog
	catalog.ProjectNumber = c.Input().Get("Pnumber")
	catalog.ProjectName = c.Input().Get("Pname")
	catalog.DesignStage = c.Input().Get("Stage")
	catalog.Section = c.Input().Get("Section")
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Category = c.Input().Get("Category")
	catalog.Page = c.Input().Get("Page")

	count1 := c.Input().Get("Count")
	if count1 != "" {
		catalog.Count, err = strconv.ParseFloat(count1, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	drawninput := c.Input().Get("Drawn")
	if drawninput != "" {
		//由uname取得user
		userdrawn, err := models.GetUserByUsername(drawninput)
		if err != nil {
			beego.Error(err)
		}
		catalog.Drawn = userdrawn.Nickname
	}

	designdinput := c.Input().Get("Designd")
	if designdinput != "" {
		//由uname取得user
		userdesignd, err := models.GetUserByUsername(designdinput)
		if err != nil {
			beego.Error(err)
		}
		catalog.Designd = userdesignd.Nickname
	}

	checkedinput := c.Input().Get("Checked")
	if checkedinput != "" {
		//由uname取得user
		userchecked, err := models.GetUserByUsername(checkedinput)
		if err != nil {
			beego.Error(err)
		}
		catalog.Checked = userchecked.Nickname
	}

	examinedinput := c.Input().Get("Examined")
	if examinedinput != "" {
		//由uname取得user
		userexamined, err := models.GetUserByUsername(examinedinput)
		if err != nil {
			beego.Error(err)
		}
		catalog.Examined = userexamined.Nickname
	}
	// catalog.Verified = c.Input().Get("Verified")
	// catalog.Approved = c.Input().Get("Approved")

	// complex, err := strconv.ParseFloat(c.Input().Get("Complex"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Complex = complex
	drawnratio1 := c.Input().Get("Drawnratio")
	if drawnratio1 != "" {
		catalog.Drawnratio, err = strconv.ParseFloat(drawnratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	designdratio1 := c.Input().Get("Designdratio")
	if designdratio1 != "" {
		catalog.Designdratio, err = strconv.ParseFloat(designdratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	// checkedratio, err := strconv.ParseFloat(c.Input().Get("Checkedratio"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Checkedratio = checkedratio
	// examinedratio, err := strconv.ParseFloat(c.Input().Get("Examinedratio"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Examinedratio = examinedratio

	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	// hours := 8
	inputdate := c.Input().Get("Date")
	// beego.Info(inputdate)
	var t1 time.Time
	// var convdate1, convdate2 string
	const lll = "2006-01-02"
	if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
		t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
		if err != nil {
			beego.Error(err)
		}
		convdate := t1.Format(lll)
		catalog.Datestring = convdate
		catalog.Date = t1
		// t1 = printtime.Add(+time.Duration(hours) * time.Hour)
	} else { //如果取系统时间，则需要加8小时
		date := time.Now()
		convdate := date.Format(lll)
		catalog.Datestring = convdate
		date, err = time.Parse(lll, convdate)
		if err != nil {
			beego.Error(err)
		}
		catalog.Date = date
	}

	catalog.Created = time.Now() //.Add(+time.Duration(hours) * time.Hour)
	catalog.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)

	var news string
	var id int64
	catalog.Author = uname
	catalog.Complex = 1
	if catalog.Checked == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
		// beego.Info(catalog.Checked)
		catalog.State = 3
		id, err, news = m.SaveCatalog(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := c.Input().Get("Link") //附件链接地址
			if link1 != "" {
				array := strings.Split(link1, ",")
				for _, v := range array {
					_, err = models.AddCatalogLink(id, v)
					if err != nil {
						beego.Error(err)
					}
				}
			}
			content1 := c.Input().Get("Content") //设计说明
			if content1 != "" {
				_, err = models.AddCatalogContent(id, content1, 1)
				if err != nil {
					beego.Error(err)
				}
			}
			data := news
			c.Ctx.WriteString(data)
		}
	} else if catalog.Designd == user.Nickname && catalog.Checked != user.Nickname && catalog.Checked != "" || catalog.Designd == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
		// beego.Info(catalog.Designd)
		catalog.State = 2
		id, err, news = m.SaveCatalog(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := c.Input().Get("Link") //附件链接地址
			if link1 != "" {
				array := strings.Split(link1, ",")
				for _, v := range array {
					_, err = models.AddCatalogLink(id, v)
					if err != nil {
						beego.Error(err)
					}
				}
			}
			content1 := c.Input().Get("Content") //设计说明
			if content1 != "" {
				_, err = models.AddCatalogContent(id, content1, 1)
				if err != nil {
					beego.Error(err)
				}
			}
			data := news
			c.Ctx.WriteString(data)
		}
	} else if catalog.Drawn == user.Nickname && catalog.Designd != user.Nickname && catalog.Designd != "" || catalog.Drawn == user.Nickname && catalog.Checked != user.Nickname && catalog.Checked != "" || catalog.Drawn == user.Nickname && catalog.Examined != user.Nickname && catalog.Examined != "" {
		// beego.Info(catalog.Drawn)
		catalog.State = 1
		id, err, news = m.SaveCatalog(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := c.Input().Get("Link") //附件链接地址
			if link1 != "" {
				array := strings.Split(link1, ",")
				for _, v := range array {
					_, err = models.AddCatalogLink(id, v)
					if err != nil {
						beego.Error(err)
					}
				}
			}
			content1 := c.Input().Get("Content") //设计说明
			if content1 != "" {
				_, err = models.AddCatalogContent(id, content1, 1)
				if err != nil {
					beego.Error(err)
				}
			}
			data := news
			c.Ctx.WriteString(data)
		}
	} else {
		data := "缺少下级 或 自己与下级同名 或 不能发起审查"
		c.Ctx.WriteString(data)
	}
	//****保存附件链接地址和设计说明
	// if err == nil {
	// 	link1 := c.Input().Get("Link") //附件链接地址
	// 	if link1 != "" {
	// 		_, err = models.AddCatalogLink(id, link1)
	// 		if err != nil {
	// 			beego.Error(err)
	// 		}
	// 	}
	// 	content1 := c.Input().Get("Content") //设计说明
	// 	if content1 != "" {
	// 		_, err = models.AddCatalogContent(id, content1, 1)
	// 		if err != nil {
	// 			beego.Error(err)
	// 		}
	// 	}
	// }
	//只能添加自己是设计者或绘图者的成果
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "添加记录")
	logs.Close()
	// err := models.ModifyCatalog(tid, title, tnumber)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Redirect("/catalog", 302)
}

//在线修改保存某个字段
func (c *Achievement) ModifyCatalog() {
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")

	ids := c.GetString("ids")
	if ids != "" { //修改选中。问题，选中的是其他几个，修改当前这个没有选中，则不修改，会不会很奇怪？
		array := strings.Split(ids, ",")
		for _, v := range array {
			// pid = strconv.FormatInt(v1, 10)
			if v != pk { //避免与下面的id重复
				//id转成64位
				idNum, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				err = m.ModifyCatalog(idNum, name, value)
				if err != nil {
					beego.Error(err)
				} else {
					// data := value
					// c.Ctx.WriteString(data)
				}
			}
		}
	}

	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//无论如何都修改当前，重复修改了
	err = m.ModifyCatalog(id, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := value
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + pk)
	logs.Close()
}

//提交记录，判断下一个人位置，状态对应改为对应值
func (c *Achievement) SendCatalog() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//取得用户名
	// if rolename > 2 && uname != username {
	var ob []CatalogLinkCont
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	for _, v := range ob {
		//查出cidnum这个有没有审查人员，没有就直接state+2
		catalog, err := m.GetCatalog(v.Id)
		if err != nil {
			beego.Error(err)
		}
		switch catalog.State {
		case 1:
			if catalog.Designd != "" {
				err = m.ModifyCatalogState(v.Id, 2)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Checked != "" {
				err = m.ModifyCatalogState(v.Id, 3)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Examined != "" {
				err = m.ModifyCatalogState(v.Id, 4)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			}
		case 2:
			if catalog.Checked != "" {
				err = m.ModifyCatalogState(v.Id, 3)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Examined != "" {
				err = m.ModifyCatalogState(v.Id, 4)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			} else {
				err = m.ModifyCatalogState(v.Id, 5)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交汇总ok!"
					c.Ctx.WriteString(data)
				}
			}
		case 3:
			if catalog.Examined != "" {
				err = m.ModifyCatalogState(v.Id, 4)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交下一级ok!"
					c.Ctx.WriteString(data)
				}
			} else {
				err = m.ModifyCatalogState(v.Id, 5)
				if err != nil {
					beego.Error(err)
				} else {
					data := "提交汇总ok!"
					c.Ctx.WriteString(data)
				}
			}
		case 4:
			err = m.ModifyCatalogState(v.Id, 5)
			if err != nil {
				beego.Error(err)
			} else {
				data := "提交汇总ok!"
				c.Ctx.WriteString(data)
			}
		}
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "提交记录" + strconv.FormatInt(v.Id, 10))
		logs.Close()
	}

}

//退回一条目录，状态降到前一个人的位置
func (c *Achievement) DownSendCatalog() {
	var ob []CatalogLinkCont
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	for _, v := range ob {
		catalog, err := m.GetCatalog(v.Id)
		if err != nil {
			beego.Error(err)
		}
		// cid := c.Input().Get("CatalogId")
		// id, err := strconv.ParseInt(cid, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// catalog, err := m.GetCatalog(id)
		// if err != nil {
		// 	beego.Error(err)
		// }
		switch catalog.State {
		case 4:
			if catalog.Checked != "" {
				err = m.ModifyCatalogState(v.Id, 3)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Designd != "" {
				err = m.ModifyCatalogState(v.Id, 2)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Drawn != "" {
				err = m.ModifyCatalogState(v.Id, 1)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			}
		case 3:
			if catalog.Designd != "" {
				err = m.ModifyCatalogState(v.Id, 2)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			} else if catalog.Drawn != "" {
				err = m.ModifyCatalogState(v.Id, 1)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			}
		case 2:
			if catalog.Drawn != "" {
				err = m.ModifyCatalogState(v.Id, 1)
				if err != nil {
					beego.Error(err)
				} else {
					data := "退回ok!"
					c.Ctx.WriteString(data)
				}
			}
		}
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "退回记录" + strconv.FormatInt(v.Id, 10))
		logs.Close()
	}
}

//删除一条目录
func (c *Achievement) DeleteCatalog() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//取得用户名

	// if rolename > 2 && uname != username {
	// cid := c.Input().Get("CatalogId")
	var ob []CatalogLinkCont
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	for _, v := range ob {
		// catalog, err := m.GetCatalog(v.Id)
		// if err != nil {
		// 	beego.Error(err)
		// }
		err := m.DeletCatalog(v.Id)
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
			logs := logs.NewLogger(1000)
			logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
			logs.EnableFuncCallDepth(true)
			logs.Info(c.Ctx.Input.IP() + " " + "删除记录" + strconv.FormatInt(v.Id, 10))
			logs.Close()
		}
	}
}

//*********下面是统计分析************
//用户参与的项目列表和自己的贡献及项目总分值
func (c *Achievement) Participate() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//4.取得客户端用户名
	var uname string
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := c.GetSession("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	//由uname取得user
	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	secid := c.Input().Get("secid")
	if secid == "" { //如果为空，则用登录的
		secid = strconv.FormatInt(user.Id, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//现在的月份
	const MMM = "2006-01"
	// date := time.Now()
	//	fmt.Println(date)
	month1 := time.Now().Format(MMM)
	// fmt.Println(convdate)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}
	catalogs, err := models.Getparticipatebyuserid(secid, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(&catalogs)
	slice1 := make([]Project, 0)
	for _, v := range catalogs {
		value1, value2, err := models.Getprojuserspecialty(secid1, v.ProjectNumber, v.DesignStage, v.Section, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		aa := make([]Project, 1)
		aa[0].Id = v.Id //没这个Id，将导致点击表格详细，无法传递项目id号
		aa[0].ProjectNumber = v.ProjectNumber
		aa[0].ProjectName = v.ProjectName
		aa[0].DesignStage = v.DesignStage
		aa[0].Section = v.Section
		aa[0].Value = value1                           //项目（阶段、专业）总分值
		aa[0].Myvalue = value2                         //用户参与的分值
		aa[0].Percent = models.Round(value2/value1, 1) //用户分值占比
		slice1 = append(slice1, aa...)
	}

	c.Data["json"] = slice1
	c.ServeJSON()
}

//点击个人参与的项目，弹出模态框，显示这个项目所有成果
//没有区分时间
//没有按时间排序！！！
func (c *Achievement) ProjectAchievement() {
	cid := c.Input().Get("CatalogId")
	// var id string
	// var cidNum int64
	var err error
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//查出项目特性:
	catalog, err := m.GetCatalog(cidNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目编号和项目阶段、专业，查出所有成果
	achievements, err := m.GetProjectAchievement(catalog.ProjectNumber, catalog.DesignStage, catalog.Section)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(achievements)
	c.Data["json"] = achievements
	c.ServeJSON()
}

//显示某个专业内成本分布情况
func (c *Achievement) Specialty() {
	ratios, err := models.GetAchievcategories()
	if err != nil {
		beego.Error(err)
	}
	var select2 []string
	for _, v := range ratios {
		aa := make([]string, 1)
		aa[0] = v.Category
		select2 = append(select2, aa...)
	}

	c.Data["Select2"] = select2
	c.Data["Ratio"] = ratios //定义的成果类型
	// c.Data["specialty"] = slice1
	c.TplName = "merit/specialty_show.tpl"
}

//显示自己一个月内成果类型情况
func (c *Achievement) Echarts() {
	//1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	return
	// }
	// //4.取得客户端用户名
	// var uname string
	// // sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// // defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// }
	// //由uname取得user
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// //4.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	if roleint > 4 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	if secid == "" { //如果为空，则用登录的
		secid = strconv.FormatInt(uid, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)

	var t1, t2 time.Time
	t2 = time.Now()
	t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
	slice1 := make([]m.Specialty, 0)
	ratios, err := models.GetAchievcategories()
	if err != nil {
		beego.Error(err)
	}
	for _, v := range ratios {
		value1, err := models.Getuserspecialty(secid1, v.Category, t1, t2)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]m.Specialty, 1)
		aa[0].Value = m.Round(value1*v.Rationum, 1)
		aa[0].Name = v.Category
		slice1 = append(slice1, aa...)
	}
	c.Data["json"] = slice1
	c.ServeJSON()
}

//显示自己一年来，成果类型情况
func (c *Achievement) Echarts2() {
	//1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	return
	// }
	// //4.取得客户端用户名
	// var uname string
	// // sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// // defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// }
	// //由uname取得user
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// //4.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	if roleint > 4 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	if secid == "" { //如果为空，则用登录的
		secid = strconv.FormatInt(uid, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//现在的月份
	const MMM = "2006-01"
	month1 := time.Now().Format(MMM)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}

	slice1 := make([]m.Specialty, 0)
	ratios, err := models.GetAchievcategories()
	if err != nil {
		beego.Error(err)
	}
	for _, v := range ratios {
		value1, err := models.Getuserspecialty(secid1, v.Category, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		aa := make([]m.Specialty, 1)
		aa[0].Value = m.Round(value1*v.Rationum, 1)
		aa[0].Name = v.Category
		slice1 = append(slice1, aa...)
	}
	c.Data["json"] = slice1
	c.ServeJSON()
}

//显示某个项目（阶段和专业）一年来，成果类型情况
func (c *Achievement) Echarts3() {
	cid := c.Input().Get("CatalogId")
	tidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//查出项目特性:
	catalog, err := m.GetCatalog(tidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取项目编号
	//获取阶段
	//获取专业
	//根据项目编号阶段和专业名称，查出图纸、报告、计算书等各类型分值
	//现在的月份
	const MMM = "2006-01"
	month1 := time.Now().Format(MMM)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}
	slice1 := make([]m.Specialty, 0)
	ratios, err := models.GetAchievcategories()
	if err != nil {
		beego.Error(err)
	}
	for _, v := range ratios {
		value1, err := models.Getspecialty(catalog.ProjectNumber, catalog.DesignStage, catalog.Section, v.Category, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		aa := make([]m.Specialty, 1)
		aa[0].Value = m.Round(value1*v.Rationum, 1)
		aa[0].Name = v.Category
		slice1 = append(slice1, aa...)
	}
	//去除空的ratio
	c.Data["json"] = slice1
	c.ServeJSON()
}

//展示一个项目所有参与人员的贡献率
func (c *Achievement) ProjectUserParticipate() {
	cid := c.Input().Get("CatalogId")
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//查出项目特性:
	catalog, err := m.GetCatalog(cidNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目编号和项目阶段、专业，查出所有成果
	achievements, err := m.GetProjectAchievement(catalog.ProjectNumber, catalog.DesignStage, catalog.Section)
	if err != nil {
		beego.Error(err)
	}
	//循环成果，得到制图设计校核审查人员名单
	//利用map的index不允许重复特性，进行名单去重
	var name map[string]int
	name = make(map[string]int)
	for i, v := range achievements {
		if v.Drawn != "" {
			name[v.Drawn] = i
		}
		if v.Designd != "" {
			name[v.Designd] = i
		}
		if v.Checked != "" {
			name[v.Checked] = i
		}
		if v.Examined != "" {
			name[v.Examined] = i
		}
	}
	//现在的月份
	const MMM = "2006-01"
	month1 := time.Now().Format(MMM)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}
	//循环名单，根据项目阶段专业，查出分值
	slice1 := make([]m.Specialty, 0)
	for k, _ := range name {
		value, err := models.Getprojuserspecialty1(k, catalog.ProjectNumber, catalog.DesignStage, catalog.Section, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		aa := make([]m.Specialty, 1)
		aa[0].Value = value
		aa[0].Name = k
		slice1 = append(slice1, aa...)
	}
	c.Data["json"] = slice1
	c.ServeJSON()
}

//*******下面是科室统计分析
//科室参与的项目列表和项目总分值
func (c *Achievement) SecParticipate() {
	//由科室id取得所有用户id
	secid := c.Input().Get("secid")
	// if secid == "" { //如果为空，则用登录的
	// 	secid = strconv.FormatInt(user.Id, 10)
	// }
	// secid1, err := strconv.ParseInt(secid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//根据科室id查所有员工
	users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
	// beego.Info(users)
	if err != nil {
		beego.Error(err)
	}
	//现在的月份
	const MMM = "2006-01"
	// date := time.Now()
	//	fmt.Println(date)
	month1 := time.Now().Format(MMM)
	// fmt.Println(convdate)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}

	catalogs := make([]*models.Catalog, 0)
	for _, w := range users { //循环用户id
		//由用户Id取得所有成果_得到参与的项目和阶段——去重
		userid := strconv.FormatInt(w.Id, 10)
		aa, err := models.Getparticipatebyuserid(userid, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		catalogs = append(catalogs, aa...)
	}
	//对成果进行去重
	bb := models.Rm_duplicate(catalogs)

	slice1 := make([]Project, 0)
	for _, v := range bb {
		//查出一个项目(阶段、专业)时间段内总分值和某个用户的总分值
		value1, _, err := models.Getprojuserspecialty(v.Id, v.ProjectNumber, v.DesignStage, v.Section, month2.AddDate(0, -11, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		cc := make([]Project, 1)
		cc[0].Id = v.Id //没这个Id，将导致点击表格详细，无法传递项目id号
		cc[0].ProjectNumber = v.ProjectNumber
		cc[0].ProjectName = v.ProjectName
		cc[0].DesignStage = v.DesignStage
		cc[0].Section = v.Section
		cc[0].Value = value1 //项目（阶段、专业）总分值
		slice1 = append(slice1, cc...)
	}

	c.Data["json"] = slice1
	c.ServeJSON()
}

//显示科室月度全部成果
func (c *Achievement) SecProjectAchievement() {
	secid := c.Input().Get("secid")
	//根据科室id查所有员工
	users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
	// beego.Info(users)
	if err != nil {
		beego.Error(err)
	}
	//现在的月份
	const MMM = "2006-01"
	// date := time.Now()
	//	fmt.Println(date)
	month1 := time.Now().Format(MMM)
	// fmt.Println(convdate)
	month2, err := time.Parse(MMM, month1)
	if err != nil {
		beego.Error(err)
	}

	catalogs := make([]*models.Catalog, 0)
	for _, w := range users { //循环用户id
		//由用户Id取得所有成果_得到参与的项目和阶段——去重
		userid := strconv.FormatInt(w.Id, 10)
		aa, err := models.Getparticipatebyuserid(userid, month2.AddDate(0, -1, 0), month2.AddDate(0, +1, 0))
		if err != nil {
			beego.Error(err)
		}
		catalogs = append(catalogs, aa...)
	}
	//对成果进行去重
	bb := models.Rm_duplicate(catalogs)

	slice1 := make([]*m.Catalog, 0)
	for _, v := range bb {
		//查出项目特性:
		catalog, err := m.GetCatalog(v.Id)
		if err != nil {
			beego.Error(err)
		}
		//根据项目编号和项目阶段、专业，查出所有成果
		achievements, err := m.GetProjectAchievement(catalog.ProjectNumber, catalog.DesignStage, catalog.Section)
		if err != nil {
			beego.Error(err)
		}
		slice1 = append(slice1, achievements...)
	}
	// beego.Info(achievements)
	c.Data["json"] = slice1
	c.ServeJSON()
}

//参考时间的转换：由string转成time data
// func (c *TaskController) AddTask() {
// 	//解析表单     表示时间的变量和字段，应为time.Time类型
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	//seconds := 10
// 	//fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
// 	hours := 8
// 	//	time.Duration(hours) * time.Hour
// 	// t1 := t.Add(time.Duration(hours) * time.Hour)
// 	// datestring = t1.Format(layout) //{{dateformat .Created "2006-01-02 15:04:05"}}
// 	// return

// 	var err error
// 	tid := c.Input().Get("tid")
// 	title := c.Input().Get("title")
// 	content := c.Input().Get("content")
// 	daterange := c.Input().Get("datefilter")
// 	array := strings.Split(daterange, " - ")
// 	// for _, v := range array {
// 	starttime1 := array[0]
// 	beego.Info(array[0])
// 	endtime1 := array[1]
// 	beego.Info(array[1])
// 	// }
// 	// starttime1 := c.Input().Get("starttime")
// 	// endtime1 := c.Input().Get("endtime")
// 	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
// 	starttime, _ := time.Parse(lll, starttime1)
// 	endtime, _ := time.Parse(lll, endtime1)
// 	t1 := starttime.Add(-time.Duration(hours) * time.Hour)
// 	t2 := endtime.Add(-time.Duration(hours) * time.Hour)
// 	//12-19-2015 22:40:24
// 	// ck, err := c.Ctx.Request.Cookie("uname")
// 	// if err != nil {
// 	// beego.Error(err)
// 	// }
// 	// uname := ck.Value
// 	if len(tid) == 0 {
// 		err = models.AddTask(title, content, t1, t2)
// 		// beego.Info(attachment)
// 		// } else {
// 		// 	err = models.UpdateTask(tid, title, content)
// 	}
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/todo", 302)
// }
