package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"time"
)

type MeritTopic struct {
	Id      int64  `PK`
	MeritId int64  `orm:"null"`
	UserId  int64  `orm:"null"`
	Title   string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)" json:"title"` //orm:"unique",
	// Choose  string `orm:"null"`
	Content string `orm:"null" json:"content"`
	State   int    //1编写状态，未提交；2编写者提交，等待审核确认;3,已经审核确认
	Active  bool   `default(true) json:"active"`
	// Mark     string    `orm:"null"` //设置分数
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
	// User    User      `gorm:"foreignkey:UserId"`
}

func init() {
	orm.RegisterModel(new(MeritTopic))
	// _db.CreateTable(&MeritTopic{})
}

//用户添加价值
func AddMerit(mid, userid int64, title, content string, active bool) (id int64, err error) {
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
		// Choose:  choose,
		Content: content,
		State:   1,
		Active:  active,
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

//根据价值id和用户id取得状态=state的价值
func GetMerit(mid, uid int64, state int) (merits []*MeritTopic, err error) {
	o := orm.NewOrm()
	merits = make([]*MeritTopic, 0)
	// category := new(Merit)
	qs := o.QueryTable("merit_topic") //这个表名Merit需要用驼峰式，
	if mid != 0 && uid != 0 {         //如果给定父id则进行过滤
		_, err = qs.Filter("merit_id", mid).Filter("user_id", uid).Filter("state", state).All(&merits) //而这个字段parentid为何又不用呢
		if err != nil {
			return merits, err
		}
		// return merits, err
	} else if mid == 0 && uid != 0 { //则取所有这个用户的
		_, err = qs.Filter("user_id", uid).Filter("state", state).All(&merits) //而这个字段parentid为何又不用呢
		if err != nil {
			return merits, err
		}
		// return merits, err
	}
	return merits, err
}

//价值表带mark分值
type UserMeritTopic struct {
	// gorm.Model
	Id         int64     `json:"id"`
	MeritCate  string    `json:"meritcate"`  //价值分类——项目管理类
	MeritTitle string    `json:"merittitle"` //价值-专业负责人
	Choose     string    `json:"choose"`     //大型、中型
	TopicTitle string    `json:"topictitle"` //价值内容名称
	Mark       int       `json:"mark"`
	Content    string    `json:"content"`
	Updated    time.Time `json:"updated"`
	State      int       `json:"state"`
	Active     bool      `json:"active"`
	User       User
	// Children []UserMerit
}

//根据用户id和价值分类2级，取得所有这个用户的价值
func GetMeritTopic(mid, uid int64, state int) (usermerittopics []*UserMeritTopic, err error) {
	db := GetDB()
	// 必须要写权select，坑爹啊 OR admin_merit.id=?
	err = db.Table("admin_merit").
		Select("merit_topic.id as id,t1.title as merit_title,t2.title as merit_cate,admin_merit.title as choose,admin_merit_mark.mark as mark,merit_topic.title as topic_title,merit_topic.content as content,merit_topic.state as state,merit_topic.updated as updated,merit_topic.active as active").
		Where("admin_merit.parent_id=?", mid).
		Joins("left JOIN merit_topic on merit_topic.merit_id=admin_merit.id").
		Where("merit_topic.user_id = ? AND merit_topic.state=?", uid, state).
		Joins("INNER JOIN admin_merit AS t1 ON t1.id = admin_merit.parent_id").
		Joins("INNER JOIN admin_merit AS t2 ON t2.id = t1.parent_id").
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = merit_topic.merit_id").
		Joins("left JOIN User on user.id = merit_topic.user_id").
		Scan(&usermerittopics).Error
	return usermerittopics, err
}

//统计用户价值表带mark分值
type UserMeritTopics struct {
	Marks      int    `json:"marks"`
	Number     int    `json:"number"`
	Nickname   string `json:"nickname"`
	Department string `json:"department"`
	UserId     int64  `json:"userid"`
	User       User   `gorm:"foreignKey:UserId"`
}

//根据用户id，取得所有这个用户的价值分值汇总和价值数量汇总
//和GetMyselfMeritTopic区别是进行了汇总
func GetMeritTopics(uid int64, state int, active bool) (usermerittopics []*UserMeritTopics, err error) {
	db := GetDB()
	// 必须要写权select，坑爹啊
	err = db.Table("merit_topic").Select("sum(admin_merit_mark.mark) as marks,count(*) as number,user.id as user_id,user.nickname as nickname,user.department as department").
		Where("merit_topic.user_id = ? AND merit_topic.state=? AND merit_topic.active=?", uid, state, active).
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = merit_topic.merit_id").
		Joins("left JOIN user").Where("user.id=?", uid).
		Group("user.id").
		Find(&usermerittopics).Error
	// err = db.Table("user").Select("sum(admin_merit_mark.mark) as marks,count(*) as number,user.id as user_id,user.nickname as user_nickname").Where("user.id = ?", uid).
	// 	Joins("left JOIN merit_topic on user.id = merit_topic.user_id").Where("merit_topic.state=?", state).Group("merit_topic.id").
	// 	Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = merit_topic.merit_id").
	// 	Joins("User").
	// 	Find(&usermerittopics).Error
	return usermerittopics, err
}

type MyMeritTopic struct {
	Id           int64  `json:"id"`
	MeritCate    string `json:"meritcate"`  //价值分类——项目管理类
	MeritTitle   string `json:"merittitle"` //价值-专业负责人
	UserNickName string `json:"usernickname"`
	TopicTitle   string `json:"topictitle"`
	Choose       string `json:"choose"` //大型、中型
	Content      string `json:"content"`
	Active       bool   `json:"active"`
	// State        int       //1编写状态，未提交；2编写者提交，等待审核确认;3,已经审核确认
	Mark int `json:"mark"`
	// Created      time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated"`
}

//取得用户id状态为state的所有价值，代替下面那个
//GetMeritTopics区别是-没有进行汇总
func GetMyselfMeritTopic(uid int64, state int, active bool) (merittopics []*MyMeritTopic, err error) {
	db := GetDB()
	// 必须要写权select，坑爹啊
	err = db.Table("merit_topic").
		Select("merit_topic.id as id,merit_topic.title as topic_title,merit_topic.content as content,merit_topic.updated as updated,merit_topic.active as active,t1.title as merit_title,t2.title as merit_cate,admin_merit_mark.mark as mark,admin_merit.title as choose,user.nickname as user_nick_name").
		Where("merit_topic.user_id = ? AND merit_topic.state=?AND merit_topic.active=?", uid, state, active).
		Joins("INNER JOIN admin_merit AS t1 ON t1.id = admin_merit.parent_id").
		Joins("INNER JOIN admin_merit AS t2 ON t2.id = t1.parent_id").
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = merit_topic.merit_id").
		Joins("left JOIN User on user.id = merit_topic.user_id").
		Joins("left JOIN admin_merit on admin_merit.id = merit_topic.merit_id").
		Scan(&merittopics).Error
	return merittopics, err
}

//取得用户id的所有状态为3的价值——作废
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
		// type Duration int64
		// const (
		// 	Nanosecond  Duration = 1
		// 	Microsecond          = 1000 * Nanosecond
		// 	Millisecond          = 1000 * Microsecond
		// 	Second               = 1000 * Millisecond
		// 	Minute               = 60 * Second
		// 	Hour                 = 60 * Minute
		// )
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
		// case "Choose":
		// 	merit.Choose = value
		// 	_, err := o.Update(&merit, "Choose", "Updated")
		// 	if err != nil {
		// 		return err
		// 	} else {
		// 		return nil
		// 	}
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
		case "Active":
			if value == "true" {
				merit.Active = true
			} else {
				merit.Active = false
			}
			_, err := o.Update(&merit, "Active", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "active":
			if value == "true" {
				merit.Active = true
			} else {
				merit.Active = false
			}
			_, err := o.Update(&merit, "Active", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "topictitle":
			merit.Title = value
			_, err := o.Update(&merit, "topict_title", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "choose":
			// 根据choose找到meritid，不容易，topicid——chooseid——parentid——所有子类，比对value
			_, err := o.Update(&merit, "Active", "Updated")
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

//查询某个用户userid的价值分值
func GetUserMeritMark(uid, mid int64, state int, active bool) (mark int, err error) {
	//获取DB Where("product.title LIKE ?", "%searchText%").不对
	db := GetDB()
	// 必须要写权select，坑爹啊
	err = db.Table("merit_topic").Select("sum(admin_merit_mark.mark) as amount").
		Where("merit_topic.state=? AND merit_topic.user_id=?AND merit_topic.active=?", state, uid, active).
		Joins("left JOIN admin_merit on admin_merit.id = merit_topic.merit_id").
		Joins("left join admin_merit_mark on admin_merit_mark.merit_id = admin_merit.id").
		Scan(&mark).Error
	return mark, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//测试用
type MeritTopicUser struct {
	AdminMeritMark `xorm:"extends"`
	MeritTopic     `xorm:"extends"`
	User           `xorm:"extends"`
}

//查询结果
// [
//   {
//     "Mark": 4,
//     "UserId": 1490,
//     "title": "阿斯顿发生发送的",
//     "content": "阿斯蒂芬",
//     "State": 3,
//     "name": "qin.xc",
//     "Nickname": "秦晓川",
//     "Password": "efadc8586557774bf7729250ca71f114",
//     "Repassword": "",
//     "Email": "",
//     "Department": "施工预算分院",
//     "Secoffice": "水工室",
//     "Remark": "",
//     "Ip": "127.0.0.1",
//     "Port": "80",
//     "Status": 1,
//     "Lastlogintime": "2020-09-08T20:45:06.1394044+08:00",
//     "Createtime": "2016-09-03T14:24:19.3756176+08:00",
//     "role": "3"
//   }
// ]

//测试xorm的join查询，返回结构体无法嵌套
func GetMeritTopicUser(uid int64, state int) ([]*MeritTopicUser, error) {
	merittopics := make([]*MeritTopicUser, 0)
	return merittopics, engine.Table("merit_topic").
		Where("merit_topic.user_id = ? AND merit_topic.state=?", uid, state).
		Join("INNER", "user", "merit_topic.user_id = user.id").
		Join("INNER", "admin_merit_mark", "admin_merit_mark.merit_id = merit_topic.merit_id").
		Find(&merittopics)
}

//测试gorm自定义嵌套结构体，无效果
type MeritTopicUser2 struct {
	AdminMeritMark AdminMeritMark
	MeritTopic     MeritTopic `json:"merittopic"`
	User           User       `json:"user"`
}

//测试gorm返回嵌套结构体，必须是建表的数据结构，不能自定义数据结构。建表的数据struct必须有foreignKey
func GetMeritTopics2(uid int64, state int) (topics []*MeritTopic, err error) {
	db := GetDB()
	// 必须要写权select，坑爹啊
	err = db.Preload("User").Preload("merit_topic").Where("user_id=?", uid).Find(&topics).Error //查询所有device记录
	return topics, err
}

//测试 价值表带mark分值
type UserMeritTopic2 struct {
	// gorm.Model
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Title1     string `json:"title1"`     //价值分类名称
	Title2     string `json:"title2"`     //价值分类名称
	Title3     string `json:"title3"`     //价值分类名称
	TopicTitle string `json:"topictitle"` //价值内容名称
	Mark       int    `json:"mark"`
	Content    string `json:"content"`
	User       User
	// Children []UserMerit
}

//测试子查询，希望查询出树状结构表的父子孙
//最后用join查询重命名来解决了。
func GetMeritTopic2(mid, uid int64, state int) (usermerittopics []*UserMeritTopic2, err error) {
	db := GetDB()
	// subQuery1 := db.Model(&AdminMerit{}).Select("title").Where("id=?", mid)
	// subQuery2 := db.Model(&AdminMerit{}).Select("title").Where("parent_id=?", mid)
	// db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
	// db.Table("(?) as title1, (?) as title2", subQuery1, subQuery2).

	// subQuery := db.Select("admin_merit.id as id,admin_merit.title as title1").Table("admin_merit").Where("id=?", mid).SubQuery()
	err = db.Table("admin_merit").Select("t2.title as title2,t1.title as title1,admin_merit.title as title,admin_merit_mark.mark as mark,merit_topic.title as topic_title,merit_topic.content as content").
		Where("admin_merit.parent_id=?", mid).
		Joins("INNER JOIN admin_merit AS t1 ON t1.id = admin_merit.parent_id").
		Joins("INNER JOIN admin_merit AS t2 ON t2.id = t1.parent_id").
		Joins("left JOIN merit_topic on merit_topic.merit_id=admin_merit.id").
		Where("merit_topic.user_id = ? AND merit_topic.state=?", uid, state).
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = merit_topic.merit_id").
		Scan(&usermerittopics).Error
	return usermerittopics, err
}

//根据价值内容id取得价值内容和分值——没用到
// func GetMeritTopic(mid, uid int64, state int) (merit *MeritTopic, marks int, err error) {
// 	o := orm.NewOrm()
// 	// category := new(Merit)
// 	qs := o.QueryTable("MeritTopic")                              //这个表名Merit需要用驼峰式，
// 	err = qs.Filter("Id", mid).Filter("state", state).One(&merit) //而这个字段parentid为何又不用呢
// 	if err != nil {
// 		return merit, 0, err
// 	}
// 	//根据choose取得adminmerit分值
// 	adminmerit, err := GetAdminMeritbyId(merit.MeritId)
// 	var ff string
// 	// 如果mark为空，则寻找选择列表的分值，如果不为空，则直接用价值的分值
// 	// 进行选择列表拆分
// 	array1 := strings.Split(adminmerit.List, ",")
// 	array2 := strings.Split(adminmerit.ListMark, ",")
// 	if adminmerit.Mark == "" {
// 		for i1, _ := range array1 {
// 			// if v1 == v.choose {
// 			ff = array2[i1]
// 			// }
// 		}
// 	} else {
// 		ff = adminmerit.Mark
// 	}
// 	markint, err := strconv.Atoi(ff)
// 	if err != nil {
// 		return merit, 0, err
// 	}
// 	return merit, markint, err
// }
