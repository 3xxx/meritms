package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/meritms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

// CMSADMIN API
type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	beego.Info(role)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	c.TplName = "admin/admin.tpl"
}

func (c *AdminController) Admin() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	if role == "1" {
		id := c.Ctx.Input.Param(":id")
		c.Data["Id"] = id

		switch id {
		//基本设置-项目分级目录
		case "010": //日历事件
			c.TplName = "admin/admin_calendar.tpl"
		case "011": //基本设置
			c.TplName = "admin/admin_base.tpl"
		case "012": //分级目录
			c.TplName = "admin/admin_category.tpl"
		case "013": //搜索引擎
			c.TplName = "admin/admin_spiderip.tpl"
		case "014": //升级数据库
			c.TplName = "admin/admin_updatedatabase.tpl"

			//组织架构-用户-角色
		case "021": //组织结构
			c.TplName = "admin/admin_department.tpl"
		case "022": //用户-组织结构
			c.TplName = "admin/admin_users.tpl"
		case "023": //用户-角色
			c.TplName = "admin/admin_users.tpl"
		case "024": //角色-成果价值
			achsecoffice := make([]AchSecoffice, 0)
			achdepart := make([]AchDepart, 0)
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
						//根据分院和科室查所有员工
						bb[0].Selectable = true
						// achemployee = make([]AchEmployee, 0) //再把slice置0
						achsecoffice = append(achsecoffice, bb...)
						// depcount = depcount + count //部门人员数=科室人员数相加
					}
				}
				//查出所有有这个部门但科室名为空的人员
				//根据分院查所有员工
				aa[0].Secoffice = achsecoffice
				aa[0].Selectable = true                //默认是false点击展开
				achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				achdepart = append(achdepart, aa...)
			}
			c.Data["json"] = achdepart
			c.TplName = "admin/admin_roleachieve.tpl"
		//项目设置
		case "031": //分级目录
			c.TplName = "admin/admin_category.tpl"
		case "032": //角色权限分配
			c.TplName = "admin/admin_role.tpl"
		case "033": //项目编辑
			c.TplName = "admin/admin_projectstree.tpl"
		case "034": //项目目录快捷编辑
			c.TplName = "admin/admin_projectseditor.tpl"
		case "035": //同步IP async
			c.TplName = "admin/admin_projectsynch.tpl"
		case "036": //IP地址段
			c.TplName = "admin_ipsegment.tpl"

			//成果设置
		case "041": //成果类型
			c.TplName = "admin/admin_achievcategory.tpl"
		case "042": //科室成果类型
			c.TplName = "admin/admin_departachievcate.tpl"
		case "043": //本周成果编辑
			c.TplName = "admin/admin_achievementseditor.tpl"
		case "044": //本月成果编辑
			c.TplName = "admin/admin_projectsynch.tpl"
		case "045": //上月成果编辑
			c.TplName = "admin/admin_projectcaterole.tpl"
		case "046": //当年成果编辑
			c.TplName = "admin/admin_projectcaterole.tpl"

			//价值设置
		case "051": //定义价值
			c.TplName = "admin/admin_merit.tpl"
		case "052": //科室价值
			achsecoffice := make([]AchSecoffice, 0)
			achdepart := make([]AchDepart, 0)
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
						bb[0].Selectable = true
						// achemployee = make([]AchEmployee, 0) //再把slice置0
						achsecoffice = append(achsecoffice, bb...)
						// depcount = depcount + count //部门人员数=科室人员数相加
					}
				}
				aa[0].Secoffice = achsecoffice
				aa[0].Selectable = true                //默认是false点击展开
				achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				achdepart = append(achdepart, aa...)
			}
			c.Data["json"] = achdepart
			c.TplName = "admin/admin_secofficemerit.tpl"

		default:
			c.TplName = "admin/admin_calendar.tpl"
		}
	}
}

//添加ip地址段
func (c *AdminController) AddIpsegment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminIpsegment(title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//修改ip地址段
func (c *AdminController) UpdateIpsegment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminIpsegment(cidNum, title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//删除IP
func (c *AdminController) DeleteIpsegment() {
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
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteAdminIpsegment(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
	Createip()
}

//查询IP地址段
func (c *AdminController) Ipsegment() {
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = ipsegments
	c.ServeJSON()
}

// 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
// 2 每个包可以拥有多个init函数
// 3 包的每个源文件也可以拥有多个init函数
// 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
// 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
// 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
//读取iprole.txt文件，作为全局变量Iprolemaps，供调用访问者ip的权限用
var (
	Iprolemaps map[string]int
)

func init() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

func Createip() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

//取得访问者的权限
func Getiprole(ip string) (role int) {
	role, ok := Iprolemaps[ip]
	if ok {
		return role
	} else {
		return 5
	}
	//元素查找，这是通用的使用方法
	// v, ok := personDB["test1"]
	// if !ok {
	//     fmt.Println(" 没有找到信息")
	//     return
	// }
}

//获取下一个IP
func nextIp(ip string) string {
	ips := strings.Split(ip, ".")
	var i int
	for i = len(ips) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(ips[i])
		if n >= 255 {
			//进位
			ips[i] = "1"
		} else {
			//+1
			n++
			ips[i] = strconv.Itoa(n)
			break
		}
	}
	if i == -1 {
		//全部IP段都进行了进位,说明此IP本身已超出范围
		return ""
	}
	ip = ""
	leng := len(ips)
	for i := 0; i < leng; i++ {
		if i == leng-1 {
			ip += ips[i]
		} else {
			ip += ips[i] + "."
		}
	}
	return ip
}

//生成IP地址列表
func processIp(startIp, endIp string) []string {
	var ips = make([]string, 0)
	for ; startIp != endIp; startIp = nextIp(startIp) {
		if startIp != "" {
			ips = append(ips, startIp)
		}
	}
	ips = append(ips, startIp)
	return ips
}

//port代替权限role
func processFlag(arg []string) (maps map[string]int) {
	//开始IP,结束IP
	var startIp, endIp string
	//端口
	var ports []int = make([]int, 0)
	index := 0
	startIp = arg[index]
	si := net.ParseIP(startIp)
	if si == nil {
		//开始IP不合法
		// fmt.Println("'startIp' Setting error")
		beego.Error("开始IP不合法")
		return nil
	}
	index++
	endIp = arg[index]
	ei := net.ParseIP(endIp)
	if ei == nil {
		//未指定结束IP,即只扫描一个IP
		endIp = startIp
	} else {
		index++
	}

	tmpPort := arg[index]
	if strings.Index(tmpPort, "-") != -1 {
		//连续端口
		tmpPorts := strings.Split(tmpPort, "-")
		var startPort, endPort int
		var err error
		startPort, err = strconv.Atoi(tmpPorts[0])
		if err != nil || startPort < 1 || startPort > 65535 {
			//开始端口不合法
			return nil
		}
		if len(tmpPorts) >= 2 {
			//指定结束端口
			endPort, err = strconv.Atoi(tmpPorts[1])
			if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
				//结束端口不合法
				// fmt.Println("'endPort' Setting error")
				beego.Error("'endPort' Setting error")
				return nil
			}
		} else {
			//未指定结束端口
			endPort = 65535
		}
		for i := 0; startPort+i <= endPort; i++ {
			ports = append(ports, startPort+i)
		}
	} else {
		//一个或多个端口
		ps := strings.Split(tmpPort, ",")
		for i := 0; i < len(ps); i++ {
			p, err := strconv.Atoi(ps[i])
			if err != nil {
				//端口不合法
				// fmt.Println("'port' Setting error")
				beego.Error("'port' Setting error")
				return nil
			}
			ports = append(ports, p)
		}
	}

	//生成扫描地址列表
	ips := processIp(startIp, endIp)
	il := len(ips)
	m1 := make(map[string]int)
	for i := 0; i < il; i++ {
		pl := len(ports)
		for j := 0; j < pl; j++ {
			//			ipAddrs <- ips[i] + ":" + strconv.Itoa(ports[j])
			//			ipAddrs := ips[i] + ":" + strconv.Itoa(ports[j])
			m1[ips[i]] = ports[j]
		}
	}
	//	fmt.Print(slice1)
	return m1
	//	close(ipAddrs)
}

// @Title Get Category list
// @Description Get Category list by some info
// @Success 200 {object} models.GetAdminCategory
// @Param   id     path   string false       "category id"
// @router /category/:id [get]
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Category() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
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
	categories, err := models.GetAdminCategory(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// @Title Get Category by title
// @Description Get Category list by title info
// @Success 200 {object} models.GetAdminCategory
// @Param   title   query   string  false       "title of search"
// @router /categorytitle [get]
//根据名称title查询分级表
func (c *AdminController) CategoryTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminCategoryTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// @Title Post Category by pid title code grade
// @Description Get Category list by title info
// @Success 200 {object} models.AddAdminCategory
// @Param   pid   query   string  false       "parentid of category"
// @Param   title   query   string  false       "title of category"
// @Param   code   query   string  false       "code of category"
// @Param   grade   query   string  false       "grade of category"
// @router /category/addcategory [post]
//添加
func (c *AdminController) AddCategory() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
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
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminCategory(pidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdateCategory() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteCategory() {
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
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminCategory(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminCategory(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminCategory(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//********************日历开始**************
//添加日历
func (c *AdminController) AddCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	color := c.Input().Get("color")
	allday1 := c.Input().Get("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	public1 := c.Input().Get("public")
	var public bool
	if public1 == "true" {
		public = true
	} else {
		public = false
	}
	const lll = "2006-01-02 15:04"
	var starttime, endtime time.Time
	var err error
	if start != "" {
		starttime, err = time.Parse(lll, start)
		// beego.Info(start)
		// beego.Info(starttime)
		if err != nil {
			beego.Error(err)
		}
	} else {
		starttime = time.Now()
	}
	if end != "" {
		endtime, err = time.Parse(lll, end)
		if err != nil {
			beego.Error(err)
		}
	} else {
		endtime = starttime
	}
	_, err = models.AddAdminCalendar(title, content, color, allday, public, starttime, endtime)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
}

//返回日历json数据
//如果是管理员，则显示全部，非管理员，显示公开
func (c *AdminController) Calendar() {
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	const lll = "2006-01-02"
	startdate, err := time.Parse(lll, start)
	if err != nil {
		beego.Error(err)
	}
	enddate, err := time.Parse(lll, end)
	if err != nil {
		beego.Error(err)
	}
	var calendars []*models.AdminCalendar
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	if role == "1" {
		calendars, err = models.GetAdminCalendar(startdate, enddate, false)
		if err != nil {
			beego.Error(err)
		}
	} else {
		calendars, err = models.GetAdminCalendar(startdate, enddate, true)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//修改
func (c *AdminController) UpdateCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	cid := c.Input().Get("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	color := c.Input().Get("color")
	allday1 := c.Input().Get("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	public1 := c.Input().Get("public")
	var public bool
	if public1 == "true" {
		public = true
	} else {
		public = false
	}
	const lll = "2006-01-02 15:04"
	var starttime, endtime time.Time
	if start != "" {
		starttime, err = time.Parse(lll, start)
		// beego.Info(start)
		// beego.Info(starttime)
		if err != nil {
			beego.Error(err)
		}
	} else {
		starttime = time.Now()
	}
	if end != "" {
		endtime, err = time.Parse(lll, end)
		if err != nil {
			beego.Error(err)
		}
	} else {
		endtime = starttime
	}
	err = models.UpdateAdminCalendar(cidNum, title, content, color, allday, public, starttime, endtime)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
	// pid := c.Ctx.Input.Param(":id")
	//
	// title := c.Input().Get("title")
	// code := c.Input().Get("code")
	// grade := c.Input().Get("grade")
	// //pid转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["json"] = "ok"
	// 	c.ServeJSON()
	// }
}

//拖曳
func (c *AdminController) DropCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	delta := c.Input().Get("delta")
	daltaint, err := strconv.Atoi(delta)
	if err != nil {
		beego.Error(err)
	}
	calendar, err := models.GetAdminCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	t1 := calendar.Starttime.AddDate(0, 0, daltaint)
	t2 := calendar.Endtime.AddDate(0, 0, daltaint)
	err = models.DropAdminCalendar(idNum, t1, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//resize
func (c *AdminController) ResizeCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	delta := c.Input().Get("delta")
	delta = delta + "h"
	deltahour, err := time.ParseDuration(delta)
	if err != nil {
		beego.Error(err)
	}
	// starttime.Add(-time.Duration(hours) * time.Hour)
	calendar, err := models.GetAdminCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	// t1 := calendar.Starttime.Add(deltahour)
	t2 := calendar.Endtime.Add(deltahour)
	err = models.ResizeAdminCalendar(idNum, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	cid := c.Input().Get("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteAdminCalendar(cidNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

func (c *AdminController) SearchCalendar() {
	title := c.Input().Get("title")
	const lll = "2006-01-02"

	var calendars []*models.AdminCalendar
	var err error
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	if role == "1" {
		calendars, err = models.SearchAdminCalendar(title, false)
		if err != nil {
			beego.Error(err)
		}
	} else {
		calendars, err = models.SearchAdminCalendar(title, true)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
}

//******编辑项目同步ip**********
//根据项目id查询ip
func (c *AdminController) SynchIp() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	synchips, err := models.GetAdminSynchIp(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = synchips
	c.ServeJSON()
}

//添加
func (c *AdminController) AddsynchIp() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	username := c.Input().Get("username")
	ip := c.Input().Get("ip")
	port := c.Input().Get("port")
	//pid转成64为
	var pidNum int64
	var err error
	// if pid != "" {
	pidNum, err = strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// } else {
	// 	pidNum = 0
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	_, err = models.AddAdminSynchIp(pidNum, username, ip, port)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdatesynchIp() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	username := c.Input().Get("username")
	ip := c.Input().Get("ip")
	port := c.Input().Get("port")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	err = models.UpdateAdminSynchIp(cidNum, username, ip, port)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除
func (c *AdminController) DeletesynchIp() {
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
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}

		err = models.DeleteAdminSynchIp(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//******后台部门结构********
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Department() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminDepartment
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	categories, err := models.GetAdminDepart(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据名称title查询分级表
func (c *AdminController) DepartmentTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminDepartTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//添加
func (c *AdminController) AddDepartment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
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

	_, err = models.AddAdminDepart(pidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdateDepartment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateAdminDepart(cidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteDepartment() {
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
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminDepart(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//批量上传首页轮播图片
func (c *AdminController) AddCarousel() {
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		//获取上传的文件
		_, h, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
		}
		// var attachment string
		// var filesize int64
		if h != nil {
			//保存附件
			// attachment = h.Filename
			// beego.Info(attachment)
			path := "./attachment/carousel/" // + h.Filename
			url := "/attachment/carousel"    //+ h.Filename
			//存入成果数据库
			//如果编号重复，则不写入，值返回Id值。
			//根据id添加成果code, title, label, principal, content string, projectid int64
			_, err := models.AddAdminCarousel(h.Filename, url)
			if err != nil {
				beego.Error(err)
			} else {
				//存入文件夹
				err = c.SaveToFile("file", path+h.Filename) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
				}
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": h.Filename, "original": h.Filename, "url": url + "/" + h.Filename}
				c.ServeJSON()
			}
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//查询所有轮播图片
func (c *AdminController) Carousel() {
	carousels, err := models.GetAdminCarousel()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = carousels
	c.ServeJSON()
}

func (c *AdminController) Test1() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.tpl"
}
func (c *AdminController) Jsoneditor() {
	c.TplName = "jsoneditor.tpl"
}

func (c *AdminController) Testdown() {
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	filename := path.Base(filePath)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/download/"+filename)
}
