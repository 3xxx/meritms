package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
	// _ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"fmt"
	"time"
)

type AchievementTopic struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	UserId   int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Choose   string    `orm:"null"`
	Content  string    `orm:"null"`
	Mark     string    `orm:"null"` //设置分数
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)"`
}

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

func init() {
	var err error
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	if db_path == "" {
		db_path = "./"
	}

	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
	_db, err = gorm.Open(db_type, dns)
	// defer _db.Close()//20200803这个不能打开。
	// _db.LogMode(true)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// defer gdb.Close()
	//禁止表名复数形式
	_db.SingularTable(true)
	// 开发的时候需要打开调试日志
	// _db.LogMode(true)
	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	// _db.CreateTable(&Pay{}, &Money{}, &Recharge{})
	// if !gdb.HasTable(&Pay1{}) {
	// 	if err = gdb.CreateTable(&Pay1{}).Error; err != nil {
	// 		panic(err)
	// 	}
	// }
	orm.RegisterModel(new(AchievementTopic))
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
// 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接
func GetDB() *gorm.DB {
	return _db
}

// func init() {
// orm.RegisterModel(new(AchievementTopic)) //, new(Article)
// orm.RegisterDriver("sqlite", orm.DRSqlite)
// orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
// }

//用户添加价值
func AddAchievementTopic(pid int64, uname, title, choose, content, mark string) (id int64, err error) {
	//先由uname取得uid
	user, err := GetUserByUsername(uname)
	if err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	topic := &AchievementTopic{
		ParentId: pid,
		UserId:   user.Id,
		Title:    title,
		Choose:   choose,
		Content:  content,
		Mark:     mark,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	// qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	id, err = o.Insert(topic)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//根据父级价值id和用户id取得所有价值——返回数量和分值
//这个是拷贝jsonmodel的，这里没有用
func GetAchievementTopic(pid, uid int64) (topics []*AchievementTopic, numbers, marks int, err error) {
	o := orm.NewOrm()
	topics = make([]*AchievementTopic, 0)
	// category := new(AchievementTopic)
	qs := o.QueryTable("merit_topic") //这个表名AchievementTopic需要用驼峰式，
	if pid != 0 {                     //如果给定父id则进行过滤
		_, err = qs.Filter("parentid", pid).Filter("userid", uid).All(&topics) //而这个字段parentid为何又不用呢
		if err != nil {
			return nil, 0, 0, err
		}
		for _, v := range topics {
			mark, err := strconv.Atoi(v.Mark)
			if err != nil {
				return nil, 0, 0, err
			}
			marks = marks + mark
		}
		numbers = len(topics)
		return topics, numbers, marks, err
	} else { //如果不给定父id（PID=0），则取所有
		_, err = qs.Filter("userid", uid).All(&topics) //而这个字段parentid为何又不用呢
		if err != nil {
			return nil, 0, 0, err
		}
		for _, v := range topics {
			mark, err := strconv.Atoi(v.Mark)
			if err != nil {
				return nil, 0, 0, err
			}
			marks = marks + mark
		}
		numbers = len(topics)
		return topics, numbers, marks, err
	}
}

//由父级id得到所有下级
// func GetPids(pid int64) ([]*Category, error) {
// 	o := orm.NewOrm()
// 	cates := make([]*Category, 0)
// 	qs := o.QueryTable("category")
// 	var err error
// 	//这里进行过滤
// 	_, err = qs.Filter("ParentId", pid).All(&cates)
// 	return cates, err
// }
