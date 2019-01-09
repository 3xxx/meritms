package controllers

import (
	"database/sql"
	// "strings"
	// "testing"
	"fmt"
	"github.com/astaxie/beego"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/js-ojus/flow"
	// "github.com/3xxx/meritms/models"
	// _ "github.com/mattn/go-sqlite3"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	"log"
)

// FLOW API
type FlowController struct {
	beego.Controller
}

// type WFApi struct {
// 	flow.Document
// }

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

// @Title show wf list
// @Description show workflow page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /workflow [get]
//页面
func (c *FlowController) WorkFlow() {
	c.TplName = "merit/workflow.tpl"
}

// @Title get wf list
// @Description get workflowlist by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /workflowdata [get]
//初始化数据库数据——测试
func (c *FlowController) WorkFlowData() {
	driver, connStr := "sqlite3", "database/meritms.db"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()
	dtID1, err := flow.DocTypes.New(tx, "图纸设计") //dtID1
	if err != nil {
		fmt.Println(err)
	}
	// dtID2, err := flow.DocTypes.New(tx, "图纸校核") //dtID1
	// if err != nil {
	// 	fmt.Println(err)
	// }
	_, err = flow.DocTypes.New(tx, "变更立项") //dtID2
	if err != nil {
		fmt.Println(err)
	}
	dsID1, err := flow.DocStates.New(tx, "设计中...") //初始化
	if err != nil {
		fmt.Println(err)
	}
	dsID2, err := flow.DocStates.New(tx, "校核中...") //委托创建
	if err != nil {
		fmt.Println(err)
	}
	_, err = flow.DocStates.New(tx, "审查中...")
	if err != nil {
		fmt.Println(err)
	}
	flow.DocStates.New(tx, "批注中...")
	flow.DocStates.New(tx, "申报中...")
	flow.DocStates.New(tx, "评估中...")
	flow.DocStates.New(tx, "审批中...")
	// flow.DocStates.New(tx, "DataApproved")
	// flow.DocStates.New(tx, "ReportGen")
	// flow.DocStates.New(tx, "ReportApproved")

	daID1, err := flow.DocActions.New(tx, "提交设计", false) //改变状态设计中...为校核中...
	if err != nil {
		fmt.Println(err)
	}
	daID2, err := flow.DocActions.New(tx, "校核", false)
	if err != nil {
		fmt.Println(err)
	}
	daID3, err := flow.DocActions.New(tx, "审查", false)
	if err != nil {
		fmt.Println(err)
	}
	daID4, err := flow.DocActions.New(tx, "核定", true)
	if err != nil {
		fmt.Println(err)
	}
	daID5, err := flow.DocActions.New(tx, "评估", true)
	if err != nil {
		fmt.Println(err)
	}
	daID6, err := flow.DocActions.New(tx, "审批", false)
	if err != nil {
		fmt.Println(err)
	}
	daID7, err := flow.DocActions.New(tx, "立项", false)
	if err != nil {
		fmt.Println(err)
	}

	workflowID1, err := flow.Workflows.New(tx, "图纸设计流程", dtID1, dsID1) //初始状态是“设计中...”
	if err != nil {
		fmt.Println(err)
	}
	// workflowID2, err := flow.Workflows.New(tx, "图纸校核流程", dtID2, dsID2) //初始状态是“设计中...”
	// if err != nil {
	// 	fmt.Println(err)
	// }

	accessContextID1, err := flow.AccessContexts.New(tx, "Context")
	if err != nil {
		beego.Error(err)
	}
	// AddNode maps the given document state to the specified node.  This
	// map is consulted by the workflow when performing a state transition
	// of the system.nodeID1
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID1, accessContextID1, workflowID1, "图纸三角校审流程-设计", flow.NodeTypeLinear)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	_, err = flow.Workflows.AddNode(tx, dtID1, dsID2, accessContextID1, workflowID1, "图纸三角校审流程-校核", flow.NodeTypeEnd)
	if err != nil {
		fmt.Println(err)
	}

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

	gID1 := fatal1(flow.Groups.New(tx, "设计人员组", "G")).(flow.GroupID)
	gID2 := fatal1(flow.Groups.New(tx, "校核人员组", "G")).(flow.GroupID)
	fatal0(flow.Groups.AddUser(tx, gID1, uID1))
	fatal0(flow.Groups.AddUser(tx, gID1, uID2))
	fatal0(flow.Groups.AddUser(tx, gID1, uID3))

	fatal0(flow.Groups.AddUser(tx, gID2, uID2))
	fatal0(flow.Groups.AddUser(tx, gID2, uID3))
	fatal0(flow.Groups.AddUser(tx, gID2, uID4))
	roleID1 := fatal1(flow.Roles.New(tx, "设计人员角色")).(flow.RoleID)
	roleID2 := fatal1(flow.Roles.New(tx, "校核人员角色")).(flow.RoleID)
	//给角色赋予action权限
	fatal0(flow.Roles.AddPermissions(tx, roleID1, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4}))
	fatal0(flow.Roles.AddPermissions(tx, roleID2, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4, daID5, daID6, daID7}))
	//给用户组赋予角色
	err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID1, roleID1)
	if err != nil {
		beego.Error(err)
	}
	err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID2, roleID2)
	if err != nil {
		beego.Error(err) //UNIQUE constraint failed: wf_ac_group_roles.ac_id已修补
	}
	//这里缺少contex的衔接？？上面的就是衔接上contex了吧。
	tx.Commit() //这个必须要！！！！！！

	wflist, err := flow.DocTypes.List(0, 0)
	if err != nil {
		beego.Error(err)
	}

	// wflist, err = flow.DocTypes.GetByName("Compute Request")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(wflist)

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
	// beego.Info(wflist1)
	// Users.List("", 0, 0)

	// Users.List("LN 4", 0, 0)
	// Groups.List(0, 0)
	// Roles.List(0, 0)
	// DocTypes.GetByName("Compute Request")
	c.Data["json"] = wflist
	c.ServeJSON()
}

// @Title post wf state
// @Description post workflow state
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgetdoctypebyname [get]
//模拟流程中提交——修改文档状态
func (c *FlowController) FlowGetDocTypeByName() {
	driver, connStr := "sqlite3", "database/meritms.db"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}
	// db.Close()
	//查询doctype
	dtID1, err := flow.DocTypes.GetByName("图纸设计")
	if err != nil {
		beego.Error(err)
	}
	beego.Info(dtID1)
	//查询context
	accessContextID1, err := flow.AccessContexts.List("Context", 0, 0)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(accessContextID1[0].ID)
	beego.Info(flow.GroupID(1))
	docNewInput := flow.DocumentsNewInput{
		DocTypeID:       dtID1.ID,
		AccessContextID: accessContextID1[0].ID,
		GroupID:         5, //groupId,
		Title:           "厂房布置图",
		Data:            "设计、制图: 秦晓川1, 校核: 秦晓川2",
	}
	// flow.Documents.New(tx, &docNewInput)
	DocumentID1, err := flow.Documents.New(tx, &docNewInput)
	if err != nil {
		beego.Error(err)
	}
	// tx.Commit() //new后面一定要跟commit啊
	beego.Info(DocumentID1)
	dsID2, err := flow.DocStates.GetByName("校核中...")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(dsID2)
	daID2, err := flow.DocActions.GetByName("提交设计")
	if err != nil {
		fmt.Println(err)
	}
	beego.Info(daID2)
	beego.Info(flow.GroupID(2))
	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   dtID1.ID, //flow.DocTypeID(1),
		DocumentID:  DocumentID1,
		DocStateID:  dsID2.ID,
		DocActionID: daID2.ID, //flow.DocActionID(2),
		GroupID:     6,
		Text:        "校核",
	}

	docEventID1, err := flow.DocEvents.New(tx, &docEventInput)
	if err != nil {
		beego.Error(err)
	}
	tx.Commit() //一个函数里只能有一个commit！！！！
	beego.Info(docEventID1)
	myDocEvent, err := flow.DocEvents.Get(docEventID1)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(myDocEvent)
	// myWorkflow, err := flow.Workflows.Get(workflowId.ID)
	// if err != nil {
	// 	beego.Error(err)
	// }
	myWorkflow, err := flow.Workflows.GetByName("图纸设计流程")
	if err != nil {
		beego.Error(err)
	}
	beego.Info(myWorkflow)
	//给出接受的组groupids
	groupIds := []flow.GroupID{flow.GroupID(6)}
	beego.Info(groupIds)
	newDocStateId, err := myWorkflow.ApplyEvent(tx, myDocEvent, groupIds)
	if err != nil {
		beego.Error(err)
	}
	tx.Commit() //一个函数里只能有一个commit！！！！
	fmt.Println("newDocStateId=", newDocStateId, err)

	// wflist, err := flow.DocTypes.GetByName("Compute Request")
	// if err != nil {
	// 	beego.Error(err)
	// }

	c.Data["json"] = "OK" //wflist
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

// Initialise DB connection.
// func init() {
// gt = t

// Connect to the database.travis
// driver, connStr := "mysql", "root:root@/flow"
// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// flow.RegisterDB(tdb)

// if tdb == nil {
// 	log.Fatal("given database handle is `nil`")
// }
// db = tdb

// return nil
// }

// func RegisterDB(sdb *sql.DB) error {
// 	if sdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db = sdb

// 	return nil
// }wflist []*flow.DocState

// @Title get wf list
// @Description get workflowlist by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /workflowdata1 [get]
//初始化数据库数据——测试
func (c *FlowController) WorkFlowData1() {
	// driver, connStr := "mysql", "root:root@/flow"
	// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	// if tdb == nil {
	// 	log.Fatal("given database handle is `nil`")
	// }
	// db := tdb
	driver, connStr := "sqlite3", "database/meritms.db"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	// var tx *sql.Tx
	defer tx.Rollback()

	// docType1, err := flow.DocTypes.New(tx, "EXAM:COMMON")
	// DocTypes.New(tx, "Stor Request")
	// DocTypes.New(tx, "Compute Request")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// docState1, err := flow.DocStates.New(tx, "Init") //初始化
	// flow.DocStates.New(tx, "EntrustCreate")          //委托创建
	// flow.DocStates.New(tx, "EntrustApproved")        //委托审核
	// flow.DocStates.New(tx, "SampleHandon")           //样本交接
	// flow.DocStates.New(tx, "PrepareFinish")          //制备完成
	// flow.DocStates.New(tx, "PrepareApproved")        //制备审核
	// flow.DocStates.New(tx, "TaskAssign")             //任务分配
	// flow.DocStates.New(tx, "DataApproved")           //数据录入审核
	// flow.DocStates.New(tx, "ReportGen")              //报告生成
	// flow.DocStates.New(tx, "ReportApproved")         //报告审核

	// docActionID1, _ := flow.DocActions.New(tx, "CreateEntrust", false)  //创建委托
	// docActionID2, _ := flow.DocActions.New(tx, "ApproveEntrust", false) //审核委托
	// docActionID3, _ := flow.DocActions.New(tx, "HandonSample", false)   //提交样本
	// docActionID4, _ := flow.DocActions.New(tx, "FinishPrepare", true)   //完成制备
	// docActionID5, _ := flow.DocActions.New(tx, "ApprovePrepare", true)  //审核制备
	// docActionID6, _ := flow.DocActions.New(tx, "AssignTask", false)     //分配任务
	// docActionID7, _ := flow.DocActions.New(tx, "ApproveData", false)    //审核数据
	// docActionID8, _ := flow.DocActions.New(tx, "GenReport", false)      //生成报告
	// docActionID9, _ := flow.DocActions.New(tx, "ApproveReport", true)   //审核报告

	// workflowId, _ := flow.Workflows.New(tx, "Examination", docType1, docState1)

	// flow.Workflows.SetActive(tx, workflowId, true)

	//创建Docments
	// contextId, _ := flow.AccessContexts.New(tx, "Context")
	// groupId, err := flow.Groups.New(tx, "Examination", "G")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// resUser, _ := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 	VALUES('admin', 'dashoo', 'admin@dashoo.com', 1)`)
	// uid, _ := resUser.LastInsertId()
	// userID1 := flow.UserID(uid)

	// flow.Groups.AddUser(tx, groupId, 5)
	// roleID1, _ := flow.Roles.New(tx, "administrator")
	flow.Roles.AddPermissions(tx, 5, 4, []flow.DocActionID{8, 9,
		10, 11, 12, 13, 14, 15, 16})

	docNewInput := flow.DocumentsNewInput{
		DocTypeID:       4, //docType1,
		AccessContextID: 1, //contextId,
		GroupID:         7,
		Title:           "entrust flow",
		Data:            "eid: 111, entrustNo: 2222",
	}
	// flow.Documents.New(tx, &docNewInput)

	_, err = flow.Documents.New(tx, &docNewInput)
	if err != nil {
		beego.Error(err)
	}
	// Documents.setState(tx, docType1, documentID1, docState2, contextId)
	tx.Commit()
	// wflist, err = flow.DocStates.List(0, 0)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// return wflist
}

// @Title post wf state
// @Description post workflow state
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgetdoctypebyname1 [get]
//模拟流程中提交——修改文档状态
func (c *FlowController) FlowGetDocTypeByName1() {

	// tx, _ := flow.db.Begin()
	var tx *sql.Tx
	defer tx.Rollback()

	/*docNewInput := DocumentsNewInput {
	  	DocTypeID: docType1,
	  	AccessContextID: contextId,
	  	GroupID: groupId,
	  	Title: "entrust flow",
	  	Data: "eid: 111, entrustNo: 2222",
	  }
	  documentID1, err := Documents.New(tx, &docNewInput)
	  fmt.Println("documentID1=", documentID1, err)*/

	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   flow.DocTypeID(4),
		DocumentID:  flow.DocumentID(1),
		DocStateID:  flow.DocStateID(9),
		DocActionID: flow.DocActionID(2),
		GroupID:     flow.GroupID(1),
		Text:        "开始审批",
	}
	groupIds := []flow.GroupID{flow.GroupID(1)}
	myWorkflow, err := flow.Workflows.Get(flow.WorkflowID(3))
	docEvent1, err := flow.DocEvents.New(tx, &docEventInput)
	tx.Commit()
	myDocEvent, err := flow.DocEvents.Get(docEvent1)
	newDocStateId, err := myWorkflow.ApplyEvent(tx, myDocEvent, groupIds)
	tx.Commit()
	fmt.Println("newDocStateId=", newDocStateId, err)
}
