package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
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

func init() {
	orm.RegisterModel(new(AchievementTopic)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

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
