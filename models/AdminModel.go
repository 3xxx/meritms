//部门结构，价值，ip地址段，日历
package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"log"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type AdminCategory struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Grade    int       `orm:"null"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

type AdminIpsegment struct {
	Id      int64     `form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	StartIp string    //`orm:"not null"`
	EndIp   string    `orm:"null"`
	Iprole  int       `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type AdminCalendar struct {
	Id        int64     `json:"id",form:"-"`
	Title     string    `json:"title",form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	Content   string    `json:"content",orm:"sie(20)"`
	Starttime time.Time `json:"start",orm:"type(datetime)"`
	Endtime   time.Time `json:"end",orm:"null;type(datetime)"`
	Allday    bool      `json:"allDay",orm:"default(0)"`
	Color     string    `json:"color",orm:"null"`
	Public    bool      `default(true)`
	// Color     string    `json:"backgroundColor",orm:"null"`
	// BColor    string    `json:"borderColor",orm:"null"`
}

//项目同步ip表
type AdminSynchIp struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	UserName string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	SynchIp  string    //`orm:"not null"`
	Port     string    `orm:"default(80)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

//科室结构
type AdminDepartment struct {
	Id       int64     `json:"id",form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)"`
}

//价值分类
type AdminMerit struct {
	Id       int64  `form:"-"`
	ParentId int64  `orm:"null"`
	Title    string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	// Mark     string    `orm:"null"`                                              //设置分数
	// List     string    `orm:"null"`                                              //选择项
	// ListMark string    `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
}

//价值分值
type AdminMeritMark struct {
	Id      int64     `form:"-"`
	MeritId int64     `orm:"null"`
	Mark    int       `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
}

type AdminDepartMerit struct {
	Id          int64 `json:"id",form:"-"`
	SecofficeId int64 `orm:"null"`
	MeritId     int64 `orm:"null"`
}

//ip地址段权限
// type AdminIpsegment struct {
// 	Id      int64  `form:"-"`
// 	Title   string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
// 	StartIp string
// 	EndIp   string    `orm:"null"`
// 	Iprole  int       `orm:"null"`
// 	Created time.Time `orm:"auto_now_add;type(datetime)"`
// 	Updated time.Time `orm:"auto_now_add;type(datetime)"`
// }

//日历
type AdminCalenda struct {
	Id        int64     `form:"-"`
	Title     string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	starttime time.Time `orm:"type(datetime)"`
	endtime   time.Time `orm:"null;type(datetime)"`
	allday    int8      `orm:"NOT NULL;default(0)"`
	color     string    `orm:"null"`
}

//首页轮播图片
type AdminCarousel struct {
	Id      int64     `form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Url     string    `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
}

// `id` int(11) NOT NULL AUTO_INCREMENT,
//   `title` varchar(100) NOT NULL,
//   `starttime` int(11) NOT NULL,
//   `endtime` int(11) DEFAULT NULL,
//   `allday` tinyint(1) NOT NULL DEFAULT '0',
//   `color` varchar(20) DEFAULT NULL,

func init() {
	orm.RegisterModel(new(AdminCategory), new(AdminIpsegment), new(AdminDepartment), new(AdminMerit), new(AdminMeritMark), new(AdminDepartMerit), new(AdminSynchIp), new(AdminCarousel), new(AdminCalendar)) //, new(Article)
	// orm.RegisterModel(new(AdminIpsegment))
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/meritms.db", 10)
	// 注册xorm
	var err error
	engine, err = xorm.NewEngine("sqlite3", "database/meritms.db")
	if err != nil {
		log.Println(err)
	}
}

//添加部门
func AddAdminDepart(pid int64, title, code string) (id int64, err error) {
	o := orm.NewOrm()
	depart := &AdminDepartment{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(depart)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateAdminDepart(cid int64, title, code string) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminDepart(cid int64) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		_, err := o.Delete(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据部门id取得所有科室
//如果父级id为0，则取所有部门
func GetAdminDepart(pid int64) (departs []*AdminDepartment, err error) {
	o := orm.NewOrm()
	departs = make([]*AdminDepartment, 0)
	qs := o.QueryTable("AdminDepartment")
	_, err = qs.Filter("parentid", pid).All(&departs)
	if err != nil {
		return nil, err
	}
	return departs, err
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
// func GetAdminCategory(pid int64) (categories []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	categories = make([]*AdminDepartment, 0)
// 	qs := o.QueryTable("AdminDepartment")
// 	_, err = qs.Filter("parentid", pid).All(&categories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, err
// }
//根据部门名字title返回自身
func GetAdminDepartName(title string) (AdminDepartment, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err := qs.Filter("title", title).One(&cate)
	if err != nil {
		return cate, err
	}
	return cate, err
}

//根据部门名字title查询所有下级科室category
func GetAdminDepartTitle(title string) (categories []*AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err = qs.Filter("title", title).One(&cate)
	categories = make([]*AdminDepartment, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
}

//根据id查科室
func GetAdminDepartbyId(id int64) (category AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return category, err
	}
	return category, err
}

//由分院id和科室 名称取得科室
func GetAdminDepartbyidtitle(id int64, title string) (*AdminDepartment, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(AdminDepartment)
	qs := o.QueryTable("AdminDepartment")
	err := qs.Filter("parentid", id).Filter("title", title).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//*********价值********************
//添加价值
func AddAdminMerit(pid int64, title string, mark int) (id int64, err error) {
	//重复性检查
	o := orm.NewOrm()
	cate := &AdminMerit{
		ParentId: pid,
		Title:    title,
		// Mark:     mark,
		// List:     list,
		// ListMark: listmark,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	if mark != 0 {
		mmark := &AdminMeritMark{
			MeritId: id,
			Mark:    mark,
			Created: time.Now(),
			Updated: time.Now(),
		}
		_, err = o.Insert(mmark)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

//由父级id（项目管理类）得到所有价值
func GetAdminMeritbyPid(pid int64) ([]*AdminMerit, error) {
	o := orm.NewOrm()
	cates := make([]*AdminMerit, 0)
	qs := o.QueryTable("AdminMerit")
	var err error
	//这里进行过滤
	_, err = qs.Filter("ParentId", pid).All(&cates)
	return cates, err
}

//价值表带mark分值
type UserMerit struct {
	// gorm.Model
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Mark  int    `json:"mark"`
	// Children []UserMerit
}

//根据父级id取到所有的价值结构
func GetAdminMerit(pid int64) (usermerit []*UserMerit, err error) {
	// o := orm.NewOrm()
	// merits := make([]*AdminMerit, 0)
	// qs := o.QueryTable("AdminMerit")
	// _, err := qs.Filter("parentid", pid).All(&merits)
	// if err != nil {
	// 	return nil, err
	// }
	// return merits, err
	db := GetDB()
	// 必须要写权select，坑爹啊
	err = db.Table("admin_merit").Select("admin_merit.id,admin_merit.title,admin_merit_mark.mark as mark").
		Where("admin_merit.parent_id=?", pid).
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = admin_merit.id").
		Scan(&usermerit).Error
	return usermerit, err
}

//由id取得价值
//要么本身能查到分值，要么要子价值查出分值
func GetAdminMeritbyId(id int64) (*AdminMerit, error) {
	o := orm.NewOrm()
	merit := new(AdminMerit)
	qs := o.QueryTable("AdminMerit")
	err := qs.Filter("id", id).One(merit)
	if err != nil {
		return nil, err
	}
	return merit, err
}

//由id取得价值分值
func GetAdminMeritMarkbyId(meritid int64) (*AdminMeritMark, error) {
	o := orm.NewOrm()
	meritmark := new(AdminMeritMark)
	qs := o.QueryTable("AdminMeritMark")
	err := qs.Filter("merit_id", meritid).One(meritmark)
	if err != nil {
		return nil, err
	}
	return meritmark, err
}

//由科室id取得所有价值分类
//根据分院名称查所有价值——适用于没有科室的部门
//查出所有价值，只有分院（部门）而没科室字段的价值
func GetSecofficeMerit(id int64) (meritcates []*AdminDepartMerit, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartMerit")
	_, err = qs.Filter("SecofficeId", id).All(&meritcates)
	if err != nil {
		return nil, err
	}
	return meritcates, err
}

//将secofficeid和meritid存入对应数据库
//如果存在，
func AddSecofficeMerit(sid, mid int64) error {
	//重复性检查
	o := orm.NewOrm()
	secmerit := &AdminDepartMerit{
		SecofficeId: sid,
		MeritId:     mid,
	}
	_, err := o.Insert(secmerit)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSecofficeMerit(sid, mid int64) error {
	o := orm.NewOrm()
	var merit AdminDepartMerit
	qs := o.QueryTable("AdminDepartMerit")
	_, err := qs.Filter("SecofficeId", sid).Filter("MeritId", mid).All(&merit)
	if err != nil {
		return err
	}
	// if o.Read(&merit) == nil {
	_, err = o.Delete(&merit) //删除分院
	if err != nil {
		return err
	}
	// }
	return err
}

//修改merit
func UpdateAdminMerit(id int64, title string, mark int) error {
	o := orm.NewOrm()
	merit := &AdminMerit{Id: id}
	var err error
	if o.Read(merit) == nil {
		merit.Title = title
		// merit.Mark = mark
		// merit.List = list
		// merit.ListMark = listmark
		merit.Updated = time.Now()
		_, err = o.Update(merit)
		if err != nil {
			return err
		}

		if mark != 0 {
			mmark := &AdminMeritMark{MeritId: id}
			//不是主键的话，要在read里加上字段名才行！！！
			if o.Read(mmark, "MeritId") == nil {
				mmark.Mark = mark
				mmark.Updated = time.Now()
				_, err = o.Update(mmark)
				if err != nil {
					return err
				}
			} else { //如果未查到，则插入
				mmark := &AdminMeritMark{
					MeritId: id,
					Mark:    mark,
					Created: time.Now(),
					Updated: time.Now(),
				}
				_, err = o.Insert(mmark)
				if err != nil {
					return err
				}
			}
		}
		// else {
		// 	array := strings.Split(list, ",")
		// 	array2 := strings.Split(listmark, ",")
		// 	for i, v := range array {
		// 		//下面有很多个，所以不行
		// 		cate := &AdminMerit{ParentId: id}
		// 		if o.Read(cate) == nil {
		// 			cate.Title = v
		// 			cate.Updated = time.Now()
		// 		}
		// 		_, err = o.Update(cate)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		markint, err := strconv.Atoi(array2[i])
		// 		if err != nil {
		// 			return err
		// 		}
		// 		mmark := &AdminMeritMark{MeritId: cate.Id}
		// 		if o.Read(mmark) == nil {
		// 			mmark.Mark = markint
		// 			mmark.Updated = time.Now()
		// 		}
		// 		_, err = o.Update(mmark)
		// 		if err != nil {
		// 			return err
		// 		}
		// 	}
		// }
	}
	return err
}

//删除价值结构
func DeleteAdminMerit(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	merit := AdminMerit{Id: id}
	var err error
	if o.Read(&merit) == nil {
		_, err = o.Delete(&merit)
		if err != nil {
			return err
		}
	}
	return err
}

//价值表带mark分值
type Merit struct {
	// gorm.Model
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	ParentId int64  `json:"parentid"`
	Mark     int    `json:"mark"`
	// Children []UserMerit
}

//取得所有的价值
func GetAllMerit() (merits []*Merit, err error) {
	db := GetDB()
	// 必须要写权select，坑爹啊 OR admin_merit.id=?
	err = db.Table("admin_merit").
		Select("admin_merit.id as id,admin_merit.title as title,admin_merit.parent_id as parent_id,admin_merit_mark.mark as mark").
		Joins("left JOIN admin_merit_mark on admin_merit_mark.merit_id = admin_merit.id").
		Scan(&merits).Error
	return merits, err
}

//****************IP******************
//添加ip地址段
func AddAdminIpsegment(title, startip, endip string, iprole int) (id int64, err error) {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{
		Title:   title,
		StartIp: startip,
		EndIp:   endip,
		Iprole:  iprole,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(ipsegment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改Ip地址段
func UpdateAdminIpsegment(cid int64, title, startip, endip string, iprole int) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		ipsegment.Title = title
		ipsegment.StartIp = startip
		ipsegment.EndIp = endip
		ipsegment.Iprole = iprole
		ipsegment.Updated = time.Now()
		_, err := o.Update(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminIpsegment(cid int64) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		_, err := o.Delete(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//查询所有Ip地址段
func GetAdminIpsegment() (ipsegments []*AdminIpsegment, err error) {
	o := orm.NewOrm()
	// ipsegments = make([]*AdminIpsegment, 0)

	qs := o.QueryTable("AdminIpsegment") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.All(&ipsegments)
	if err != nil {
		return nil, err
	}

	return ipsegments, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}

//添加
func AddAdminCategory(pid int64, title, code string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	// var category AdminCategory
	// if pid == "" {
	// 	category := &AdminCategory{
	// 		ParentId: 0,
	// 		Title:    title,
	// 		Code:     code,
	// 		Grade:    grade,
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	id, err = o.Insert(category)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }
	category := &AdminCategory{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Grade:    grade,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(category)
	if err != nil {
		return 0, err
	}
	// }
	return id, nil
}

//修改
func UpdateAdminCategory(cid int64, title, code string, grade int) error {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	category := &AdminCategory{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Grade = grade
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminCategory(cid int64) error {
	o := orm.NewOrm()
	category := &AdminCategory{Id: cid}
	if o.Read(category) == nil {
		_, err := o.Delete(category)
		if err != nil {
			return err
		}
	}
	// catalog := Catalog{Id: cidNum}
	// if o.Read(&catalog) == nil {
	// 	_, err = o.Delete(&catalog)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
func GetAdminCategory(pid int64) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	categories = make([]*AdminCategory, 0)

	qs := o.QueryTable("AdminCategory") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.Filter("parentid", pid).All(&categories) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}

	return categories, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}

//根据类别名字title查询所有下级分级category
func GetAdminCategoryTitle(title string) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")
	var cate AdminCategory
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := AdminCategory{Title: title}这句无效
	categories = make([]*AdminCategory, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}

//根据id查分级
func GetAdminCategorybyId(id int64) (category []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")

	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//添加ip地址段
// func AddAdminIpsegment(title, startip, endip string, iprole int) (id int64, err error) {
// 	o := orm.NewOrm()
// 	ipsegment := &AdminIpsegment{
// 		Title:   title,
// 		StartIp: startip,
// 		EndIp:   endip,
// 		Iprole:  iprole,
// 		Created: time.Now(),
// 		Updated: time.Now(),
// 	}
// 	id, err = o.Insert(ipsegment)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

//修改Ip地址段
// func UpdateAdminIpsegment(cid int64, title, startip, endip string, iprole int) error {
// 	o := orm.NewOrm()
// 	ipsegment := &AdminIpsegment{Id: cid}
// 	if o.Read(ipsegment) == nil {
// 		ipsegment.Title = title
// 		ipsegment.StartIp = startip
// 		ipsegment.EndIp = endip
// 		ipsegment.Iprole = iprole
// 		ipsegment.Updated = time.Now()
// 		_, err := o.Update(ipsegment)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//删除
// func DeleteAdminIpsegment(cid int64) error {
// 	o := orm.NewOrm()
// 	ipsegment := &AdminIpsegment{Id: cid}
// 	if o.Read(ipsegment) == nil {
// 		_, err := o.Delete(ipsegment)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//查询所有Ip地址段
// func GetAdminIpsegment() (ipsegments []*AdminIpsegment, err error) {
// 	o := orm.NewOrm()
// 	// ipsegments = make([]*AdminIpsegment, 0)

// 	qs := o.QueryTable("AdminIpsegment") //这个表名AchievementTopic需要用驼峰式，
// 	// if pid != "" {                      //如果给定父id则进行过滤
// 	//pid转成64为
// 	// pidNum, err := strconv.ParseInt(pid, 10, 64)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	_, err = qs.All(&ipsegments)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ipsegments, err
// 	// } else { //如果不给定父id（PID=0），则取所有一级
// 	// _, err = qs.Filter("parentid", 0).All(&categories)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return categories, err
// 	// }
// }

//********日历********
//添加
func AddAdminCalendar(title, content, color string, allday, public bool, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	calendar := &AdminCalendar{
		Title:   title,
		Content: content,
		Color:   color,
		Allday:  allday,
		Public:  public,
		// BColor:    color,
		Starttime: start,
		Endtime:   end,
	}
	id, err = o.Insert(calendar)
	if err != nil {
		return id, err
	}
	// }
	return id, err
}

//取所有——要修改为支持时间段的，比如某个月份
func GetAdminCalendar(start, end time.Time, public bool) (calendars []*AdminCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*AdminCalendar, 0)
	// qs := o.QueryTable("AdminCalendar")
	if public { //只取公开的
		_, err = qs.Filter("public", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//取所有——要修改为支持时间段的，比如某个月份
func SearchAdminCalendar(title string, public bool) (calendars []*AdminCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("title__contains", title).Or("content__contains", title)

	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	qs = qs.SetCond(cond1)

	calendars = make([]*AdminCalendar, 0)
	// qs := o.QueryTable("AdminCalendar")
	if public { //只取公开的
		_, err = qs.Filter("title", title).Filter("public", true).OrderBy("-Starttime").Distinct().All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").Distinct().All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//修改
func UpdateAdminCalendar(cid int64, title, content, color string, allday, public bool, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Title = title
		calendar.Content = content
		calendar.Color = color
		calendar.Allday = allday
		calendar.Public = public
		// calendar.BColor = color
		calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//拖曳
func DropAdminCalendar(cid int64, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Starttime = start
		calendar.Endtime = end
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//resize
func ResizeAdminCalendar(cid int64, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		// calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据id查询事件
func GetAdminCalendarbyid(id int64) (calendar AdminCalendar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	err = qs.Filter("id", id).One(&calendar)
	if err != nil {
		return calendar, err
	}
	return calendar, err
}

//删除事件
func DeleteAdminCalendar(cid int64) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		_, err := o.Delete(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//****项目同步ip****
//添加AdminSynchIp
func AddAdminSynchIp(pid int64, username, synchip, port string) (id int64, err error) {
	o := orm.NewOrm()
	adminsynchip := &AdminSynchIp{
		ParentId: pid,
		UserName: username,
		SynchIp:  synchip,
		Port:     port,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(adminsynchip)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateAdminSynchIp(cid int64, username, synchip, port string) error {
	o := orm.NewOrm()
	adminsynchip := &AdminSynchIp{Id: cid}
	if o.Read(adminsynchip) == nil {
		adminsynchip.UserName = username
		adminsynchip.SynchIp = synchip
		adminsynchip.Port = port
		adminsynchip.Updated = time.Now()
		_, err := o.Update(adminsynchip)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminSynchIp(cid int64) error {
	o := orm.NewOrm()
	synchip := &AdminSynchIp{Id: cid}
	if o.Read(synchip) == nil {
		_, err := o.Delete(synchip)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据父级id取得所有AdminSynchIp
func GetAdminSynchIp(pid int64) (synchips []*AdminSynchIp, err error) {
	o := orm.NewOrm()
	synchips = make([]*AdminSynchIp, 0)
	qs := o.QueryTable("AdminSynchIp")                 //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", pid).All(&synchips) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return synchips, err
}

//***科室结构********

//添加部门
// func AddAdminDepart(pid int64, title, code string) (id int64, err error) {
// 	o := orm.NewOrm()
// 	depart := &AdminDepartment{
// 		ParentId: pid,
// 		Title:    title,
// 		Code:     code,
// 		Created:  time.Now(),
// 		Updated:  time.Now(),
// 	}
// 	id, err = o.Insert(depart)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

//修改
// func UpdateAdminDepart(cid int64, title, code string) error {
// 	o := orm.NewOrm()
// 	category := &AdminDepartment{Id: cid}
// 	if o.Read(category) == nil {
// 		category.Title = title
// 		category.Code = code
// 		category.Updated = time.Now()
// 		_, err := o.Update(category)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//删除
// func DeleteAdminDepart(cid int64) error {
// 	o := orm.NewOrm()
// 	category := &AdminDepartment{Id: cid}
// 	if o.Read(category) == nil {
// 		_, err := o.Delete(category)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//根据部门id取得所有科室
//如果父级id为0，则取所有部门
// func GetAdminDepart(pid int64) (departs []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	departs = make([]*AdminDepartment, 0)
// 	qs := o.QueryTable("AdminDepartment")
// 	_, err = qs.Filter("parentid", pid).All(&departs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return departs, err
// }

//根据父级id取得所有
//如果父级id为空，则取所有一级category
// func GetAdminCategory(pid int64) (categories []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	categories = make([]*AdminDepartment, 0)
// 	qs := o.QueryTable("AdminDepartment")
// 	_, err = qs.Filter("parentid", pid).All(&categories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, err
// }
//根据部门名字title返回自身
// func GetAdminDepartName(title string) (AdminDepartment, error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("AdminDepartment")
// 	var cate AdminDepartment
// 	err := qs.Filter("title", title).One(&cate)
// 	if err != nil {
// 		return cate, err
// 	}
// 	return cate, err
// }

//根据部门名字title查询所有下级科室category
// func GetAdminDepartTitle(title string) (categories []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("AdminDepartment")
// 	var cate AdminDepartment
// 	err = qs.Filter("title", title).One(&cate)
// 	categories = make([]*AdminDepartment, 0)
// 	_, err = qs.Filter("parentid", cate.Id).All(&categories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, err
// }

//根据id查科室
// func GetAdminDepartbyId(id int64) (category AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("AdminDepartment")
// 	err = qs.Filter("id", id).One(&category)
// 	if err != nil {
// 		return category, err
// 	}
// 	return category, err
// }

//由分院id和科室 名称取得科室
// func GetAdminDepartbyidtitle(id int64, title string) (*AdminDepartment, error) {
// 	o := orm.NewOrm()
// 	// cate := &Category{Id: id}
// 	category := new(AdminDepartment)
// 	qs := o.QueryTable("AdminDepartment")
// 	err := qs.Filter("parentid", id).Filter("title", title).One(category)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return category, err
// }

//添加轮播动画
func AddAdminCarousel(title, url string) (id int64, err error) {
	o := orm.NewOrm()
	carousel := &AdminCarousel{
		Title: title,
		// Path:    code,
		Url:     url,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(carousel)
	if err != nil {
		return 0, err
	}
	// }
	return id, nil
}

func GetAdminCarousel() (carousels []*AdminCarousel, err error) {
	o := orm.NewOrm()
	carousels = make([]*AdminCarousel, 0)
	qs := o.QueryTable("AdminCarousel")
	_, err = qs.OrderBy("-created").All(&carousels)
	if err != nil {
		return nil, err
	}
	return carousels, err
}
