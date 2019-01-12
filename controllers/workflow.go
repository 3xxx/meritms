package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	// "strings"
	// "testing"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/js-ojus/flow"
	// "github.com/3xxx/meritms/models"
	// _ "github.com/mattn/go-sqlite3"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	"log"
	"strconv"
	// "time"
)

// Flowtest API
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// var db *sql.DB

// func init() {
// 	driver, connStr := "mysql", "root:root@/flow"
// 	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	// flow.RegisterDB(tdb)

// 	if tdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db = tdb
// }

// @Title show wf page
// @Description show workflow page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /workflow [get]
// 页面
func (c *MainController) WorkFlow() {
	c.TplName = "index.tpl"
}

// @Title post wf doctype...
// @Description post workflowdoctype..
// @Param name query string  true "The name of doctype"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtype [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowType() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义流程类型
	_, err := flow.DocTypes.New(tx, name) //"图纸设计流程"
	if err != nil {
		fmt.Println(err)
	}
	// dtID2, err := flow.DocTypes.New(tx, "合同评审流程")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// beego.Info(dtID2)
	// dtID3, err := flow.DocTypes.New(tx, "变更立项流程")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// beego.Info(dtID3)

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf docstate...
// @Description post workflowdocstate..
// @Param name query string  true "The name of docstate"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowstate [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowState() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义流程状态
	_, err := flow.DocStates.New(tx, name) //"设计中..."
	if err != nil {
		fmt.Println(err)
	}
	// dsID2, err := flow.DocStates.New(tx, "校核中...")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dsID3, err := flow.DocStates.New(tx, "审查中...")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// flow.DocStates.New(tx, "批准中...")
	// flow.DocStates.New(tx, "申报中...")
	// flow.DocStates.New(tx, "评估中...")
	// flow.DocStates.New(tx, "审批中...")

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf docaction...
// @Description post workflowdocaction..
// @Param name query string  true "The name of docaction"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowaction [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowAction() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义流程动作类型
	_, err := flow.DocActions.New(tx, name, false) //"设计完成后提交"改变状态设计中...为校核中...
	if err != nil {
		fmt.Println(err)
	}
	// daID2, err := flow.DocActions.New(tx, "校核完成后提交", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// daID3, err := flow.DocActions.New(tx, "审查完成后提交", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// daID4, err := flow.DocActions.New(tx, "核定完成后提交", true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// daID5, err := flow.DocActions.New(tx, "编制完成后提交", true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// daID6, err := flow.DocActions.New(tx, "审批完成后提交", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// daID7, err := flow.DocActions.New(tx, "立项完成后提交", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //添加流程规则1:oldstate1 action1 newstate2
	// err = flow.DocTypes.AddTransition(tx, dtID1, dsID1, daID1, dsID2)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// //添加流程规则2:oldstate2 action2 newstate3
	// err = flow.DocTypes.AddTransition(tx, dtID1, dsID2, daID2, dsID3)
	// if err != nil {
	// 	beego.Error(err)
	// }

	// //定义流程类型doctype下的唯一流程workflow
	// workflowID1, err := flow.Workflows.New(tx, "图纸设计-三级校审流程", dtID1, dsID1) //初始状态是“设计中...”——校核——审查——完成
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// beego.Info(workflowID1)
	// // workflowID2, err := flow.Workflows.New(tx, "图纸设计-二级校审流程", dtID1, dsID1) //初始状态是“设计中...”-“校核”——完成
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // beego.Info(workflowID2)
	// //定义合同评审下的流程类型：部门合同流程，总院合同流程
	// //略

	// //定义用户、组、角色、权限集合
	// accessContextID1, err := flow.AccessContexts.New(tx, "Context")
	// if err != nil {
	// 	beego.Error(err)
	// }

	// //定义流程类型workflow下的具体每个节点node，用户对文件执行某个动作（event里的action）后，会沿着这些节点走
	// // AddNode maps the given document state to the specified node.  This
	// // map is consulted by the workflow when performing a state transition
	// // of the system.nodeID1
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID1, accessContextID1, workflowID1, "图纸设计-三级校审流程-设计", flow.NodeTypeBegin)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID2, accessContextID1, workflowID1, "图纸设计-三级校审流程-校核", flow.NodeTypeLinear)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID3, accessContextID1, workflowID1, "图纸设计-三级校审流程-审查", flow.NodeTypeEnd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //定义用户-组-角色-权限关系
	// res, err := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-1', 'email1@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ := res.LastInsertId()
	// uID1 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID1)

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-2', 'email2@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID2 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID2)

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-3', 'email3@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID3 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID3)

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-4', 'email4@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID4 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID4)

	// gID1 := fatal1(flow.Groups.New(tx, "设计人员组", "G")).(flow.GroupID)
	// gID2 := fatal1(flow.Groups.New(tx, "校核人员组", "G")).(flow.GroupID)
	// fatal0(flow.Groups.AddUser(tx, gID1, uID1))
	// fatal0(flow.Groups.AddUser(tx, gID1, uID2))
	// fatal0(flow.Groups.AddUser(tx, gID1, uID3))

	// fatal0(flow.Groups.AddUser(tx, gID2, uID2))
	// fatal0(flow.Groups.AddUser(tx, gID2, uID3))
	// fatal0(flow.Groups.AddUser(tx, gID2, uID4))
	// roleID1 := fatal1(flow.Roles.New(tx, "设计人员角色")).(flow.RoleID)
	// roleID2 := fatal1(flow.Roles.New(tx, "校核人员角色")).(flow.RoleID)
	// //给角色role赋予action权限
	// fatal0(flow.Roles.AddPermissions(tx, roleID1, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4}))
	// fatal0(flow.Roles.AddPermissions(tx, roleID2, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4, daID5, daID6, daID7}))
	// //给用户组group赋予角色role
	// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID1, roleID1)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// //将group和role加到accesscontext里——暂时不理解
	// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID2, roleID2)
	// if err != nil {
	// 	beego.Error(err) //UNIQUE constraint failed: wf_ac_group_roles.ac_id已修补
	// }

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf transition...
// @Description post transition..
// @Param name query string  true "The name of transition"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtransition [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowTransition() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	dtID := c.Input().Get("dtID")
	//pid转成64为
	dtID1, err := strconv.ParseInt(dtID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsID1str := c.Input().Get("dsID1")
	dsID1, err := strconv.ParseInt(dsID1str, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	daID := c.Input().Get("daID")
	daID1, err := strconv.ParseInt(daID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsID2str := c.Input().Get("dsID2")
	dsID2, err := strconv.ParseInt(dsID2str, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//添加流程规则1:oldstate1 action1 newstate2
	err = flow.DocTypes.AddTransition(tx, flow.DocTypeID(dtID1), flow.DocStateID(dsID1), flow.DocActionID(daID1), flow.DocStateID(dsID2))
	if err != nil {
		beego.Error(err)
	}
	//添加流程规则2:oldstate2 action2 newstate3
	// err = flow.DocTypes.AddTransition(tx, flow.DocTypeID(dtID1), flow.DocStateID(dsID2), flow.DocActionID(daID2), flow.DocStateID(dsID3))
	// if err != nil {
	// 	beego.Error(err)
	// }

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf Workflow...
// @Description post Workflow..
// @Param name query string  true "The name of Workflow"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowworkflow [post]
// 管理员定义流程Workflow
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowWorkflow() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()
	dtID := c.Input().Get("dtID")
	//pid转成64为
	dtID1, err := strconv.ParseInt(dtID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsID1str := c.Input().Get("dsID1")
	dsID1, err := strconv.ParseInt(dsID1str, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//定义流程类型doctype下的唯一流程workflow
	workflowID1, err := flow.Workflows.New(tx, "图纸设计-三级校审流程", flow.DocTypeID(dtID1), flow.DocStateID(dsID1)) //初始状态是“设计中...”——校核——审查——完成
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(workflowID1)
	// workflowID2, err := flow.Workflows.New(tx, "图纸设计-二级校审流程", dtID1, dsID1) //初始状态是“设计中...”-“校核”——完成
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// beego.Info(workflowID2)
	//定义合同评审下的流程类型：部门合同流程，总院合同流程
	//略

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf AccessContext...
// @Description post AccessContext..
// @Param name query string  true "The name of AccessContext"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowaccesscontext [post]
// 管理员定义流程AccessContext
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowAccessContext() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义用户、组、角色、权限集合
	_, err := flow.AccessContexts.New(tx, name) //"Context"
	if err != nil {
		beego.Error(err)
	}

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf Node...
// @Description post Node..
// @Param name query string  true "The name of Node"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flownode [post]
// 管理员定义流程Node
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowNode() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	dtID := c.Input().Get("dtID")
	//pid转成64为
	dtID1, err := strconv.ParseInt(dtID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsID1str := c.Input().Get("dsID1")
	dsID1, err := strconv.ParseInt(dsID1str, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	acID1str := c.Input().Get("acID1")
	acID1, err := strconv.ParseInt(acID1str, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据doctypeid获得workflow
	workflow, err := flow.Workflows.GetByDocType(flow.DocTypeID(dtID1))
	//定义流程类型workflow下的具体每个节点node，用户对文件执行某个动作（event里的action）后，会沿着这些节点走
	// AddNode maps the given document state to the specified node.  This
	// map is consulted by the workflow when performing a state transition
	// of the system.nodeID1
	_, err = flow.Workflows.AddNode(tx, flow.DocTypeID(dtID1), flow.DocStateID(dsID1), flow.AccessContextID(acID1), workflow.ID, name, flow.NodeTypeBegin)
	if err != nil {
		fmt.Println(err)
	}
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID2, accessContextID1, workflowID1, "图纸设计-三级校审流程-校核", flow.NodeTypeLinear)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID3, accessContextID1, workflowID1, "图纸设计-三级校审流程-审查", flow.NodeTypeEnd)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf user...
// @Description post user..
// @Param name query string  true "The name of user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowuser [post]
// 管理员定义流程user
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowUser() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	//定义用户-组-角色-权限关系
	res, err := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('秦', '晓川-1', 'email1@example.com', 1)`)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	uid, _ := res.LastInsertId()
	uID1 := flow.UserID(uid)
	_, err = flow.Groups.NewSingleton(tx, uID1)

	res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('秦', '晓川-2', 'email2@example.com', 1)`)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	uid, _ = res.LastInsertId()
	uID2 := flow.UserID(uid)
	_, err = flow.Groups.NewSingleton(tx, uID2)

	res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('秦', '晓川-3', 'email3@example.com', 1)`)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	uid, _ = res.LastInsertId()
	uID3 := flow.UserID(uid)
	_, err = flow.Groups.NewSingleton(tx, uID3)

	res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('秦', '晓川-4', 'email4@example.com', 1)`)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	uid, _ = res.LastInsertId()
	uID4 := flow.UserID(uid)
	_, err = flow.Groups.NewSingleton(tx, uID4)

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf Group...
// @Description post Group..
// @Param name query string  true "The name of Group"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgroup [post]
// 管理员定义流程Group
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowGroup() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	_, err := flow.Groups.New(tx, "设计人员组", "G") //).(flow.GroupID)
	if err != nil {
		beego.Error(err)
	}
	// gID2 := fatal1(flow.Groups.New(tx, "校核人员组", "G")).(flow.GroupID)

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf GroupUser...
// @Description post Group..
// @Param name query string  true "The name of GroupUser"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgroupuser [post]
// 管理员定义流程GroupUser
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowGroupUser() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	gID := c.Input().Get("gID")
	//pid转成64为
	gID1, err := strconv.ParseInt(gID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	uID := c.Input().Get("uID")
	//pid转成64为
	uID1, err := strconv.ParseInt(uID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = flow.Groups.AddUser(tx, flow.GroupID(gID1), flow.UserID(uID1))
	if err != nil {
		beego.Error(err)
	}
	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf Role...
// @Description post Role..
// @Param name query string  true "The name of Role"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowrole [post]
// 管理员定义流程Role
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowRole() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	_, err := flow.Roles.New(tx, "设计人员角色")
	if err != nil {
		beego.Error(err)
	}
	// roleID2 := fatal1(flow.Roles.New(tx, "校核人员角色")).(flow.RoleID)

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf Permission...
// @Description post Permission..
// @Param name query string  true "The name of Permission"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowpermission [post]
// 管理员定义流程Permission
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowPermission() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	roleID := c.Input().Get("roleID1")
	//pid转成64为
	roleID1, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dtID := c.Input().Get("dtID")
	//pid转成64为
	dtID1, err := strconv.ParseInt(dtID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据用户选择的动作
	var actions []flow.DocActionID //[]flow.DocActionID{daID1, daID2, daID3, daID4}
	//给角色role赋予action权限
	err = flow.Roles.AddPermissions(tx, flow.RoleID(roleID1), flow.DocTypeID(dtID1), actions)

	if err != nil {
		beego.Error(err)
	}
	// fatal0(flow.Roles.AddPermissions(tx, roleID2, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4, daID5, daID6, daID7}))

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf GroupRole...
// @Description post GroupRole..
// @Param name query string  true "The name of GroupRole"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgrouprole [post]
// 管理员定义流程GroupRole
// 流程动作action、流程流向transition、流程事件event
func (c *MainController) FlowGroupRole() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	acID := c.Input().Get("acID1")
	//pid转成64为
	acID1, err := strconv.ParseInt(acID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gID := c.Input().Get("gID1")
	//pid转成64为
	gID1, err := strconv.ParseInt(gID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	roleID := c.Input().Get("roleID1")
	//pid转成64为
	roleID1, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//给用户组group赋予角色role
	err = flow.AccessContexts.AddGroupRole(tx, flow.AccessContextID(acID1), flow.GroupID(gID1), flow.RoleID(roleID1))
	if err != nil {
		beego.Error(err)
	}
	//将group和role加到accesscontext里——暂时不理解
	// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID2, roleID2)
	// if err != nil {
	// 	beego.Error(err) //UNIQUE constraint failed: wf_ac_group_roles.ac_id已修补
	// }

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf state
// @Description post workflow state
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdocevent [get]
// 每次新建一个文件，自动对文件进行流程初始化，即，进行定义动作事件
func (c *MainController) FlowDocEvent() {
	//连接数据库
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}

	//查询预先定义的doctype流程类型
	dtID1, err := flow.DocTypes.GetByName("图纸设计")
	if err != nil {
		beego.Error(err)
	}
	beego.Info(dtID1)
	//查询预先定义的docstate状态1
	dsID1, err := flow.DocStates.GetByName("设计中...")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(dsID1)
	//查询预先定义的docstate状态2
	dsID2, err := flow.DocStates.GetByName("校核中...")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(dsID2)
	//查询预先定义的docstate状态3
	dsID3, err := flow.DocStates.GetByName("审查中...")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(dsID3)
	//查询预先定义的action动作1
	daID1, err := flow.DocActions.GetByName("提交设计")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(daID1)
	//查询预先定义的action动作2
	daID2, err := flow.DocActions.GetByName("校核") //应该叫"提交校核"
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(daID2)
	//查询预先定义的action动作3
	daID3, err := flow.DocActions.GetByName("审查") //应该叫"提交审查"
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(daID3)
	//查询预先定义的流程类型workflow，这个相当于doctype下面再分很多种流程
	//比如doctype为图纸设计流程，下面可以分为二级校审流程，三级校审流程，四级校审流程
	myWorkflow, err := flow.Workflows.GetByName("图纸设计-三级校审流程")
	if err != nil {
		beego.Error(err)
	}
	beego.Info(myWorkflow)
	//查询context——这个应该是管理用户-组-权限的
	accessContextID1, err := flow.AccessContexts.List("Context", 0, 0)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(accessContextID1[0].ID)

	beego.Info(flow.GroupID(1))
	//开始为具体一个文件设立流程-此处是新建一个文件。对于旧文件应该怎么操作来着？
	//document根据doctype取得唯一workflow的state作为document的state
	docNewInput := flow.DocumentsNewInput{
		DocTypeID:       dtID1.ID,                //属于图纸设计类型的流程
		AccessContextID: accessContextID1[0].ID,  //所有用户权限符合这个contex的要求
		GroupID:         11,                      //groupId,初始状态下的用户组，必须是个人用户组（一个用户也可以成为一个独特的组，因为用户无法赋予角色，所以必须将用户放到组里）
		Title:           "厂房布置图",                 //这个文件的名称
		Data:            "设计、制图: 秦晓川1, 校核: 秦晓川2", //文件的描述
	}
	// flow.Documents.New(tx, &docNewInput)
	DocumentID1, err := flow.Documents.New(tx, &docNewInput)
	if err != nil {
		beego.Error(err)
	}
	// tx.Commit() //new后面一定要跟commit
	beego.Info(DocumentID1)

	beego.Info(daID2)
	beego.Info(flow.GroupID(12))
	//建立好document，循环建立events，根据哪个来建立？
	//根据document的Doctypes.Transitions获取state和action
	//循环建立events，然后展示给客户端
	//用户点开这个文件，根据文件的状态，list出所有这个状态的events，比如前进，后退等
	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   dtID1.ID, //flow.DocTypeID(1),
		DocumentID:  DocumentID1,
		DocStateID:  dsID1.ID, //document state must be this state，文档的现状状态
		DocActionID: daID2.ID, //flow.DocActionID(2),
		GroupID:     12,       //必须是个人用户组
		Text:        "校核",
	}

	docEventID1, err := flow.DocEvents.New(tx, &docEventInput)
	if err != nil {
		beego.Error(err)
	}
	tx.Commit() //一个函数里只能有一个commit！
	beego.Info(docEventID1)
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// @Title post wf doclist
// @Description post workflow doclist
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdoclist [get]
// 文件列表页，水平显示每个文件的状态
func (c *MainController) FlowDocList() {

	documentslistinput := flow.DocumentsListInput{
		DocTypeID:       1, // Documents of this type are listed; required
		AccessContextID: 1, // Access context from within which to list; required
		GroupID:         1, // 本人所在的组List documents created by this (singleton) group
		DocStateID:      1, // 忽略List documents currently in this state
		//CtimeStarting:   time.Now(), // List documents created after this time
		//CtimeBefore:     time.Now(), // List documents created before this time
		//TitleContains:   string,     // List documents whose title contains the given text; expensive operation
		//RootOnly:        bool,       // List only root (top-level) documents
	}
	var offset, limit int64
	offset = 0
	limit = 0
	_, err := flow.Documents.List(&documentslistinput, offset, limit)
	if err != nil {
		beego.Error(err)
	}

}

// @Title post wf state
// @Description post workflow state
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdocshow [get]
// 显示具体文件的操作按钮-events：
func (c *MainController) FlowDocShow() {
	var tx *sql.Tx
	dtID := c.Input().Get("dtID1")
	//pid转成64为
	dtID1, err := strconv.ParseInt(dtID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	docID := c.Input().Get("docid")
	//pid转成64为
	docID1, err := strconv.ParseInt(docID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据客户端document的id取得document
	DocumentID1, err := flow.Documents.Get(tx, flow.DocTypeID(dtID1), flow.DocumentID(docID1))
	if err != nil {
		beego.Error(err)
	}
	//用户点开这个文件，根据文件的状态，list出所有这个状态的events，比如前进，后退等
	//doctypeid从哪来？所有操作都带doctype吧，因为当前走的流程都属于这个doctype下的
	docEventListInput := flow.DocEventsListInput{
		DocTypeID:       flow.DocTypeID(dtID1), // Events on documents of this type are listed
		AccessContextID: DocumentID1.AccCtx.ID, // Access context from within which to list
		GroupID:         DocumentID1.Group.ID,  // List events created by this (singleton) group
		DocStateID:      DocumentID1.State.ID,  // List events acting on this state
		// CtimeStarting:   time.Time,             // List events created after this time
		// CtimeBefore:     time.Time,             // List events created before this time
		Status: flow.EventStatusAll, // EventStatusAll,List events that are in this state of application
	}
	var offset, limit int64
	offset = 0
	limit = 0
	myDocEvent, err := flow.DocEvents.List(&docEventListInput, offset, limit)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(myDocEvent)
}

// @Title post wf next
// @Description post workflow next
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flownext [post]
// 用户点击提交，前进——或，回退 等
func (c *MainController) FlowNext() {
	//连接数据库var db *sql.DB
	// driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	// if tdb == nil {
	// 	log.Fatal("given database handle is `nil`")
	// }
	// db := tdb
	// tx, err := db.Begin()
	// if err != nil {
	// 	beego.Error(err)
	// }
	var tx *sql.Tx
	//客户界面设计上：设计中，然后点击按钮提交，这个提交动作怎么自动赋予的？
	//是系统根据文件的状态列出所有可能的events
	//根据documentid，查出document，
	//根据doctypeid，查出workflow

	//给出接受的组groupids
	//用户选择接受者，或系统自动根据action权限来选择组
	// groupIds := []flow.GroupID{flow.GroupID(13)}
	// beego.Info(groupIds)
	//查询workflow
	// myWorkflow, err := flow.Workflows.GetByName("图纸设计-三级校审流程")
	// if err != nil {
	// 	beego.Error(err)
	// }
	wfID := c.Input().Get("wfID1")
	//pid转成64为
	wfID1, err := strconv.ParseInt(wfID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	deID := c.Input().Get("deID1")
	deID1, err := strconv.ParseInt(deID, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// gID := c.Input().Get("gID1")
	// gID1, err := strconv.ParseInt(gID, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	myWorkflow, err := flow.Workflows.Get(flow.WorkflowID(wfID1))
	if err != nil {
		beego.Error(err)
	}
	beego.Info(myWorkflow)

	//针对具体一个文件定义动作事件，从"校核中……"状态通过动作"校核"将它修改为"审查中……"
	myDocEvent, err := flow.DocEvents.Get(flow.DocEventID(deID1))
	if err != nil {
		beego.Error(err)
	}
	var groupIds []flow.GroupID
	newDocStateId, err := myWorkflow.ApplyEvent(tx, myDocEvent, groupIds)
	if err != nil {
		beego.Error(err)
	}
	tx.Commit() //一个函数里只能有一个commit
	fmt.Println("newDocStateId=", newDocStateId, err)

	c.Data["json"] = "OK"
	c.ServeJSON()
}

// beego.Info(wflist)

// wflist, err = flow.DocStates.List(0, 0)
// if err != nil {
// 	beego.Error(err)
// }
// // beego.Info(wflist1)
// wflist, err = DocActions.List(0, 0)
// if err != nil {
// 	beego.Error(err)
// }
// wflist1, err = flow.Workflows.List(0, 0)
// if err != nil {
// 	beego.Error(err)
// }
// fatal1 expects a value and an error value as its arguments.
func fatal1(val1 interface{}, err error) interface{} {
	if err != nil {
		fmt.Println("%v", err)
	}
	return val1
}

// error0 expects only an error value as its argument.
func error0(err error) error {
	if err != nil {
		fmt.Println("%v", err)
	}
	return err
}

// error1 expects a value and an error value as its arguments.
func error1(val1 interface{}, err error) interface{} {
	if err != nil {
		fmt.Println("%v", err)
		return nil
	}
	return val1
}

// fatal0 expects only an error value as its argument.
func fatal0(err error) {
	if err != nil {
		fmt.Println("%v", err)
	}
}
