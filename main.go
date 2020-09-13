package main

import (
	"github.com/3xxx/meritms/models"
	_ "github.com/3xxx/meritms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"time"
)

//自定义模板函数，序号加1
func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}

func main() {
	// InitDB()//导致数据库lock，不要用！
	beego.AddFuncMap("indexaddone", Indexaddone) //模板中使用{{indexaddone $index}}或{{$index|indexaddone}}
	beego.AddFuncMap("loadtimes", loadtimes)

	//开启orm调试模式
	// orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	models.InsertUser()

	beego.Run()
}

//显示页面加载时间
func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

// var db *sql.DB

// func InitDB() {
// 	var err error
// 	db, err = sql.Open("sqlite3", "database/meritms.db")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		log.Panic(err)
// 	}

// 	// driver, connStr := "sqlite3", "database/meritms.db"
// 	// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	// if tdb == nil {
// 	// 	log.Fatal("given database handle is `nil`")
// 	// }
// 	// db := tdb
// }

// // fatal1 expects a value and an error value as its arguments.
// func fatal1(val1 interface{}, err error) interface{} {
// 	if err != nil {
// 		fmt.Println("%v", err)
// 	}
// 	return val1
// }

//错误描述：当controllers中的func中没有使用models中的func时，提示没有定义default数据库。
//也就是controllers中不使用models时，models中的init()不起作用

// <orm.RegisterModel> table name `category` repeat register, must be unique
//因为"merit/models"没有改到，原来是quick/models
