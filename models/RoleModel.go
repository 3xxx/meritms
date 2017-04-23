package models

import (
	"errors"
	"log"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//角色表
type Role struct {
	Id     int64
	Title  string `orm:"size(100)" form:"Title"  valid:"Required"`
	Name   string `orm:"size(100)" form:"Name"  valid:"Required"`
	Remark string `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status int    `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	// Users  []*User `orm:"reverse(many)"`
	// Node   []*Node `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Role)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

// func (r *Role) TableName() string {
// 	return beego.AppConfig.String("rbac_role_table")
// }

// func init() {
// 	orm.RegisterModel(new(Role))
// }

func checkRole(g *Role) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&g)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//get role list
func GetRolelist(page int64, page_size int64, sort string) (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&roles)
	count, _ = qs.Count()
	return roles, count
}

func AddRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := new(Role)
	role.Title = r.Title
	role.Name = r.Name
	role.Remark = r.Remark
	role.Status = r.Status

	id, err := o.Insert(role)
	return id, err
}

func UpdateRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := make(orm.Params)
	if len(r.Title) > 0 {
		role["Title"] = r.Title
	}
	if len(r.Name) > 0 {
		role["Name"] = r.Name
	}
	if len(r.Remark) > 0 {
		role["Remark"] = r.Remark
	}
	if r.Status != 0 {
		role["Status"] = r.Status
	}
	if len(role) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Role
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(role)
	return num, err
}

func DelRoleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Role{Id: Id})
	return status, err
}

// func GetNodelistByRoleId(Id int64) (nodes []orm.Params, count int64) {
// 	o := orm.NewOrm()
// 	node := new(Node)
// 	count, _ = o.QueryTable(node).Filter("Role__Role__Id", Id).Values(&nodes)
// 	return nodes, count
// }

// func DelGroupNode(roleid int64, groupid int64) error {
// 	var nodes []*Node
// 	var node Node
// 	role := Role{Id: roleid}
// 	o := orm.NewOrm()
// 	num, err := o.QueryTable(node).Filter("Group", groupid).RelatedSel().All(&nodes)
// 	if err != nil {
// 		return err
// 	}
// 	if num < 1 {
// 		return nil
// 	}
// 	for _, n := range nodes {
// 		m2m := o.QueryM2M(n, "Role")
// 		_, err1 := m2m.Remove(&role)
// 		if err1 != nil {
// 			return err1
// 		}
// 	}
// 	return nil
// }
// func AddRoleNode(roleid int64, nodeid int64) (int64, error) {
// 	o := orm.NewOrm()
// 	role := Role{Id: roleid}
// 	node := Node{Id: nodeid}
// 	m2m := o.QueryM2M(&node, "Role")
// 	num, err := m2m.Add(&role)
// 	return num, err
// }

func DelUserRole(roleid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_roles").Filter("role_id", roleid).Delete()
	return err
}

func AddRoleUser(roleid, userid int64) (num int64, err error) {
	//要判断是否已经存在，如果存在，则update
	o := orm.NewOrm()
	user := User{Id: userid}
	// var role1 Role
	// err = o.QueryTable("user_roles").Filter("user_id", userid).One(&role1)
	m2m := o.QueryM2M(&user, "Roles")
	if m2m.Exist(&User{Id: userid}) {
		// 	// fmt.Println("Tag Exist")
		num, err = m2m.Remove(&User{Id: userid})
		role := Role{Id: roleid}
		// 	user := User{Id: userid}
		num, err = m2m.Add(&role)
	} else {
		role := Role{Id: roleid}
		// 	user := User{Id: userid}
		num, err = m2m.Add(&role)
	}

	// user1 := new(User)
	// _, err := o.QueryTable("user_roles").Filter("user_id", userid).All(&user1)
	// if err == orm.ErrNoRows {
	// o := orm.NewOrm()
	// role := Role{Id: roleid}
	// user := User{Id: userid}
	// m2m := o.QueryM2M(&user, "Roles")
	// num, err := m2m.Add(&role)
	return num, err
	// } else {
	// }
}

func UpdateRoleUser(roleid1, roleid2, userid int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleid1}
	user := User{Id: userid}
	m2m := o.QueryM2M(&user, "Roles")
	num, err := m2m.Remove(&role)
	role2 := Role{Id: roleid2}
	_, err = m2m.Add(&role2)
	return num, err
}

func GetUserByRoleId(roleid int64) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	count, _ = o.QueryTable(user).Filter("Role__Role__Id", roleid).Values(&users)
	return users, count
}

//根据role等级取得roleid
func GetRoleIdbyTitle(roletitle string) (roleid int64, err error) {
	o := orm.NewOrm()
	var role Role
	err = o.QueryTable("role").Filter("title", roletitle).One(&role, "Id")
	if err != nil { //如果不存在这个权限
		return 0, err
	}
	return role.Id, nil
}

// var user User
// err := o.QueryTable("user").Filter("name", "slene").One(&user)
// if err == orm.ErrMultiRows {
//     // 多条的时候报错
//     fmt.Printf("Returned Multi Rows Not One")
// }
// if err == orm.ErrNoRows {
//     // 没有找到记录
//     fmt.Printf("Not row found")
// }
// 可以指定返回的字段：

// // 只返回 Id 和 Title
// var post Post
// o.QueryTable("post").Filter("Content__istartswith", "prefix string").One(&post, "Id", "Title")

// func AccessList(uid int64) (list []orm.Params, err error) {
// 	var roles []orm.Params
// 	o := orm.NewOrm()
// 	role := new(Role)
// 	_, err = o.QueryTable(role).Filter("User__User__Id", uid).Values(&roles)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var nodes []orm.Params
// 	node := new(Node)
// 	for _, r := range roles {
// 		_, err := o.QueryTable(node).Filter("Role__Role__Id", r["Id"]).Values(&nodes)
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, n := range nodes {
// 			list = append(list, n)
// 		}
// 	}
// 	return list, nil
// }
