package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/3xxx/meritms/models"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"github.com/3xxx/meritms/controllers"
	_ "github.com/3xxx/meritms/routers"
)

//自定义模板函数，序号加1
func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}

func main() {
	beego.AddFuncMap("indexaddone", Indexaddone) //模板中使用{{indexaddone $index}}或{{$index|indexaddone}}
	beego.AddFuncMap("loadtimes", loadtimes)

	//开启orm调试模式
	// orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	models.InsertUser()
	// time1 := "0/" + time + " * * * * *"

	time1 := "0 30 23 * * *"
	tk1 := toolbox.NewTask("tk1", time1, func() error { controllers.Postdata(); return nil }) //func() error { fmt.Println("tk1"); return nil }
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}

//显示页面加载时间
func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}
//错误描述：当controllers中的func中没有使用models中的func时，提示没有定义default数据库。
//也就是controllers中不使用models时，models中的init()不起作用

// <orm.RegisterModel> table name `category` repeat register, must be unique
//因为"merit/models"没有改到，原来是quick/models
