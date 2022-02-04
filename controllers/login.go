package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
	"github.com/3xxx/meritms/controllers/utils"
	"github.com/3xxx/meritms/models"
	"strconv"
	// "github.com/astaxie/beego/session"
	"encoding/json"
	"net/http"
)

// CMSWX login API
type LoginController struct {
	beego.Controller
}

// func (c *LoginController) Get() {
// 	isExit := c.Input().Get("exit") == "true"
// 	// secofficeshow?secid=1643&level=3&key=modify
// 	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
// 	url2 := c.Input().Get("level")
// 	url3 := c.Input().Get("key")
// 	var url string
// 	if url2 == "" {
// 		url = url1
// 	} else {
// 		url = url1 + "&level=" + url2 + "&key=" + url3
// 	}
// 	c.Data["Url"] = url
// 	if isExit {
// 		// c.Ctx.SetCookie("uname", "", -1, "/")
// 		// c.Ctx.SetCookie("pwd", "", -1, "/")
// 		// c.DelSession("gosessionid")
// 		// c.DelSession("uname")//这个不行
// 		// c.Destroy/Session()
// 		// c.Ctx.Input.CruSession.Delete("gosessionid")这句与上面一句重复
// 		// c.Ctx.Input.CruSession.Flush()
// 		// beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
// 		v := c.GetSession("uname")
// 		// islogin := false
// 		if v != nil {
// 			//删除指定的session
// 			c.DelSession("uname")
// 			//销毁全部的session
// 			c.DestroySession()
// 		}
// 		// sess.Flush()//这个不灵
// 		c.Redirect("/", 301)
// 		return
// 	}
// 	c.TplName = "login.tpl"
// }

// func (c *LoginController) Loginerr() {
// 	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
// 	url2 := c.Input().Get("level")
// 	url3 := c.Input().Get("key")
// 	var url string
// 	if url2 == "" {
// 		url = url1
// 	} else {
// 		url = url1 + "&level=" + url2 + "&key=" + url3
// 	}
// 	// port := strconv.Itoa(c.Ctx.Input.Port())
// 	// url := c.Ctx.Input.Site() + ":" + port + c.Ctx.Request.URL.String()
// 	c.Data["Url"] = url
// 	// beego.Info(url)
// 	c.TplName = "loginerr.tpl"
// }

//登录页面
func (c *LoginController) Login() {
	c.TplName = "login.tpl"
}

//login页面输入用户名和密码后登陆提交
func (c *LoginController) Post() {
	// uname := c.Input().Get("uname")
	// url := c.Input().Get("returnUrl")
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("level")
	url3 := c.Input().Get("key")
	var url string
	if url2 == "" && url1 != "" {
		url = url1
	} else if url2 != "" {
		url = url1 + "&level=" + url2 + "&key=" + url3
	} else {
		url = c.Input().Get("referrer")
	}
	var user models.User
	user.Username = c.Input().Get("uname")
	Pwd1 := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {
		c.SetSession("uname", user.Username)
		c.SetSession("pwd", user.Password)
		utils.FileLogs.Info(user.Username + " " + "login" + " 成功")
		User, err := models.GetUserByUsername(user.Username)
		if err != nil {
			beego.Error(err)
			utils.FileLogs.Error(user.Username + " 查询用户 " + err.Error())
		}
		if User.Ip == "" {
			err = models.UpdateUser(User.Id, "Ip", c.Ctx.Input.IP())
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 添加用户ip " + err.Error())
			}
		} else {
			//更新user表的lastlogintime
			err = models.UpdateUserlastlogintime(user.Username)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 更新用户登录时间 " + err.Error())
			}
		}
		if url != "" {
			c.Redirect(url, 301)
		} else {
			c.Redirect("/", 301)
		}
	} else {
		c.Redirect("/loginerr?url="+url, 302)
	}
	return
}

//login弹框输入用户名和密码后登陆提交
func (c *LoginController) LoginPost() {
	var user models.User
	user.Username = c.Input().Get("uname")
	// uname := c.GetString("uname")
	Pwd1 := c.GetString("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"
	islogin := 0
	// maxAge := 0
	// if autoLogin {
	// 	maxAge = 1<<31 - 1
	// }
	// c.Ctx.SetCookie("uname", uname, maxAge, "/")
	// c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	beego.Info(user.Password)
	err := models.ValidateUser(user)
	if err == nil {
		c.SetSession("uname", user.Username)
		c.SetSession("pwd", user.Password)
		utils.FileLogs.Info(user.Username + " " + "login" + " 成功")
		User, err := models.GetUserByUsername(user.Username)
		if err != nil {
			beego.Error(err)
			utils.FileLogs.Error(user.Username + " 查询用户 " + err.Error())
		}
		if User.Ip == "" {
			err = models.UpdateUser(User.Id, "Ip", c.Ctx.Input.IP())
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 添加用户ip " + err.Error())
			}
		} else {
			//更新user表的lastlogintime
			err = models.UpdateUserlastlogintime(user.Username)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 更新用户登录时间 " + err.Error())
			}
		}
	} else {
		islogin = 1
	}
	// if name == "admin" && pwd == "123456" {
	// 	c.SetSession("loginuser", "adminuser")
	// 	fmt.Println("当前的session:")
	// 	fmt.Println(c.CruSession)
	c.Data["json"] = map[string]interface{}{"islogin": islogin}
	c.ServeJSON()
}

//退出登录
func (c *LoginController) Logout() {
	v := c.GetSession("uname")
	islogin := false
	if v != nil {
		//删除指定的session
		c.DelSession("uname")
		//销毁全部的session
		c.DestroySession()
		islogin = true
	}
	c.Data["json"] = map[string]interface{}{"islogin": islogin}
	c.ServeJSON()
}

//作废20180915
func (c *LoginController) Loginerr() {
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("level")
	url3 := c.Input().Get("key")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&level=" + url2 + "&key=" + url3
	}
	// port := strconv.Itoa(c.Ctx.Input.Port())
	// url := c.Ctx.Input.Site() + ":" + port + c.Ctx.Request.URL.String()
	c.Data["Url"] = url
	// beego.Info(url)
	c.TplName = "loginerr.tpl"
}

// @Title post wx login
// @Description post wx login
// @Param id path string  true "The id of wx"
// @Param code path string  true "The jscode of wxuser"
// @Success 200 {object} success
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /wxlogin/:id [get]
//微信小程序访问微信服务器获取用户信息
func (c *LoginController) WxLogin() {
	id := c.Ctx.Input.Param(":id")
	JSCODE := c.Input().Get("code")
	// beego.Info(JSCODE)
	var APPID, SECRET string
	if id == "1" {
		APPID = "wx7f77b90a1a891d93"
		SECRET = "f58ca4f28cbb52ccd805d66118060449"
	} else if id == "2" {
		APPID = beego.AppConfig.String("wxAPPID2")
		SECRET = beego.AppConfig.String("wxSECRET2")
	} else if id == "3" {
		APPID = beego.AppConfig.String("wxAPPID3")
		SECRET = beego.AppConfig.String("wxSECRET3")
	}

	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		beego.Error(err)
		// return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		// return
	}
	// beego.Info(data)
	if _, ok := data["session_key"]; !ok {
		errcode := data["errcode"]
		errmsg := data["errmsg"].(string)
		// return
		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		// c.ServeJSON()
	} else {
		var openID string
		var sessionKey string
		// var unionId string
		openID = data["openid"].(string)
		sessionKey = data["session_key"].(string)
		// unionId = data["unionid"].(string)
		// beego.Info(openID)
		// beego.Info(sessionKey)
		// beego.Info(unionId)
		//如果数据库存在记录，则存入session？
		//上传文档的时候，检查session？
		c.SetSession("uname", openID)
		c.SetSession("pwd", sessionKey)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "success", "data": "3rd_session"}
		c.ServeJSON()
	}
}

// [login.go:224] 0716J5410OfFLF1Daw610f6a4106J54e
// [login.go:247] map[session_key:3NaIB1t/AOjCQKitWx1fr
// Q== openid:oRgfy5MQlRRxyyNrENpZWnhniO-I]
// 2018/09/09 18:57:04.791 [C] [asm_amd64.s:509] the request url is  /wx/wxlogin
// 2018/09/09 18:57:04.807 [C] [asm_amd64.s:509] Handler crashed with error interfa
// ce conversion: interface {} is nil, not string
// 2018/09/09 18:57:04.807 [C] [asm_amd64.s:509] D:/gowork/src/github.com/3xxx/engi
// neercms/controllers/login.go:260
//判断用户是否登录
func checkAccount(ctx *context.Context) bool {
	var user models.User
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	//但是我还是建议大家采用 SetSession、GetSession、DelSession 三个方法来操作，避免自己在操作的过程中资源没释放的问题
	// sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	// defer sess.SessionRelease(ctx.ResponseWriter)
	v := ctx.Input.CruSession.Get("uname")
	if v == nil {
		return false
		//     this.SetSession("asta", int(1))
		//     this.Data["num"] = 0
	} else {
		//     this.SetSession("asta", v.(int)+1)
		//     this.Data["num"] = v.(int)
		user.Username = v.(string)
		v = ctx.Input.CruSession.Get("pwd")
		user.Password = v.(string) //ck.Value
		err := models.ValidateUser(user)
		if err == nil {
			return true
		} else {
			return false
		}
	}
}

func checkRole(ctx *context.Context) (role string, err error) { //这里返回用户的role
	v := ctx.Input.CruSession.Get("uname")
	var user models.User
	user.Username = v.(string) //ck.Value
	user, err = models.GetUserByUsername(user.Username)
	if err != nil {
		beego.Error(err)
	}
	return user.Role, err
}

// type Session struct {
// 	Session int
// }
// type Login struct {
// 	UserName string
// 	Password string
// }

//用户登录，则role是1则是admin，其余没有意义
//ip区段，casbin中表示，比如9楼ip区段作为用户，赋予了角色，这个角色具有访问项目目录权限
func checkprodRole(ctx *context.Context) (uname, role string, uid int64, isadmin, islogin bool) {
	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	// beego.Info(v)                          //qin.xc
	var userrole string
	var user models.User
	var err error
	var iprole int
	if v != nil { //如果登录了
		islogin = true
		uname = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		} else {
			uid = user.Id
			if user.Role == "0" {
				isadmin = false
				userrole = "4"
			} else if user.Role == "1" {
				isadmin = true
				userrole = user.Role
			} else {
				isadmin = false
				userrole = user.Role
			}
		}
	} else { //如果没登录,查询ip对应的用户
		islogin = false
		isadmin = false
		uid = 0
		uname = ctx.Input.IP()
		// beego.Info(uname)
		user, err = models.GetUserByIp(uname)
		beego.Info(user)
		if err != nil { //如果查不到，则用户名就是ip，role再根据ip地址段权限查询
			// beego.Error(err)
			iprole = Getiprole(ctx.Input.IP()) //查不到，则是5——这个应该取消，采用casbin里的ip区段
			beego.Info(iprole)
			userrole = strconv.Itoa(iprole)
		} else { //如果查到，则role和用户名
			if user.Role == "1" {
				isadmin = true
			}
			uid = user.Id
			userrole = user.Role
			uname = user.Username
			islogin = true
		}
	}
	return uname, userrole, uid, isadmin, islogin
}

// func checkRole(ctx *context.Context) (roles []*models.Role, err error) {
// 	ck, err := ctx.Request.Cookie("uname")
// 	if err != nil {
// 		return roles, err
// 	}
// 	var user models.User
// 	user.Username = ck.Value

// 	roles, _ = models.GetRoleByUsername(user.Username)
// 	if err == nil {
// 		return roles, err
// 	} else {
// 		return roles, err
// 	}
// }

// func GetRoleByUserId(userid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
// 	roles = make([]*Role, 0)
// 	o := orm.NewOrm()
// 	// role := new(Role)
// 	count, _ = o.QueryTable("role").Filter("Users__User__Id", userid).All(&roles)
// 	return roles, count
// 	// 通过 post title 查询这个 post 有哪些 tag
// 	// var tags []*Tag
// 	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

// }
