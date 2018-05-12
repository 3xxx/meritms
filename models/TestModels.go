package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id      int64
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id    int64
	Age   int16
	Users *Users `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

// Post 和 User 是 ManyToOne 关系，也就是 ForeignKey 为 User
type Post struct {
	Id    int64
	Title string
	Users *Users `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int64
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users), new(Post), new(Profile), new(Tag))
}

func GetUser1(id int64) (users *Users, err error) {
	o := orm.NewOrm()
	users = &Users{Id: id}
	o.Read(users)
	if users.Profile != nil {
		// 	o.Read(user1.Profile)
		return users, err
	}

	// 直接关联查询：
	// user1 = &User1{}
	// o.QueryTable("User1").Filter("Id", 1).RelatedSel().One(user1)
	// 自动查询到 Profile
	// fmt.Println(user1.Profile)
	// 因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。
	// fmt.Println(user1.Profile.User1)
	return users, err
}

func Getprofile(id int64) (profile *Profile, err error) {
	o := orm.NewOrm()
	// 通过 User 反向查询 Profile：
	// var profile Profile
	err = o.QueryTable("Profile").Filter("Users__Id", id).One(&profile)
	// if err == nil {
	// 	fmt.Println(profile)
	// }
	return profile, err
}

func GetPost(id int64) (posts []*Post, err error) {
	o := orm.NewOrm()
	// var posts []*Post
	_, err = o.QueryTable("post").Filter("Users__Id", 1).RelatedSel().All(&posts)
	// if err == nil {
	// fmt.Printf("%d posts read\n", num)
	// for _, post := range posts {
	// 	fmt.Printf("Id: %d, UserName: %d, Title: %s\n", post.Id, post.User1.Name, post.Title)
	// }
	// }
	return posts, err
}

func GetUser(title string) (users *Users, err error) {
	o := orm.NewOrm()
	// 根据 Post.Title 查询对应的 User：
	// RegisterModel 时，ORM 也会自动建立 User 中 Post 的反向关系，所以可以直接进行查询
	// var user1 User1
	err = o.QueryTable("users").Filter("Post__Title", title).Limit(1).One(&users)
	// if err == nil {
	// 	fmt.Printf(user1)
	// }
	return users, err
}

// Post 和 Tag 是 ManyToMany 关系
// 设置 rel(m2m) 以后，ORM 会自动创建中间表
// type Post struct {
//     Id    int
//     Title string
//     User  *User  `orm:"rel(fk)"`
//     Tags  []*Tag `orm:"rel(m2m)"`
// }
// type Tag struct {
//     Id    int
//     Name  string
//     Posts []*Post `orm:"reverse(many)"`
// }
// func GetPosts(id int64) (posts []*Post, err error) {
// 	// 一条 Post 纪录可能对应不同的 Tag 纪录,一条 Tag 纪录可能对应不同的 Post 纪录，所以 Post 和 Tag 属于多对多关系,通过 tag name 查询哪些 post 使用了这个 tag
// 	// var posts []*Post
// 	num, err := dORM.QueryTable("post").Filter("Tags__Tag__Name", "golang").All(&posts)
// 	// 通过 post title 查询这个 post 有哪些 tag
// 	var tags []*Tag
// 	num, err = dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)
// }
