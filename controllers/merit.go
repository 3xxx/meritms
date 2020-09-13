// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
//在线价值登记
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"github.com/3xxx/meritms/models"
	// "sort"
	"strconv"
	"strings"
	"time"
)

type MeritController struct {
	beego.Controller
}

type MeritMark struct {
	Choose string `json:"choose"`
	Mark1  string `json:"mark1"` //打分1
}

type List1 struct {
	Choose string `json:"choose"`
	Mark1  string `json:"mark1"` //打分1
}

//价值内容
type MeritList struct {
	Id         int64  `json:"Id"`
	Pid        int64  `form:"-"`
	Title      string `json:"text"`
	Tags       [2]int `json:"tags"`
	Mark       string `json:"mark2"` //打分2
	List       string //大型、中型……
	ListMark   string //对应列表打分
	Level      string `json:"Level"` //4
	Selectable bool   `json:"selectable"`
}

//价值分类
type MeritCategory struct { //项目负责人——链接——大、中、小
	Id         int64       `json:"Id"`
	Pid        int64       `form:"-"` //这个为0表示是分类
	Title      string      `json:"text"`
	Tags       [2]int      `json:"tags"`
	List       []MeritList `json:"nodes"`
	Level      string      `json:"Level"` //4
	Selectable bool        `json:"selectable"`
	// Parent2 string
	// Href    string `json:"href"`
}

type MeritSecoffice struct { //专业室：水工、施工……
	Id         int64           `json:"Id"`
	Pid        int64           `form:"-"`
	Title      string          `json:"text"`
	List       []MeritCategory `json:"nodes"`
	Level      string          `json:"Level"` //2
	Selectable bool            `json:"selectable"`
}

type MeritDepartment struct { //分院：施工预算、水工分院……
	Id         int64            `json:"Id"`
	Pid        int64            `form:"-"`
	Title      string           `json:"text"` //这个后面json仅仅对于encode解析有用
	Selectable bool             `json:"selectable"`
	Secoffice  []MeritSecoffice `json:"nodes"`
	Level      string           `json:"Level"` //1
}

type Person struct {
	// Id int64 `json:"Id"`
	// Name       string `json:"Name"`//用户名
	// Department string `json:"Department"`
	// Secoffice  string `json:"Keshi"` //当controller返回json给view的时候，必须用text作为字段
	// Numbers    int    //记录个数
	Marks int //分值
	// UserId     int64  //用户id
	User models.User
}

type MeritTopicSlice struct {
	Id           int64
	MeritCate    string `json:"meritcate"` //价值分类
	Merit        string `json:"merit"`     //价值
	MeritId      int64
	UserId       int64
	UserNickName string `json:"usernickname"`
	Title        string `json:"title"`
	Choose       string
	Content      string    `json:"content"`
	State        int       //1编写状态，未提交；2编写者提交，等待审核确认;3,已经审核确认
	Mark         int       `json:"mark"`
	Created      time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//struct排序
// type person1 []Person
type person1 []models.UserMeritTopics

func (list person1) Len() int {
	return len(list)
}

func (list person1) Less(i, j int) bool {
	if list[i].Marks > list[j].Marks {
		return true
	} else if list[i].Marks < list[j].Marks {
		return false
	} else {
		return list[i].User.Nickname > list[j].User.Nickname
	}
}

func (list person1) Swap(i, j int) {
	// var temp Person = list[i]
	var temp models.UserMeritTopics = list[i]
	list[i] = list[j]
	list[j] = temp
}

//管理员进行人员价值排序查看
//排序第一排序为部门，第二排序为科室，第三排序为分值
// func (c *MeritController) GetPerson() {
// 	//1.首先判断是否注册
// 	if !checkAccount(c.Ctx) {
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/login?url="+route, 302)
// 		return
// 	}
// 	//2.取得客户端用户名
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		c.Data["Uname"] = v.(string)
// 	}
// 	//3.取出用户的权限等级
// 	role, _ := checkRole(c.Ctx) //login里的
// 	//4.进行逻辑分析
// 	if role > 2 { //
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/roleerr?url="+route, 302)
// 		return
// 	}

// 	var numbers1, marks1 int
// 	slice1 := make([]Person, 0)
// 	users, _ := models.GetAllusers(1, 2000, "Id")
// 	for i1, _ := range users {
// 		//根据价值id和用户id，得到成果，统计数量和分值
// 		//取得用户的价值数量和分值
// 		_, numbers, marks, err := models.GetMerit(0, users[i1].Id)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		marks1 = marks1 + marks
// 		numbers1 = numbers1 + numbers
// 		aa := make([]Person, 1)
// 		aa[0].Id = users[i1].Id //这里用for i1,v1,然后用v1.Id一样的意思
// 		aa[0].Name = users[i1].Nickname
// 		aa[0].Department = users[i1].Department
// 		aa[0].Keshi = users[i1].Secoffice
// 		aa[0].Numbers = numbers1
// 		aa[0].Marks = marks1
// 		slice1 = append(slice1, aa...)
// 		marks1 = 0
// 		numbers1 = 0
// 	}
// 	c.Data["person"] = slice1
// 	c.TplName = "admin_person.tpl"
// }

//管理员1登录后显示侧栏：所有部门——科室——价值分类——价值
//分院2登录侧栏显示：       分院——科室——价值分类——价值
//科室主任3登录侧栏显示：         科室——价值分类——价值
//用户4登录后侧栏显示：                 价值分类——价值
func (c *MeritController) GetMerit() {
	c.Data["IsMerit"] = true
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
	// beego.Info(role)
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
	//4.取出用户的权限等级
	if roleint > 4 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	meritlist := make([]MeritList, 0)
	meritcategory := make([]MeritCategory, 0)
	meritsecoffice := make([]MeritSecoffice, 0)
	meritdepartment := make([]MeritDepartment, 0)
	//由uname取得user,获得user的分院名称
	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}
	switch role {
	case "1": //管理员登录显示的侧栏：所有部门——科室——价值分类——价值
		depart, err := models.GetAdminDepart(0) //得到多个分院
		if err != nil {
			beego.Error(err)
		}
		for i1, v1 := range depart {
			aa := make([]MeritDepartment, 1)
			aa[0].Id = depart[i1].Id
			aa[0].Level = "1"
			// aa[0].Pid = category[0].Id
			aa[0].Title = depart[i1].Title //分院名称
			// beego.Info(category1[i1].Title)
			secoffice, err := models.GetAdminDepart(depart[i1].Id) //得到多个科室
			if err != nil {
				beego.Error(err)
			}
			//如果返回科室为空，则直接取得价值分类
			//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
			//应该是反过来找出所有没有科室字段的人员，把他放在部门下
			if len(secoffice) > 0 {
				for i2, v2 := range secoffice {
					bb := make([]MeritSecoffice, 1)
					bb[0].Id = v2.Id
					bb[0].Level = "2"
					bb[0].Pid = v1.Id
					bb[0].Title = v2.Title //科室名称
					// beego.Info(category2[i2].Title)
					//根据分院和科室查所有价值分类
					meritcates, err := models.GetMeritsbySec(depart[i1].Title, secoffice[i2].Title) //得到员工姓名
					if err != nil {
						beego.Error(err)
					}
					for i3, v3 := range meritcates {
						cc := make([]MeritCategory, 1)
						cc[0].Id = meritcates[i3].Id
						cc[0].Level = "3"
						cc[0].Pid = secoffice[i2].Id
						merittitle, err := models.GetAdminMeritbyId(v3.MeritId) //因为这个数据库只是科室和分类的对应表
						if err != nil {
							beego.Error(err)
						}
						cc[0].Title = merittitle.Title //名称
						//由价值分类取得所有价值列表
						merits, err := models.GetAdminMerit(v3.MeritId)
						if err != nil {
							beego.Error(err)
						}
						//循环价值分类，取得价值列表
						for _, v4 := range merits {
							ee := make([]MeritList, 1)
							ee[0].Id = v4.Id
							ee[0].Level = "4"
							ee[0].Pid = v3.Id

							ee[0].Title = v4.Title //名称

							// beego.Info(users[i3].Nickname)
							ee[0].Selectable = true
							//根据userid和meritid，取得项数和分值

							// ee[0].Tags[0]=//有几项
							// ee[0].Tags[1]=//总共多少分

							meritlist = append(meritlist, ee...)
						}
						cc[0].List = meritlist
						// cc[0].Selectable = false

						// cc[0].Tags[0]=
						// cc[0].Tags[1]=

						meritlist = make([]MeritList, 0)
						meritcategory = append(meritcategory, cc...)
					}
					// bb[0].Tags[0] = strconv.Itoa(count)
					bb[0].List = meritcategory
					bb[0].Selectable = true
					meritcategory = make([]MeritCategory, 0) //再把slice置0
					meritsecoffice = append(meritsecoffice, bb...)
					// depcount = depcount + count //部门人员数=科室人员数相加
				}
				// aa[0].Secoffice = achsecoffice
				// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				// achdepart = append(achdepart, aa...)
			}
			//查出所有有这个部门但科室名为空的价值分类
			//根据分院查所有员工
			// beego.Info(category1[i1].Title)
			meritcates, err := models.GetMeritsbySecOnly(depart[i1].Title) //得到价值分类
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(meritcates)
			for _, v5 := range meritcates {
				dd := make([]MeritSecoffice, 1) //把科室当作价值分类
				dd[0].Id = v5.Id
				dd[0].Level = "3"
				// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
				dd[0].Pid = v1.Id
				merittitle, err := models.GetAdminMeritbyId(v5.MeritId) //因为这个数据库只是科室和分类的对应表
				if err != nil {
					beego.Error(err)
				}
				dd[0].Title = merittitle.Title //名称
				//由价值分类取得所有价值列表
				merits, err := models.GetAdminMerit(v5.MeritId)
				if err != nil {
					beego.Error(err)
				}
				//循环价值分类，取得价值列表
				for _, v6 := range merits {
					ee := make([]MeritCategory, 1)
					ee[0].Id = v6.Id
					ee[0].Level = "4"
					ee[0].Pid = v5.Id

					ee[0].Title = v6.Title //名称

					// beego.Info(users[i3].Nickname)
					ee[0].Selectable = true
					meritcategory = append(meritcategory, ee...)
				}
				dd[0].List = meritcategory
				meritcategory = make([]MeritCategory, 0) //再把slice置0
				// dd[0].Title = strconv.FormatInt(meritcates[i3].Id, 10) //名称——关键，把价值分类当作科室名
				dd[0].Selectable = false
				meritsecoffice = append(meritsecoffice, dd...)
			}
			// aa[0].Tags[0] = count + depcount
			// count = 0
			// depcount = 0
			aa[0].Secoffice = meritsecoffice
			aa[0].Selectable = false                   //默认是false_点击展开
			meritsecoffice = make([]MeritSecoffice, 0) //再把slice置0
			meritdepartment = append(meritdepartment, aa...)
		}
		c.Data["json"] = meritdepartment
	case "2": //分院管理人员登录
		depart, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}

		aa := make([]MeritDepartment, 1)
		aa[0].Id = depart.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = depart.Title //分院名称
		// beego.Info(category1[i1].Title)
		secoffice, err := models.GetAdminDepart(depart.Id) //得到多个科室
		if err != nil {
			beego.Error(err)
		}
		//如果返回科室为空，则直接取得价值分类
		//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
		//应该是反过来找出所有没有科室字段的人员，把他放在部门下
		if len(secoffice) > 0 {
			for _, v2 := range secoffice {
				bb := make([]MeritSecoffice, 1)
				bb[0].Id = v2.Id
				bb[0].Level = "2"
				bb[0].Pid = depart.Id
				bb[0].Title = v2.Title //科室名称
				// beego.Info(category2[i2].Title)
				//根据分院和科室查所有价值分类
				meritcates, err := models.GetMeritsbySec(depart.Title, v2.Title) //得到员工姓名
				if err != nil {
					beego.Error(err)
				}
				for _, v3 := range meritcates {
					cc := make([]MeritCategory, 1)
					cc[0].Id = v3.Id
					cc[0].Level = "3"
					cc[0].Pid = v2.Id
					merittitle, err := models.GetAdminMeritbyId(v3.MeritId) //因为这个数据库只是科室和分类的对应表
					if err != nil {
						beego.Error(err)
					}
					cc[0].Title = merittitle.Title //名称
					//由价值分类取得所有价值列表
					merits, err := models.GetAdminMerit(v3.MeritId)
					if err != nil {
						beego.Error(err)
					}
					//循环价值分类，取得价值列表
					for _, v4 := range merits {
						ee := make([]MeritList, 1)
						ee[0].Id = v4.Id
						ee[0].Level = "4"
						ee[0].Pid = v3.Id

						ee[0].Title = v4.Title //名称

						// beego.Info(users[i3].Nickname)
						ee[0].Selectable = true
						meritlist = append(meritlist, ee...)
					}
					cc[0].List = meritlist
					// cc[0].Selectable = false
					meritlist = make([]MeritList, 0)
					meritcategory = append(meritcategory, cc...)
				}
				// bb[0].Tags[0] = strconv.Itoa(count)
				bb[0].List = meritcategory
				bb[0].Selectable = true
				meritcategory = make([]MeritCategory, 0) //再把slice置0
				meritsecoffice = append(meritsecoffice, bb...)
				// depcount = depcount + count //部门人员数=科室人员数相加
			}
			// aa[0].Secoffice = achsecoffice
			// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			// achdepart = append(achdepart, aa...)
			aa[0].Selectable = false
		} else {
			//查出所有有这个部门但科室名为空的价值分类
			//根据分院查所有员工
			// beego.Info(category1[i1].Title)
			meritcates, err := models.GetMeritsbySecOnly(depart.Title) //得到价值分类
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(meritcates)
			for _, v5 := range meritcates {
				dd := make([]MeritSecoffice, 1) //把科室当作价值分类
				dd[0].Id = v5.Id
				dd[0].Level = "3"
				// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
				dd[0].Pid = depart.Id
				merittitle, err := models.GetAdminMeritbyId(v5.MeritId) //因为这个数据库只是科室和分类的对应表
				if err != nil {
					beego.Error(err)
				}
				dd[0].Title = merittitle.Title //名称
				//由价值分类取得所有价值列表
				merits, err := models.GetAdminMerit(v5.MeritId)
				if err != nil {
					beego.Error(err)
				}
				//循环价值分类，取得价值列表
				for _, v6 := range merits {
					ee := make([]MeritCategory, 1)
					ee[0].Id = v6.Id
					ee[0].Level = "4"
					ee[0].Pid = v5.Id

					ee[0].Title = v6.Title //名称
					//根据userid和meritid，取得项数和分值

					// ee[0].Tags[0]=//有几项
					// ee[0].Tags[1]=//总共多少分
					// beego.Info(users[i3].Nickname)
					ee[0].Selectable = true
					meritcategory = append(meritcategory, ee...)
				}
				dd[0].List = meritcategory
				meritcategory = make([]MeritCategory, 0) //再把slice置0
				// dd[0].Title = strconv.FormatInt(meritcates[i3].Id, 10) //名称——关键，把价值分类当作科室名
				dd[0].Selectable = false
				meritsecoffice = append(meritsecoffice, dd...)
			}
			aa[0].Selectable = true
		}
		// aa[0].Tags[0] = count + depcount
		// count = 0
		// depcount = 0
		aa[0].Secoffice = meritsecoffice
		// aa[0].Selectable = false                   //默认是false_点击展开
		meritsecoffice = make([]MeritSecoffice, 0) //再把slice置0
		meritdepartment = append(meritdepartment, aa...)
		c.Data["json"] = meritdepartment
	case "3": //主任登录
		depart, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		// aa := make([]MeritDepartment, 1)
		// aa[0].Id = depart.Id
		// aa[0].Level = "1"
		// aa[0].Title = depart.Title //分院名称
		//由分院id和科室名称取得科室
		secoffice, err := models.GetAdminDepartbyidtitle(depart.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		//如果返回科室为空，则直接取得价值分类
		//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
		//应该是反过来找出所有没有科室字段的人员，把他放在部门下
		// if len(secoffice) > 0 {
		// for _, v2 := range secoffice {
		bb := make([]MeritSecoffice, 1)
		bb[0].Id = secoffice.Id
		bb[0].Level = "2"
		bb[0].Pid = depart.Id
		bb[0].Title = secoffice.Title //科室名称
		// beego.Info(category2[i2].Title)
		//根据分院和科室查所有价值分类：项目管理类的id
		meritcates, err := models.GetMeritsbySec(depart.Title, secoffice.Title)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(meritcates)
		//取得所有儿子价值，如果有孙子价值（大型、中型……）则用孙子

		for _, v3 := range meritcates {
			cc := make([]MeritCategory, 1)
			cc[0].Id = v3.Id
			cc[0].Level = "3"
			cc[0].Pid = secoffice.Id
			merittitle, err := models.GetAdminMeritbyId(v3.MeritId) //因为这个数据库只是科室和分类的对应表
			if err != nil {
				beego.Error(err)
			}
			beego.Info(merittitle)         //20 项目管理类
			cc[0].Title = merittitle.Title //名称
			//由价值分类(项目管理）取得所有儿子价值列表
			adminmerits, err := models.GetAdminMeritbyPid(v3.MeritId)
			if err != nil {
				beego.Error(err)
			}
			beego.Info(adminmerits) //专业负责类，
			//循环价值分类，取得价值列表
			for _, v4 := range adminmerits {
				ee := make([]MeritList, 1)
				ee[0].Id = v4.Id
				ee[0].Level = "4"
				ee[0].Pid = v3.Id

				ee[0].Title = v4.Title //名称

				//根据userid和meritid，取得项数和分值
				//取得用户的价值topic数量和分值
				// merits, err := models.GetMerit(0, uid, 3)
				merits, err := models.GetMeritTopic(v4.Id, uid, 3)
				if err != nil {
					beego.Error(err)
				}
				beego.Info(merits)
				//用聚合查询查出分值
				ee[0].Tags[0] = len(merits) //有几项
				for _, v5 := range merits {
					ee[0].Tags[1] = ee[0].Tags[1] + v5.Mark //总共多少分
				}

				ee[0].Selectable = true
				meritlist = append(meritlist, ee...)
			}
			cc[0].List = meritlist

			// cc[0].Tags[0]=//有几项
			// cc[0].Tags[1]=//总共多少分

			// cc[0].Selectable = false
			meritlist = make([]MeritList, 0)
			meritcategory = append(meritcategory, cc...)
		}

		// bb[0].Tags[0] = strconv.Itoa(count)
		bb[0].List = meritcategory
		bb[0].Selectable = true
		meritcategory = make([]MeritCategory, 0) //再把slice置0
		meritsecoffice = append(meritsecoffice, bb...)
		// depcount = depcount + count //部门人员数=科室人员数相加
		// }
		// aa[0].Secoffice = achsecoffice
		// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		// achdepart = append(achdepart, aa...)
		// }
		// aa[0].Secoffice = meritsecoffice
		// aa[0].Selectable = false                   //默认是false_点击展开
		// meritsecoffice = make([]MeritSecoffice, 0) //再把slice置0
		// meritdepartment = append(meritdepartment, aa...)
		c.Data["json"] = meritsecoffice
	default: //如果是用户role=4登录
		depart, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		// aa := make([]MeritDepartment, 1)
		// aa[0].Id = depart.Id
		// aa[0].Level = "1"
		// aa[0].Title = depart.Title //分院名称
		//由分院id和科室名称取得科室
		secoffice, err := models.GetAdminDepartbyidtitle(depart.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		//如果返回科室为空，则直接取得价值分类
		//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
		//应该是反过来找出所有没有科室字段的人员，把他放在部门下
		// bb := make([]MeritSecoffice, 1)
		// bb[0].Id = secoffice.Id
		// bb[0].Level = "2"
		// bb[0].Pid = depart.Id
		// bb[0].Title = secoffice.Title //科室名称
		// beego.Info(category2[i2].Title)
		//根据分院和科室查所有价值分类：项目管理类的id
		meritcates, err := models.GetMeritsbySec(depart.Title, secoffice.Title)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(meritcates)
		//取得所有儿子价值，如果有孙子价值（大型、中型……）则用孙子
		for _, v3 := range meritcates {
			cc := make([]MeritCategory, 1)
			cc[0].Id = v3.Id
			cc[0].Level = "3"
			cc[0].Pid = secoffice.Id
			merittitle, err := models.GetAdminMeritbyId(v3.MeritId) //因为这个数据库只是科室和分类的对应表
			if err != nil {
				beego.Error(err)
			}
			beego.Info(merittitle)         //20 项目管理类
			cc[0].Title = merittitle.Title //名称
			//由价值分类(项目管理）取得所有儿子价值列表
			adminmerits, err := models.GetAdminMeritbyPid(v3.MeritId)
			if err != nil {
				beego.Error(err)
			}
			beego.Info(adminmerits) //专业负责类，
			//循环价值分类，取得价值列表
			for _, v4 := range adminmerits {
				ee := make([]MeritList, 1)
				ee[0].Id = v4.Id
				ee[0].Level = "4"
				ee[0].Pid = v3.Id

				ee[0].Title = v4.Title //名称

				merits, err := models.GetMeritTopic(v4.Id, uid, 3)
				if err != nil {
					beego.Error(err)
				}
				beego.Info(merits)
				//用聚合查询查出分值
				ee[0].Tags[0] = len(merits) //有几项
				for _, v5 := range merits {
					ee[0].Tags[1] = ee[0].Tags[1] + v5.Mark //总共多少分
				}

				ee[0].Selectable = true
				meritlist = append(meritlist, ee...)
			}
			cc[0].List = meritlist
			meritlist = make([]MeritList, 0)
			meritcategory = append(meritcategory, cc...)
		}

		// bb[0].List = meritcategory
		// bb[0].Selectable = true
		// meritcategory = make([]MeritCategory, 0) //再把slice置0
		// meritsecoffice = append(meritsecoffice, bb...)

		// aa[0].Secoffice = meritsecoffice
		// aa[0].Selectable = false                   //默认是false_点击展开
		// meritsecoffice = make([]MeritSecoffice, 0) //再把slice置0
		// meritdepartment = append(meritdepartment, aa...)
		c.Data["json"] = meritcategory
	}

	c.TplName = "merit/merit.tpl"
}

//下面关于用户名的取得需要重新来改改
//上面那个是显示侧栏
//这个是显示右侧iframe框架内容——科室内人员情况统计
func (c *MeritController) Secofficeshow() {
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
		c.Redirect("/login?url="+route, 302)
		return
	}
	//由uname取得user
	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}

	//分院——科室——价值分类——价值
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
		beego.Info(t2)
	} else {
		t2 = time.Now()
		// beego.Info(t1):2016-08-19 23:27:29.7463081 +0800 CST
		// starttime, _ := time.Parse("2006-01-02", starttime1)
		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
		// beego.Info(t2)
	}

	switch level {
	case "1": //如果是部门、分院，则显示全部科室
		categoryname, err := models.GetAdminDepartbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		//权限判断，并且属于这个分院
		if role == "1" || role == "2" && user.Department == categoryname.Title {
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["Deptitle"] = categoryname.Title
			c.TplName = "merit/merit_secoffice.tpl" //"merit_depoffice.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			return
		}
	case "2": //如果是科室，则显示全部人员情况，此时secid为科室id
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
		if role == "1" || role == "3" && user.Secoffice == categoryname.Title || role == "2" && user.Department == categoryname1.Title {
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Sectitle"] = categoryname.Title
			c.Data["Level"] = level
			c.TplName = "merit/merit_secoffice.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			return
		}
	case "4": //如果是个人，则显示个人详细情况。如果是价值列表，此时secid为价值id
		//取得价值
		merit, err := models.GetAdminMeritbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		//取得所有子价值和分值
		meritarray, err := models.GetAdminMerit(merit.Id)
		if err != nil {
			beego.Error(err)
		}
		//取得价值分类
		meritcate, err := models.GetAdminMeritbyId(merit.ParentId)
		if err != nil {
			beego.Error(err)
		}
		//进行选择列表拆分
		// array1 := strings.Split(merit.List, ",")
		// beego.Info(merit.List)
		slice1 := make([]string, 0)
		for _, v := range meritarray {
			slice1 = append(slice1, v.Title)
		}
		// c.Data["list"] = slice1
		c.Data["list"] = meritarray
		c.Data["select2"] = slice1 //array1//这个用meritarray中的[]title代替
		// topics, err := models.GetAllMeritTopic(user.Id)
		// 	c.Data["topics"] = topics
		c.Data["UserNickname"] = user.Nickname
		c.Data["Meritcate"] = meritcate
		c.Data["Merit"] = merit
		c.Data["Secid"] = secid
		c.Data["Level"] = level
		// 	//查出所有用户的价值资料
		// 	//前提是价值资料里要带用户id
		// 	category5, _ := models.GetAllCategory()
		// 	if err != nil {
		// 		beego.Error(err)
		// 	}
		// 	c.Data["category"] = category5

		//这里给定secid的meritcategoryid和名，merit的名和id
		c.TplName = "merit/merit_employee.tpl"
	default:
		// case "3": //科室主任
		//分2部分，一部分是已经完成状态的，state是4，另一部分是状态分别是3待审查通过,2，1的
		usernickname := models.GetUserByUserId(secid1)
		//1.进行权限读取，室主任以上并且属于这个科室，或者或本人
		if role == "1" || role == "3" && user.Secoffice == usernickname.Secoffice || role == "2" && user.Department == usernickname.Department || user.Nickname == usernickname.Nickname {
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			//下面这个catalogs用于employee_show.tpl
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["UserNickname"] = usernickname.Nickname

			if key == "modify" { //新窗口显示处理页面
				//如果是本人，则显示
				c.TplName = "merit/merit_employeework.tpl"
			} else { //直接查看页面
				//如果是本人，则显示带处理按钮的
				if usernickname.Nickname == user.Nickname {
					c.Data["IsMe"] = true
				} else { //别人查看，不显示处理按钮
					c.Data["IsMe"] = false
				}
				c.TplName = "merit/merit_myself.tpl"
			}
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			return
		}
	}
}

//上面那个是显示右侧页面
//这个是填充数据——科室内人员成果情况统计
func (c *MeritController) SecofficeData() {
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	// secid1, err := strconv.ParseInt(secid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// level := c.Input().Get("level")
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

	//取得科室名称
	// secoffice, err := models.GetAdminDepartbyId(secid1)
	// if err != nil {
	// 	beego.Error(err)
	// }

	// employeevalue := make([]models.Employeeachievement, 0)
	//根据科室id查所有员工
	users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
	beego.Info(users)
	if err != nil {
		beego.Error(err)
	}

	// var numbers, marks int
	// slice1 := make([]Person, 0)
	slice1 := make([]*models.UserMeritTopics, 0)
	for _, v1 := range users {
		//根据价值id和用户id，得到成果，统计数量和分值
		//取得用户的价值topic数量和分值
		// merits, err := models.GetMerit(0, users[i1].Id, 3)
		beego.Info(v1.Id)
		merittopics, err := models.GetMeritTopics(v1.Id, 3, true)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(merittopics)
		slice1 = append(slice1, merittopics...)
		// for _, v := range merits {
		//根据choose取得adminmerit分值
		// adminmerit, err := models.GetAdminMeritbyId(v.MeritId)
		// adminmeritmark, err := models.GetAdminMeritMarkbyId(v.MeritId)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// var ff string
		// if adminmerit.Mark == "" {
		// 	// 如果mark为空，则寻找选择列表的分值，如果不为空，则直接用价值的分值
		// 	// 进行选择列表拆分
		// 	array1 := strings.Split(adminmerit.List, ",")
		// 	array2 := strings.Split(adminmerit.ListMark, ",")
		// 	for i2, v2 := range array1 {
		// 		if v2 == v.Choose {
		// 			ff = array2[i2]
		// 		}
		// 	}
		// } else {
		// 	ff = adminmerit.Mark
		// }
		// markint, err := strconv.Atoi(ff)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// marks = marks + markint
		// marks = marks + adminmeritmark.Mark
		// }
		// numbers = len(merits)
		// marks1 = marks1 + marks
		// numbers1 = numbers1 + numbers
		// aa := make([]Person, 1)
		// aa[0].Id = users[i1].Id //这里用for i1,v1,然后用v1.Id一样的意思
		// aa[0].Name = users[i1].Nickname
		// aa[0].Department = users[i1].Department
		// aa[0].Secoffice = users[i1].Secoffice
		// aa[0].Numbers = numbers
		// aa[0].Marks = marks
		// aa[0].Marks = marks
		// aa[0].UserId = v1.Id
		// slice1 = append(slice1, aa...)
		// marks = 0
		// numbers = 0
	}

	// for _, v := range users {
	// 	//由username查出所有编制成果总分、设计总分……合计
	// 	employee, _, err := models.Getemployeevalue(v.Nickname, t1, t2)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	employeevalue = append(employeevalue, employee...)
	// }
	//排序
	// pList := person1(slice1)
	// sort.Sort(pList)
	// sort.Sort(slice1)
	c.Data["Starttime"] = t1
	c.Data["Endtime"] = t2
	// c.Data["Secid"] = secid
	// c.Data["Sectitle"] = secoffice.Title
	// c.Data["Level"] = level
	// c.Data["json"] = pList
	c.Data["json"] = slice1
	c.ServeJSON()
}

// @Title get usertopics list
// @Description get usertopics by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /testxorm [get]
//测试取得xorm取得结构体嵌套_查不出来，奇怪
func (c *MeritController) TestXorm() {
	// merittopics, err := models.GetMeritTopicUser(1490, 3)
	// merittopics, err := models.GetMeritTopics2(1490, 3)
	merittopics, err := models.GetMeritTopic2(27, 1490, 3)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = merittopics
	c.ServeJSON()
}

//这个是显示右侧iframe框架内容——用户登录默认页面，自己的全部价值内容列表
//如果是主任以上权限人查看，则id代表用户名id，个人查看，id则代表价值id
//要修改——已经完成的*************
func (c *MeritController) Myself() {
	var userid int64
	var err error
	uid := c.Input().Get("userid")
	if uid != "" {
		userid, err = strconv.ParseInt(uid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else { //显示登录用户的
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
			c.Redirect("/login?url="+route, 302)
			return
		}
		userid = uid
	}

	merittopics, err := models.GetMyselfMeritTopic(userid, 3, true) //这个其实是merittopic
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = merittopics
	c.ServeJSON()
}

//显示指定价值下的登录用户待提交，已提交，已经完成的价值内容——计算分值
func (c *MeritController) MeritSend() {
	id := c.Ctx.Input.Param(":id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		beego.Error(err)
	}
	mid := c.Input().Get("meritid")
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//如果是主任以上权限人查看，则id代表用户名id，个人查看，id则代表价值id
	//1.首先判断是否注册
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
		c.Redirect("/login?url="+route, 302)
		return
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(idint)
	// beego.Info(user.Id)
	merittopics, err := models.GetMeritTopic(midNum, user.Id, idint)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = merittopics
	c.ServeJSON()
}

//显示需要审核的价值内容——计算分值
func (c *MeritController) MeritExamined() {
	//如果是主任以上权限人查看，则id代表用户名id，个人查看，id则代表价值id
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
		c.Redirect("/login?url="+route, 302)
		return
	}

	// merittopics := make([]MeritTopicSlice, 0)

	merittopics := make([]*models.MyMeritTopic, 0)
	//如果权限<=2级，则取出科室所有人员，循环人员取出待审核的价值内容
	if roleint <= 2 {
		depart, err := models.GetAdminDepart(0) //得到多个分院
		if err != nil {
			beego.Error(err)
		}
		beego.Info(depart)
		for _, v1 := range depart {
			secoffice, err := models.GetAdminDepart(v1.Id) //得到多个科室
			if err != nil {
				beego.Error(err)
			}
			beego.Info(secoffice)
			//如果返回科室为空，则直接取得价值分类
			//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
			//应该是反过来找出所有没有科室字段的人员，把他放在部门下
			if len(secoffice) > 0 {
				for _, v2 := range secoffice {
					//根据分院和科室查所有员工
					users, _, err := models.GetUsersbySec(v1.Title, v2.Title) //得到员工姓名
					if err != nil {
						beego.Error(err)
					}
					beego.Info(users)
					for _, v3 := range users {
						// cc, err := getusermerit(v3.Id)
						cc, err := models.GetMyselfMeritTopic(v3.Id, 2, true)
						if err != nil {
							beego.Error(err)
						}
						merittopics = append(merittopics, cc...)
					}
				}

			}
			//查出所有有这个部门但科室名为空的人员
			//根据分院查所有员工
			// beego.Info(v1.Title)
			users, _, err := models.GetUsersbySecOnly(v1.Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(users)
			for _, v4 := range users {
				// cc, err := getusermerit(v4.Id)
				cc, err := models.GetMyselfMeritTopic(v4.Id, 2, true)
				if err != nil {
					beego.Error(err)
				}
				merittopics = append(merittopics, cc...)
			}
		}
	}
	c.Data["json"] = merittopics
	c.ServeJSON()
}

//函数——取得用户的价值内容，待审核的
//作废，用cc, err := models.GetMyselfMeritTopic(v3.Id,2)代替
func getusermerit(userid int64) (merittopics []MeritTopicSlice, err error) {
	merits, err := models.GetMerit(0, userid, 2)
	if err != nil {
		beego.Error(err)
	}
	// merittopics := make([]MeritTopicSlice, 0)
	for _, v := range merits {
		//根据choose取得adminmerit分值
		adminmerit, err := models.GetAdminMeritbyId(v.MeritId)
		if err != nil {
			beego.Error(err)
		}
		//价值分类
		meritcate, err := models.GetAdminMeritbyId(adminmerit.ParentId)
		if err != nil {
			beego.Error(err)
		}
		// var ff string
		// if adminmerit.Mark == "" {
		// 	// 如果mark为空，则寻找选择列表的分值，如果不为空，则直接用价值的分值
		// 	// 进行选择列表拆分
		// 	array1 := strings.Split(adminmerit.List, ",")
		// 	array2 := strings.Split(adminmerit.ListMark, ",")
		// 	for i1, v1 := range array1 {
		// 		if v1 == v.Choose {
		// 			ff = array2[i1]
		// 		}
		// 	}
		// } else {
		// 	ff = adminmerit.Mark
		// }
		// var markint int
		// if ff != "" {
		// 	markint, err = strconv.Atoi(ff)
		// 	if err != nil {
		// 		beego.Error(err)
		// 	}
		// }
		adminmeritmark, err := models.GetAdminMeritMarkbyId(v.MeritId)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]MeritTopicSlice, 1)
		aa[0].Id = v.Id
		aa[0].MeritId = v.MeritId
		aa[0].MeritCate = meritcate.Title //价值类型
		aa[0].Merit = adminmerit.Title    //价值
		aa[0].UserId = v.UserId
		//usernickname
		user := models.GetUserByUserId(v.UserId)
		aa[0].UserNickName = user.Nickname
		aa[0].Title = v.Title
		// aa[0].Choose = v.Choose
		aa[0].Content = v.Content
		aa[0].State = v.State
		// aa[0].Mark = markint
		aa[0].Mark = adminmeritmark.Mark
		aa[0].Created = v.Created
		aa[0].Updated = v.Updated
		merittopics = append(merittopics, aa...)
	}
	return merittopics, err
}

// @Title post user merit
// @Description post user merit
// @Param meritid query string true "The id of merit"
// @Param title query string true "The title of merit_topic"
// @Param content query string true "The content of merit_topic"
// @Param active query string true "The actuve of merit_topic"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 merit not found
// @router /addmerit [post]
// 用户进行价值添加
func (c *MeritController) AddMerit() {
	// mcid := c.Input().Get("mcid")
	// mcidNum, err := strconv.ParseInt(mcid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	mid := c.Input().Get("meritid")
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	active := c.Input().Get("active")
	var activebool bool
	if active == "true" {
		activebool = true
	} else {
		activebool = false
	}
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	_, err = models.AddMerit(midNum, uid, title, content, activebool)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//用户修改价值
func (c *MeritController) UpdateMerit() {
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateMerit(id, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/merit.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "修改merit" + pk)
	// logs.Close()
}

//用户传递价值
func (c *MeritController) SendMerit() {
	// name := "state" //c.Input().Get("name")
	value := c.Input().Get("state")
	pk := c.Input().Get("meritid")
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	stateint, err := strconv.Atoi(value)
	if err != nil {
		beego.Error(err)
	}
	stateint1 := stateint + 1
	statestring := strconv.Itoa(stateint1)
	// beego.Info(statestring)
	// beego.Info(id)
	err = models.UpdateMerit(id, "State", statestring)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/merit.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "修改merit" + pk)
	// logs.Close()
}

//用户回退价值
func (c *MeritController) DownSendMerit() {
	// name := "state" //c.Input().Get("name")
	value := c.Input().Get("state")
	pk := c.Input().Get("meritid")
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	stateint, err := strconv.Atoi(value)
	if err != nil {
		beego.Error(err)
	}
	stateint = stateint - 1
	statestring := strconv.Itoa(stateint)

	err = models.UpdateMerit(id, "State", statestring)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/merit.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "修改merit" + pk)
	// logs.Close()
}

//删除价值
func (c *MeritController) Delete() {
	mid := c.Input().Get("meritid")
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteMerit(midNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//显示——修改价值结构中的项目
// func (c *MeritController) UpdateMerit() {
// 	//1.首先判断是否注册
// 	if !checkAccount(c.Ctx) {
// 		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/login?url="+route, 302)
// 		// c.Redirect("/login", 302)
// 		return
// 	}
// 	//2.取得文章的作者
// 	//3.由用户id取得用户名
// 	//4.取得客户端用户名
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		c.Data["Uname"] = v.(string)
// 	}
// 	// uname := v.(string) //ck.Value
// 	//4.取出用户的权限等级
// 	role, _ := checkRole(c.Ctx) //login里的
// 	// beego.Info(role)
// 	//5.进行逻辑分析：
// 	// rolename, _ := strconv.ParseInt(role, 10, 64)
// 	if role > 2 { //
// 		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/roleerr?url="+route, 302)
// 		// c.Redirect("/roleerr", 302)
// 		return
// 	}

// 	//4.取得价值列表choose和mark
// 	id := c.Input().Get("id")
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	category, err := models.GetCategory(idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//如果mark为空，则寻找选择列表的分值
// 	//进行选择列表拆分
// 	array1 := strings.Split(category.List, ",")
// 	array2 := strings.Split(category.ListMark, ",")
// 	//进行选择列表拆分
// 	// array1 := strings.Split(category.List, ",")
// 	slice1 := make([]List1, 0)
// 	for i, v := range array1 {
// 		ee := make([]List1, 1)
// 		ee[0].Choose = v
// 		ee[0].Mark1 = array2[i]
// 		slice1 = append(slice1, ee...)
// 	}

// 	slice2 := make([]List1, 0)
// 	for _, w := range array2 {
// 		ff := make([]List1, 1)
// 		ff[0].Mark1 = w
// 		slice2 = append(slice2, ff...)
// 	}
// 	c.TplName = "admin_json_modify.tpl"
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	c.Data["category"] = category
// 	c.Data["list"] = slice1
// 	c.Data["mark"] = slice2
// 	// beego.Info(slice2)
// 	c.Data["Cid"] = id
// 	c.Data["IsCategory"] = true
// }

//提交修改—修改价值结构中的项目
// func (c *MeritController) UpdateMerit() {
// 	title := c.Input().Get("title")
// 	mark := c.Input().Get("mark")
// 	list := c.Input().Get("list")
// 	listmark := c.Input().Get("listmark")
// 	//4.取得价值列表choose和mark
// 	id := c.Input().Get("id")
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	err = models.Modifyjson(idNum, title, mark, "", list, listmark)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/json", 301)
// }

//删除价值结构中的项目
// func (c *MeritController) DeleteMerit() {
// 	id := c.Input().Get("id")
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	err = models.Deletejson(idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/json", 301)
// }

//传递merit
//退回
//已提交——待审核
//待提交
//待审核
//导入

//导入json数据
// func (c *MeritController) ImportJson() {
// 	//获取上传的文件
// 	_, h, err := c.GetFile("json")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	var path string
// 	if h != nil {
// 		//保存附件
// 		path = ".\\attachment\\" + h.Filename
// 		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
// 		err = c.SaveToFile("json", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	contents, _ := ioutil.ReadFile(path)
// 	var r List6
// 	//var r JsonStruct//空结构对于系统unmarshal不行。
// 	//	var r map[string]interface{}//空接口可行
// 	//	var r []interface{}//这个对于系统unmarshal不行
// 	err = json.Unmarshal([]byte(contents), &r)
// 	if err != nil {
// 		fmt.Printf("err was %v", err)
// 	}
// 	// fmt.Println(r)
// 	// beego.Info(r)

// 	js, err := simplejson.NewJson([]byte(contents))
// 	if err != nil {
// 		panic("json format error")
// 	}
// 	//1.获取省水利院
// 	text, err := js.Get("text").String()
// 	//存入数据库——单位
// 	Id, err := models.AddCategory(0, text, "", "", "")
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	arr, err := js.Get("nodes").Array()
// 	if err != nil {
// 		fmt.Println("decode error: get array failed!")
// 		// return
// 	}
// 	for i, _ := range arr {
// 		// beego.Info(v)是map[string]interface{}
// 		text1, _ := js.Get("nodes").GetIndex(i).Get("text").String()
// 		//存入数据库——分院
// 		Id1, err := models.AddCategory(Id, text1, "", "", "")
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		arr1, err := js.Get("nodes").GetIndex(i).Get("nodes").Array()
// 		for i1, _ := range arr1 {
// 			text2, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("text").String()
// 			//存入数据库——科室
// 			Id2, err := models.AddCategory(Id1, text2, "", "", "")
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			arr2, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").Array()
// 			for i2, _ := range arr2 {
// 				text3, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("text").String()
// 				//存入数据库——管理类
// 				Id3, err := models.AddCategory(Id2, text3, "", "", "")
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				arr3, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").Array()
// 				for i3, _ := range arr3 {
// 					text4, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("text").String()
// 					text5, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("mark").String()
// 					//循环取出选择项，拼接字符串
// 					//循环取出每个选择项的打分，拼接字符串
// 					arr4, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").Array()
// 					var text8, text9 string
// 					for i4, _ := range arr4 {
// 						text6, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("text").String()
// 						text7, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("mark").String()
// 						if i == 0 {
// 							text8 = text6
// 							text9 = text7
// 						} else {
// 							text8 = text8 + "," + text6
// 							text9 = text9 + "," + text7
// 						}
// 					}
// 					//存入数据库——项目负责人
// 					// url:="/"+"add?id="+
// 					_, err = models.AddCategory(Id3, text4, text5, text8, text9)
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 				}

// 			}
// 		}
// 	}
// }

//根据conf目录下的json.json文件初始化价值结构
// func (c *MeritController) InitJson() {
// 	contents, _ := ioutil.ReadFile("./conf/json.json")
// 	var r List6
// 	err := json.Unmarshal([]byte(contents), &r)
// 	if err != nil {
// 		fmt.Printf("err was %v", err)
// 	}
// 	// fmt.Println(r)
// 	// beego.Info(r)

// 	js, err := simplejson.NewJson([]byte(contents))
// 	if err != nil {
// 		panic("json format error")
// 	}
// 	//1.获取省水利院
// 	text, err := js.Get("text").String()
// 	//存入数据库——单位
// 	Id, err := models.AddCategory(0, text, "", "", "")
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	arr, err := js.Get("nodes").Array()
// 	if err != nil {
// 		fmt.Println("decode error: get array failed!")
// 		// return
// 	}
// 	for i, _ := range arr {
// 		// beego.Info(v)是map[string]interface{}
// 		text1, _ := js.Get("nodes").GetIndex(i).Get("text").String()
// 		//存入数据库——分院
// 		Id1, err := models.AddCategory(Id, text1, "", "", "")
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		arr1, err := js.Get("nodes").GetIndex(i).Get("nodes").Array()
// 		for i1, _ := range arr1 {
// 			text2, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("text").String()
// 			//存入数据库——科室
// 			Id2, err := models.AddCategory(Id1, text2, "", "", "")
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			arr2, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").Array()
// 			for i2, _ := range arr2 {
// 				text3, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("text").String()
// 				//存入数据库——管理类
// 				Id3, err := models.AddCategory(Id2, text3, "", "", "")
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				arr3, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").Array()
// 				for i3, _ := range arr3 {
// 					text4, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("text").String()
// 					text5, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("mark").String()
// 					//循环取出选择项，拼接字符串
// 					//循环取出每个选择项的打分，拼接字符串
// 					arr4, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").Array()
// 					var text8, text9 string
// 					for i4, _ := range arr4 {
// 						text6, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("text").String()
// 						text7, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("mark").String()
// 						if i == 0 {
// 							text8 = text6
// 							text9 = text7
// 						} else {
// 							text8 = text8 + "," + text6
// 							text9 = text9 + "," + text7
// 						}
// 					}
// 					//存入数据库——项目负责人
// 					// url:="/"+"add?id="+
// 					_, err = models.AddCategory(Id3, text4, text5, text8, text9)
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 				}

// 			}
// 		}
// 	}
// }

// func NewJsonStruct() *JsonStruct {
// 	return &JsonStruct{}
// }

// func (self *JsonStruct) Load(filename string, v interface{}) {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return
// 	}
// 	datajson := []byte(data)
// 	err = json.Unmarshal(datajson, v)
// 	if err != nil {
// 		return
// 	}
// }

// type ValueTestAtmp struct {
// 	StringValue    string
// 	NumericalValue int
// 	BoolValue      bool
// }

// type testdata struct {
// 	ValueTestA ValueTestAtmp
// }

// package main
// import (
//     "encoding/json"
//     "fmt"
// )
// func test() {
// 	b := []byte(`{
//     "Title":"go programming language",
//     "Author":["john","ada","alice"],
//     "Publisher":"qinghua",
//     "IsPublished":true,
//     "Price":99
//   }`)
// 	//先创建一个目标类型的实例对象，用于存放解码后的值
// 	var inter interface{}
// 	err := json.Unmarshal(b, &inter)
// 	if err != nil {
// 		fmt.Println("error in translating,", err.Error())
// 		return
// 	}
// 	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
// 	book, ok := inter.(map[string]interface{})
// 	//然后通过for循环一一访问解码后的目标数据
// 	if ok {
// 		for k, v := range book {
// 			switch vt := v.(type) {
// 			case float64:
// 				fmt.Println(k, " is float64 ", vt)
// 			case string:
// 				fmt.Println(k, " is string ", vt)
// 			case []interface{}:
// 				fmt.Println(k, " is an array:")
// 				for i, iv := range vt {
// 					fmt.Println(i, iv)
// 				}
// 			default:
// 				fmt.Println("illegle type")
// 			}
// 		}
// 	}
// }

// 今天遇到个接口需要处理一个json的map类型的数组，开始想法是用simple—json
// 里的Array读取数组，然后遍历数组取出每个map，然后读取对应的值，在进行后续操作，
// 貌似很简单的工作，却遇到了一个陷阱。
// Json格式类似下边：
// {"code":0
// ,"request_id": xxxx
// ,"code_msg":""
// ,"body":[{
//         "device_id": "xxxx"
//         ,"device_hid": "xxxx"
// }]
// , "count":0}
//     很快按上述想法写好了带码，但是以外发生了，编译不过，看一看代码逻辑没有
// 问题，问题出在哪里呢？
//     原来是interface{} Array方法返回的是一个interface{}类型的，我们都在golang
// 里interface是一个万能的接受者可以保存任意类型的参数，但是却忽略了一点，它是
// 不可以想当然的当任意类型来用，在使用之前一定要对interface类型进行判断。我开始
// 就忽略了这点，想当然的使用interface变量造成了错误。
//     下面写了个小例子

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/bitly/go-simplejson"
// )

// func test2() {
// 	//拼凑json   body为map数组
// 	var rbody []map[string]interface{}
// 	t := make(map[string]interface{})
// 	t["device_id"] = "dddddd"
// 	t["device_hid"] = "ddddddd"

// 	rbody = append(rbody, t)
// 	t1 := make(map[string]interface{})
// 	t1["device_id"] = "aaaaa"
// 	t1["device_hid"] = "aaaaa"

// 	rbody = append(rbody, t1)

// 	cnnJson := make(map[string]interface{})
// 	cnnJson["code"] = 0
// 	cnnJson["request_id"] = 123
// 	cnnJson["code_msg"] = ""
// 	cnnJson["body"] = rbody
// 	cnnJson["page"] = 0
// 	cnnJson["page_size"] = 0

// 	b, _ := json.Marshal(cnnJson)
// 	cnnn := string(b)
// 	fmt.Println("cnnn:%s", cnnn)
// 	cn_json, _ := simplejson.NewJson([]byte(cnnn))
// 	cn_body, _ := cn_json.Get("body").Array()

// 	for _, di := range cn_body {
// 		//就在这里对di进行类型判断
// 		newdi, _ := di.(map[string]interface{})
// 		device_id := newdi["device_id"]
// 		device_hid := newdi["device_hid"]
// 		fmt.Println(device_hid, device_id)
// 	}

// }

// 第一步，得到json的内容
// contents, _ := ioutil.ReadAll(res.Body)
// js, js_err := simplejson.NewJson(contents)

// 第二部，根据json的格式，选择使用array或者map储存数据
// var nodes = make(map[string]interface{})
// nodes, _ = js.Map()

// 第三步，将nodes当作map处理即可，如果map的value仍是一个json结构，回到第二步。
// for key,_ := range nodes {
// ...
// }
