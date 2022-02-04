package controllers

import (
	// "encoding/json"
	m "github.com/3xxx/meritms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"regexp"
	"strconv"
	"strings"
	// "engineercms/controllers/validator"
	// "github.com/astaxie/beego/context"
	// "github.com/asofdate/hauth/core/groupcache"
	// "github.com/asofdate/hauth/core/hrpc"
	// "github.com/asofdate/hauth/core/models"
	// "github.com/asofdate/hauth/utils"
	// "github.com/asofdate/hauth/utils/hret"
	// "github.com/asofdate/hauth/utils/i18n"
	// "github.com/asofdate/hauth/utils/jwt"
	// "github.com/asofdate/hauth/utils/logs"
	// "github.com/asofdate/hauth/utils/validator"
)

type RoleController struct {
	beego.Controller
}

type Userrole struct {
	Id         int64
	Rolename   string `json:"name"`
	Rolenumber string
	Status     string `json:"role"`
	Level      string
}

type Tree struct {
	Id    int64  `json:"id"`
	Nodes []Tree `json:"nodes"`
}

// type CasbinRule struct {
// 	Id    int
// 	PType string
// 	V0    string
// 	V1    string
// 	V2    string
// 	V3    string
// 	V4    string
// 	V5    string
// }

// type RoleController struct {
// 	models models.RoleModel
// }

// var RoleCtl = &RoleController{
// 	models.RoleModel{},
// }
var e *casbin.Enforcer
var a *beegoormadapter.Adapter

//记得database换成meritms.db
func init() {
	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	a = beegoormadapter.NewAdapter("sqlite3", "database/meritms.db", true) // Your driver and data source.
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	// e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	e = casbin.NewEnforcer("conf/rbac_model.conf", a)
	// Load the policy from DB.
	e.LoadPolicy()
}

func (c *RoleController) Test() {
	// Check the permission.
	//请求的资源v1/v2/aaa.jpg
	//得到资源的扩展名Suffix-jpg，输入enforce中间那个(?i:pdf)不分大小写
	beego.Info(e.Enforce("alice", "/v1/v2/aaa.jpg", "write", "jpg"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.PDF", "delete", "PDF"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.jpg", "write", "jpg"))
	// beego.Info(e.Enforce("bob", "/v1/aaa.pdf", "read", "pdf"))
	// beego.Info(e.Enforce("bob", "/v1/v2/aaa.dwg", "read", "dwg"))
	// beego.Info(e.Enforce("bob", "/v1/v2/aaa.pdf", "write", "pdf"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.ttt", "read", "ttt")) //任意扩展名
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	// Save the policy back to DB.
	// e.SavePolicy()
}

// swagger:operation GET /v1/auth/role/page StaticFiles RoleController
//
// 角色管理页面
//
// 如果用户被授权访问角色管理页面,则系统返回角色管理页面内容,否则返回404错误
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
// func (RoleController) Page(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	rst, err := groupcache.GetStaticFile("AsofdateRolePage")
// 	if err != nil {
// 		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
// 		return
// 	}
// 	ctx.ResponseWriter.Write(rst)
// }

// swagger:operation GET /v1/auth/role/get RoleController RoleController
//
// 查询角色信息
//
// 查询指定域中的角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//这个作废，用下面的从casbin中获取用户角色
// func (c *RoleController) Get() {
// 	id := c.Ctx.Input.Param(":id")
// 	c.Data["Id"] = id
// 	c.Data["Ip"] = c.Ctx.Input.IP()
// 	// if id == "" { //如果id为空，则查询
// 	roles, err := m.GetRoles()
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	if id != "" {
// 		//pid转成64为
// 		idNum, err := strconv.ParseInt(id, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		//查出用户的角色，处于勾选状态
// 		userroles, err := m.GetRoleByUserId(idNum)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		userrole := make([]Userrole, 0)
// 		var level string
// 		level = "2"
// 		for _, v1 := range roles {
// 			for _, v2 := range userroles {
// 				if v2.RoleId == v1.Id {
// 					level = "1"
// 				}
// 			}
// 			aa := make([]Userrole, 1)
// 			aa[0].Id = v1.Id
// 			aa[0].Rolename = v1.Rolename
// 			aa[0].Rolenumber = v1.Rolenumber
// 			aa[0].Level = level
// 			userrole = append(userrole, aa...)
// 			aa = make([]Userrole, 0)
// 			level = "2"
// 		}
// 		c.Data["json"] = userrole
// 		c.ServeJSON()
// 	}
// 	c.Data["json"] = roles
// 	c.ServeJSON()
// }

func (c *RoleController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// if id == "" { //如果id为空，则查询
	roles, err := m.GetRoles()
	if err != nil {
		beego.Error(err)
	}
	//如果设置了role,用于onlyoffice的权限设置
	role := c.Input().Get("role")
	if role != "" {
		// roleint, err := strconv.Atoi(role)
		// if err != nil {
		// 	beego.Error(err)
		// }
		for _, v := range roles {
			v.Status = role
		}
	}

	if id != "" { //如果选中了用户，则显示用户所具有的角色
		//pid转成64为
		// idNum, err := strconv.ParseInt(id, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		//查出用户的角色，处于勾选状态，来自casbin\rbac_api.go
		userroles, err := e.GetRolesForUser(id)
		if err != nil {
			beego.Error(err)
		}
		userrole := make([]Userrole, 0)
		var level string
		level = "2"
		for _, v1 := range roles {
			for _, v2 := range userroles {
				ridNum, err := strconv.ParseInt(strings.Replace(v2, "role_", "", -1), 10, 64)
				if err != nil {
					beego.Error(err)
				}
				if ridNum == v1.Id {
					level = "1" //if (row.Level === "1") checked: true
				}
			}
			aa := make([]Userrole, 1)
			aa[0].Id = v1.Id
			aa[0].Rolename = v1.Rolename
			aa[0].Rolenumber = v1.Rolenumber
			aa[0].Level = level
			aa[0].Status = v1.Status
			userrole = append(userrole, aa...)
			aa = make([]Userrole, 0)
			level = "2"
		}
		c.Data["json"] = userrole //用户所具有的角色，勾选
		c.ServeJSON()
	} else {
		c.Data["json"] = roles //角色列表
		c.ServeJSON()
	}
}

// swagger:operation POST /v1/auth/role/post RoleController RoleController
//
// 新增角色信息
//
// 在某个指定的域中,新增角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success

//添加角色
func (c *RoleController) Post() {
	// u := m.Role{}
	// if err := c.ParseForm(&u); err != nil {
	// 	beego.Error(err.Error)
	// 	return
	// }
	var role m.Role
	role.Rolename = c.Input().Get("rolename")
	role.Rolenumber = c.Input().Get("rolenumber")
	// statusint, err := strconv.Atoi(c.Input().Get("status"))
	// if err != nil {
	// 	beego.Error(err)
	// }
	role.Status = c.Input().Get("status")
	// role.Createtime = time.Now()
	id, err := m.SaveRole(role)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err)
		c.Data["json"] = "wrong"
		c.ServeJSON()
		// return
	}
}

//向userid里添加权限——这个作废，用casbin的
// func (c *RoleController) UserRole() {
// 	uid := c.GetString("uid") //secofficeid
// 	//id转成64位
// 	uidNum, err := strconv.ParseInt(uid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//取出所有uidnum的role
// 	userroles, err := m.GetRoleByUserId(uidNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	ids := c.GetString("ids") //roleid
// 	// beego.Info(ids)
// 	array := strings.Split(ids, ",")
// 	// beego.Info(array)
// 	bool := false
// 	for _, v1 := range array {
// 		// pid = strconv.FormatInt(v1, 10)
// 		//id转成64位
// 		idNum, err := strconv.ParseInt(v1, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		for _, v2 := range userroles {
// 			//没有找到则插入
// 			if v2.RoleId == idNum {
// 				bool = true
// 			}
// 		}
// 		if bool == false {
// 			//存入数据库
// 			err = m.AddUserRole(uidNum, idNum)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			// beego.Info(uidNum)
// 			// beego.Info(idNum)
// 		}
// 		bool = false
// 	}

// 	for _, v3 := range userroles {
// 		for _, v4 := range array {
// 			//id转成64位
// 			idNum, err := strconv.ParseInt(v4, 10, 64)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			//没有找到则删除
// 			if v3.RoleId == idNum {
// 				bool = true
// 			}
// 		}

// 		if bool == false {
// 			//删除数据库
// 			err = m.DeleteUserRole(uidNum, v3.RoleId)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 		}
// 		bool = false
// 	}
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		c.Data["json"] = "ok"
// 		c.ServeJSON()
// 	}
// }

//AddPolicy(sec string, ptype string, rule []string)
//添加用户角色
//先删除用户所有角色
func (c *RoleController) UserRole() {
	//要支持批量分配角色，循环用户id
	uid := c.GetString("uid") //secofficeid
	//先删除用户的权限
	e.DeleteRolesForUser(uid) //数据库没有删掉！
	//删除数据库中角色中的用户
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err := qs.Filter("PType", "g").Filter("v0", uid).Delete()
	// if err != nil {
	// 	beego.Error(err)
	// }
	//再添加，如果没有选择，相当于删除了全部角色
	ids := c.GetString("ids") //roleid
	if ids != "" {
		array := strings.Split(ids, ",")
		// var rule []string
		for _, v1 := range array {
			// rule = append(rule, uid, v1)
			// e.AddPolicy(uid, v1)
			e.AddGroupingPolicy(uid, "role_"+v1) //management_api.go
			//应该用AddRoleForUser()//rbac_api.go
			// rule = make([]string, 0)
		}
		// a.SavePolicy(e.GetModel())//autosave默认是true
		// 	[{0 p 12 1    } {0 g 8 1    } {0 g 7 1
		//    } {0 g 7 2    } {0 g 5 1    } {0 g 5 2    }]
		// lines := [7][4]string{{"0", "p", "100", "1"}, {"0", "p", "101", "1"}}
		// _, err := a.o.InsertMulti(len(lines), lines)
		// return err
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//给角色赋查看科室成果、科室价值的权限——根据权限显示侧栏菜单
//先删除角色对于这个组织架构的所有权限
func (c *RoleController) RoleAchieve() {
	var success bool
	var nodeidint int

	var err error
	roleids := c.GetString("roleids")
	rolearray := strings.Split(roleids, ",")

	treeids := c.GetString("treeids") //项目目录id，25001,25002
	treearray := strings.Split(treeids, ",")
	// beego.Info(treearray)
	treenodeids := c.GetString("treenodeids") //项目目录的nodeid 0.0.0-0.0.1-0.1.0-0.1.0
	treenodearray := strings.Split(treenodeids, ",")
	// beego.Info(treenodearray)
	//取出项目目录的顶级
	var nodesid, nodesids []string
	if len(treenodearray) > 1 {
		nodesids, err = highest(treenodearray, nodesid, 0)
		if err != nil {
			beego.Error(err)
		}
	} else {
		nodesids = []string{"0"} //append(nodesids, "0")
	}
	// beego.Info(nodesids)

	//删除这些角色的全部查看成果权限
	for _, v1 := range rolearray { //其实只有一个值
		// o := orm.NewOrm()
		// qs := o.QueryTable("casbin_rule")
		// _, err := qs.Filter("PType", "p").Filter("v0", "role_"+v1).Filter("v1__contains", "secoffice_").Delete()
		// if err != nil {
		// 	beego.Error(err)
		// }
		e.RemoveFilteredPolicy(0, "role_"+v1, "secoffice")
	}
	// e.LoadPolicy() //重载权限
	if treeids != "" {
		for _, v1 := range rolearray {
			for _, v3 := range nodesids {
				// beego.Info(v3)
				nodeidint, err = strconv.Atoi(v3)
				if err != nil {
					beego.Error(err)
				}
				//id转成64位
				// pidNum, err := strconv.ParseInt(treearray[nodeidint], 10, 64)
				// if err != nil {
				// 	beego.Error(err)
				// }
				//根据projid取出路径
				// proj, err := m.GetProj(pidNum)
				// if err != nil {
				// 	beego.Error(err)
				// }
				// success = e.AddPolicy("role_"+v1, "secoffice_"+treearray[nodeidint]) //来自casbin\management_api.go
				//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
				success = e.AddPermissionForUser("role_"+v1, "secoffice", "secoffice_"+treearray[nodeidint], "sec")
			}
		}
	} else {
		success = true
	}
	// e.LoadPolicy() //重载权限

	if success == true {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "wrong"
	}
	c.ServeJSON()
}

//迭代查出最高级的树状目录
//nodesid是数组的序号
//nodeid是节点号：0.0.1   0.0.1.0
func highest(nodeid, nodesid []string, i int) (nodesid1 []string, err error) {
	if i == 0 {
		nodesid = append(nodesid, "0")
	}
	var i1 int
	for i1 = i; i1 < len(nodeid)-1; i1++ {
		matched, err := regexp.MatchString("(?i:"+nodeid[i]+")", nodeid[i1+1])
		// fmt.Println(matched)
		if err != nil {
			beego.Error(err)
		}
		if !matched {
			i = i1 + 1
			nodesid = append(nodesid, strconv.Itoa(i1+1))
			break
		} else {
			if i == len(nodeid)-2 {
				return nodesid, err
			}
		}
	}
	if i1 < len(nodeid)-1 {
		nodesid, err = highest(nodeid, nodesid, i)
	}
	return nodesid, err
}

//查询角色所具有的权限对应的项目目录
func (c *RoleController) GetRolePermission() {
	roleid := c.GetString("roleid") //角色id
	action := c.GetString("action")
	projectid := c.GetString("projectid")
	sufids := c.GetString("sufids") //扩展名
	// beego.Info(sufids)
	var suf string
	switch sufids {
	case "任意":
		suf = ".*"
	case "":
		suf = "(?i:PDF)"
	case "PDF":
		suf = "(?i:PDF)"
	}
	// beego.Info(suf)
	// beego.Info(roleid)
	// beego.Info(action)
	// beego.Info(projectid)

	// myRes := e.GetPermissionsForUser(roleid)
	// beego.Info(myRes)

	// 	2018/01/03 21:42:15 [I] [roleControllers.go:543] [[1 /25001/* POST .*] [1 /25001
	// /* PUT .*] [1 /25001/* DELETE .*] [1 /25001/* GET .*] [1 /25001/25003/* GET (?i:
	// PDF)] [1 /25001/25002/25013/* GET (?i:PDF)] [1 /25001/25002/25012/* GET (?i:PDF)
	// ] [1 /25001/25002/25011/* GET (?i:PDF)] [1 /25001/* GET (?i:PDF)] [1 /25001/2500
	// 4/* POST .*]]
	// Permissions, err := models.GetPermissions(roleid,projectid,action)
	// if err != nil {
	// 	beego.Error(err)
	// }
	var paths []beegoormadapter.CasbinRule
	o := orm.NewOrm()
	qs := o.QueryTable("casbin_rule")
	if action == "GET" || action == "" {
		_, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1__contains", "/"+projectid+"/").Filter("v2", "GET").Filter("v3", suf).All(&paths)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(paths)
	} else {
		_, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1__contains", "/"+projectid+"/").Filter("v2", action).All(&paths)
		if err != nil {
			beego.Error(err)
		}
	}
	// beego.Info(paths)
	var projids []string
	for _, v1 := range paths {
		projid := strings.Replace(v1.V1, "/*", "", -1)
		projids = append(projids, path.Base(projid))
	}
	// beego.Info(projids)
	c.Data["json"] = projids
	c.ServeJSON()
}

//查询角色所具有的查询科室成果、科室价值的权限
func (c *RoleController) GetRoleAchieve() {
	roleid := c.GetString("roleid") //角色id
	// beego.Info(roleid)
	// var paths []beegoormadapter.CasbinRule
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1__contains", "secoffice_").All(&paths)
	// if err != nil {
	// 	beego.Error(err)
	// }

	//GetPermissionsForUser只能过滤一个字段，所以不行
	paths := e.GetFilteredPolicy(0, "role_"+roleid, "secoffice")
	var secids []string
	for _, v1 := range paths {
		secid := strings.Replace(v1[2], "secoffice_", "", -1)
		secids = append(secids, secid)
	}
	// beego.Info(secids)
	c.Data["json"] = secids
	c.ServeJSON()
}

// swagger:operation POST /v1/auth/role/delete RoleController RoleController
//
// 删除角色信息
//
// 删除某个指定域中的角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (c *RoleController) Delete() {
	roleids := c.GetString("ids")
	rolearray := strings.Split(roleids, ",")
	for _, v1 := range rolearray {
		// rid, _ := c.GetInt64("roleid")
		//id转成64位
		idNum, err := strconv.ParseInt(v1, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		_, err = m.DeleteRole(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

// swagger:operation PUT /v1/auth/role/put RoleController RoleController
//
// 更新角色信息
//
// 更新某个域中的角色信息,角色编码不能更新
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (c *RoleController) Update() {
	var role m.Role
	roleid := c.Input().Get("roleid")
	idNum, err := strconv.ParseInt(roleid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	role.Id = idNum
	role.Rolename = c.Input().Get("rolename")
	role.Rolenumber = c.Input().Get("rolenumber")

	// statusint, err := strconv.Atoi(c.Input().Get("status"))
	// if err != nil {
	// 	beego.Error(err)
	// }
	role.Status = c.Input().Get("status")

	err = m.UpdateRole(role)
	if err == nil {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err)
		// return
	}
}
