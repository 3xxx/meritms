package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	// "fmt"
	"time"
)

type ProjGant struct {
	Id          int64     `form:"-"`
	Code        string    `orm:"null"` //编号
	Title       string    `orm:"null"` //项目
	DesignStage string    `orm:"null"` //-阶段
	Section     string    `orm:"null"` //专业
	Label       string    `orm:"null"` //标签
	Desc        string    `orm:"null"`
	CustomClass string    `orm:"null"`
	DataObj     string    //['ha','ha2']
	Starttime   time.Time `orm:"not null;type(datetime)"`
	Endtime     time.Time `orm:"null;type(datetime)"`
	Show        bool
	Created     time.Time `orm:"null;index","auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"null;index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(ProjGant))
}

//添加项目进度
func AddProjGant(code, title, designstage, section, label, desc, customclass, dataobj string, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	//关闭写同步
	// o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	// var ProjGant ProjGant
	// if pid == "" {
	projgant := &ProjGant{
		Code:        code,
		Title:       title,
		DesignStage: designstage,
		Section:     section,
		Label:       label,
		Desc:        desc,
		CustomClass: customclass,
		DataObj:     dataobj,
		Starttime:   start,
		Endtime:     end,
		Show:        true,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	id, err = o.Insert(projgant)
	if err != nil {
		return 0, err
	}
	// } else {

	// }
	return id, nil
}

//修改——还没改
func UpdateProjGant(cid int64, code, title, label string) error {
	o := orm.NewOrm()
	project := &ProjGant{Id: cid}
	if o.Read(project) == nil {
		project.Code = code
		project.Title = title
		project.Label = label
		project.Updated = time.Now()
		_, err := o.Update(project)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteProjGant(id int64) error {
	o := orm.NewOrm()
	project := &ProjGant{Id: id}
	if o.Read(project) == nil {
		_, err := o.Delete(project)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func CloseProjGant(id int64) error {
	o := orm.NewOrm()
	project := &ProjGant{Id: id}
	if o.Read(project) == nil {
		_, err := o.Delete(project)
		if err != nil {
			return err
		}
	}
	return nil
}

//取得所有项目进度，按结束时间早到晚排列
func GetProjGants() (projgant []*ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant") //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("show", true).OrderBy("Endtime").All(&projgant)
	if err != nil {
		return projgant, err
	}
	return projgant, err
}

//根据id取得项目目录
func GetProjGant(id int64) (projgant ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&projgant)
	if err != nil {
		return projgant, err
	}
	return projgant, err
}

//根据名字title查询到项目目录
func GetProjGantTitle(title string) (projgant ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant")
	// var cate ProjGant
	err = qs.Filter("title", title).One(&projgant)
	// if pid != "" {
	// cate := ProjGant{Title: title}这句无效
	// categories = make([]*ProjGant, 0)
	// _, err = qs.Filter("parentid", cate.Id).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	return projgant, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}
