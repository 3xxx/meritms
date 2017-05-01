//成果类型
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "math"
	"strconv"
	"time"
)

type AdminAchievcategory struct {
	Id         int64
	Category   string  //成果类型及单位
	Rationum   float64 //折标系数
	Ismaterial bool    //是否是实物工作量
	// Data     time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index;auto_now_add;type(datetime)"`
}

func init() {
	// orm.RegisterModel(new(AdminAchievcategory))
}

func SaveAchievcategory(category AdminAchievcategory) (cid int64, err error) {
	//重复提交问题
	category1 := AdminAchievcategory{Category: category.Category}
	// 	user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	o := orm.NewOrm()
	err1 := o.Read(&category1, "Category")
	if err1 == orm.ErrNoRows {
		cid, err = o.Insert(&category)
		return cid, err
	} else {
		return 0, err1
	}

}

func GetAchievcategories() (categories []*AdminAchievcategory, err error) {
	categories = make([]*AdminAchievcategory, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("AdminAchievcategory")
	_, err = qs.All(&categories)
	return categories, err
}

//根据成果类型，查系数等特性
func GetAchcatebycate(category string) (ratio AdminAchievcategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminAchievcategory") //这里要大写吗？
	// var ratio1 Ratio
	err = qs.Filter("category", category).One(&ratio)
	if err != nil {
		return ratio, err
	}
	return ratio, err

}

func UpdateAchievcategory(cid int64, ratio1 AdminAchievcategory) error {
	o := orm.NewOrm()
	ratio := &AdminAchievcategory{Id: cid}
	var err error
	if o.Read(ratio) == nil {
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
		ratio.Category = ratio1.Category
		// ratio.Unit = ratio1.Unit
		ratio.Ismaterial = ratio1.Ismaterial
		ratio.Rationum = ratio1.Rationum
		ratio.Updated = ratio1.Updated
		// _, err = o.Update(&ratio, "ProjectName", "DesignStage", "Section", "Tnumber", "Name", "Category", "Page", "Count", "Drawn", "Designd", "Checked", "Examined", "Data", "Updated", "Author")
		_, err = o.Update(ratio, "Category", "Rationum", "Ismaterial", "Updated") //这里不能用&ratio
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteAchievcategory(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	var err error
	ratio := AdminAchievcategory{Id: id}
	if o.Read(&ratio) == nil {
		_, err = o.Delete(&ratio)
		if err != nil {
			return err
		}
	}
	return err
}

func GetAchievcategory(tid string) (*AdminAchievcategory, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	ratio := new(AdminAchievcategory)
	qs := o.QueryTable("AdminAchievcategory")
	err = qs.Filter("id", tidNum).One(ratio)
	if err != nil {
		return nil, err
	}
	return ratio, err
}
