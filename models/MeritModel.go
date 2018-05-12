package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)

type MeritTopic struct {
	Id      int64  `PK`
	MeritId int64  `orm:"null"`
	UserId  int64  `orm:"null"`
	Title   string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Choose  string `orm:"null"`
	Content string `orm:"null"`
	State   int    //1编写状态，未提交；2编写者提交，等待审核确认;3,已经审核确认

	// Mark     string    `orm:"null"` //设置分数
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(MeritTopic)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

//用户添加价值
func AddMerit(mid, userid int64, title, choose, content string) (id int64, err error) {
	//先由uname取得uid
	// user, err := GetUserByUsername(uname)
	// if err != nil {
	// 	return 0, err
	// }
	o := orm.NewOrm()
	merit := &MeritTopic{
		MeritId: mid,
		UserId:  userid,
		Title:   title,
		Choose:  choose,
		Content: content,
		State:   1,
		Created: time.Now(),
		Updated: time.Now(),
	}
	// qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	id, err = o.Insert(merit)
	if err != nil {
		return id, err
	}
	return id, err
}

//根据价值内容id取得价值内容和分值
func GetMeritTopic(mid, uid int64, state int) (merit *MeritTopic, marks int, err error) {
	o := orm.NewOrm()
	// category := new(Merit)
	qs := o.QueryTable("MeritTopic")                              //这个表名Merit需要用驼峰式，
	err = qs.Filter("Id", mid).Filter("state", state).One(&merit) //而这个字段parentid为何又不用呢
	if err != nil {
		return merit, 0, err
	}
	//根据choose取得adminmerit分值
	adminmerit, err := GetAdminMeritbyId(merit.MeritId)
	var ff string
	// 如果mark为空，则寻找选择列表的分值，如果不为空，则直接用价值的分值
	// 进行选择列表拆分
	array1 := strings.Split(adminmerit.List, ",")
	array2 := strings.Split(adminmerit.ListMark, ",")
	if adminmerit.Mark == "" {
		for i1, _ := range array1 {
			// if v1 == v.choose {
			ff = array2[i1]
			// }
		}
	} else {
		ff = adminmerit.Mark
	}
	markint, err := strconv.Atoi(ff)
	if err != nil {
		return merit, 0, err
	}
	return merit, markint, err
}

//根据价值id和用户id取得已经完成的价值
func GetMerit(mid, uid int64, state int) (merits []*MeritTopic, err error) {
	o := orm.NewOrm()
	merits = make([]*MeritTopic, 0)
	// category := new(Merit)
	qs := o.QueryTable("MeritTopic") //这个表名Merit需要用驼峰式，
	if mid != 0 && uid != 0 {        //如果给定父id则进行过滤
		_, err = qs.Filter("MeritId", mid).Filter("userid", uid).Filter("state", state).All(&merits) //而这个字段parentid为何又不用呢
		if err != nil {
			return merits, err
		}
		// return merits, err
	} else if mid == 0 && uid != 0 { //则取所有这个用户的
		_, err = qs.Filter("userid", uid).Filter("state", state).All(&merits) //而这个字段parentid为何又不用呢
		if err != nil {
			return merits, err
		}
		// return merits, err
	}
	return merits, err
}

//取得用户id的所有状态为3的价值
func GetMyselfMerit(uid int64) ([]*MeritTopic, error) {
	o := orm.NewOrm()
	topics := make([]*MeritTopic, 0)
	// category := new(Merit)
	qs := o.QueryTable("merit_topic")                                  //这个表名Merit需要用驼峰式，
	_, err := qs.Filter("userid", uid).Filter("state", 3).All(&topics) //而这个字段userid为何又不用呢
	if err != nil {
		return topics, err
	}
	return topics, err
}

//根据topicid取得topic
func GetMeritbyId(tid string) (*MeritTopic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	Merit := new(MeritTopic)
	qs := o.QueryTable("merit_topic")
	err = qs.Filter("id", tidNum).One(Merit)
	if err != nil {
		return nil, err
	}
	return Merit, err
}

//删除Merit
func DeleteMerit(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	Merit := MeritTopic{Id: id}
	if o.Read(&Merit) == nil {
		_, err := o.Delete(&Merit)
		if err != nil {
			return err
		}
	}
	return nil
}

//用户修改一个用户的某个字段
func UpdateMerit(id int64, fieldname, value string) error {
	o := orm.NewOrm()
	var merit MeritTopic
	// user := &User{Id: cid}
	err := o.QueryTable("MeritTopic").Filter("Id", id).One(&merit)
	// err:=o.Read(user).One()
	if err == nil {
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

		const lll = "2006-01-02"
		merit.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		switch fieldname {
		case "Title":
			merit.Title = value
			_, err := o.Update(&merit, "Title", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Choose":
			merit.Choose = value
			_, err := o.Update(&merit, "Choose", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Content":
			merit.Content = value
			_, err := o.Update(&merit, "Content", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "State":
			merit.State, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
			_, err := o.Update(&merit, "State", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}

		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	}
	return err
}

//修改Merit
// func ModifyMerit(mid, title, choose, content string) error {
// 	tidNum, err := strconv.ParseInt(mid, 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	o := orm.NewOrm()
// 	Merit := &MeritTopic{Id: tidNum}
// 	if o.Read(Merit) == nil {
// 		Merit.Title = title
// 		Merit.Choose = choose
// 		Merit.Content = content
// 		// Merit.Mark = mark
// 		Merit.Updated = time.Now()
// 		_, err = o.Update(Merit)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return err
// }

//管理员取得所有价值

//根据分院和科室名称查所有价值
func GetMeritsbySec(department, secoffice string) (meritcates []*AdminDepartMerit, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	qs := o.QueryTable("AdminDepartment")
	//这里进行过滤
	//由分院和科室名称获得科室id
	var depart AdminDepartment
	var secoff AdminDepartment
	err = qs.Filter("Title", department).One(&depart)
	if err != nil {
		return nil, err
	}

	err = qs.Filter("ParentId", depart.Id).Filter("Title", secoffice).One(&secoff)
	if err != nil {
		return nil, err
	}
	//取得所有价值分类id
	qs2 := o.QueryTable("AdminDepartMerit")
	_, err = qs2.Filter("SecofficeId", secoff.Id).All(&meritcates)
	if err != nil {
		return nil, err
	}
	return meritcates, err
}

//根据分院名称查所有价值——适用于没有科室的部门
//查出所有价值，只有分院（部门）而没科室字段的价值
func GetMeritsbySecOnly(department string) (meritcates []*AdminDepartMerit, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	//这里进行过滤
	//由分院和科室名称获得科室id
	var depart AdminDepartment
	err = qs.Filter("title", department).One(&depart)
	if err != nil {
		return nil, err
	}
	qs2 := o.QueryTable("AdminDepartMerit")
	_, err = qs2.Filter("SecofficeId", depart.Id).All(&meritcates)
	if err != nil {
		return nil, err
	}
	return meritcates, err
}

//根据科室id查所有价值
func GetMeritsbySecId(secofficeid string) (users []*User, count int, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	qs := o.QueryTable("user")
	//这里进行过滤
	secid, err := strconv.ParseInt(secofficeid, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	//由secid查自身科室名称
	secoffice, err := GetAdminDepartbyId(secid)
	if err != nil {
		return nil, 0, err
	}
	//由secoffice的pid查分院名称
	department, err := GetAdminDepartbyId(secoffice.ParentId)
	if err != nil {
		return nil, 0, err
	}
	//由分院名称和科室名称查所有用户
	_, err = qs.Filter("Department", department.Title).Filter("Secoffice", secoffice.Title).OrderBy("Username").All(&users)
	if err != nil {
		return nil, 0, err
	}
	// _, err = qs.OrderBy("-created").All(&cates)
	// _, err := qs.All(&cates)
	count = len(users)
	return users, count, err
}
