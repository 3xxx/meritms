package controllers

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/pborman/uuid"
	// "image/png"
	// "io"
	// "log"
	// "net/http"
	"os"
	"path"
	// "hydrocms/models"
	// "encoding/base64"
	"github.com/3xxx/meritms/models"
	// "io/ioutil"
	// "regexp"
	"strconv"
	// "strings"
	"time"
)

// CMSWX froala API
type FroalaController struct {
	beego.Controller
}

type UploadimgFroala struct {
	url      string
	title    string
	original string
	state    string
	// "url": fmt.Sprintf("/static/upload/%s", filename),
	// "title": "demo.jpg",
	// "original": header.Filename,
	// "state": "SUCCESS"
}

// @Title post merit user img
// @Description post user merit img
// @Param meritid query string true "The id of adminmerit"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadmeritimg [post]
// merit添加用户价值图片上传
func (c *FroalaController) UploadMeritImg() {
	var user models.User
	var err error
	// username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	v := c.Ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	beego.Info(v.(string))
	// var userid, roleid, userrole string
	if v != nil { //如果登录了
		uname := v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	//获取用户昵称user.Nickname
	//获取用户科室和部门department/secoffice
	//获取用户价值id和价值名称
	mid := c.Input().Get("meritid")
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	merit, err := models.GetAdminMeritbyId(midNum)
	if err != nil {
		beego.Error(err)
	}
	parentmerit, err := models.GetAdminMeritbyId(merit.ParentId)
	if err != nil {
		beego.Error(err)
	}
	//建立这个目录
	err = os.MkdirAll("./attachment/merit/"+user.Department+"/"+user.Secoffice+"/"+user.Nickname+"/"+parentmerit.Title+"/"+merit.Title, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		path = "./attachment/merit/" + user.Department + "/" + user.Secoffice + "/" + user.Nickname + "/" + parentmerit.Title + "/" + merit.Title + "/" + newname
		Url := "/attachment/merit/" + user.Department + "/" + user.Secoffice + "/" + user.Nickname + "/" + parentmerit.Title + "/" + merit.Title + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "photo": "", "title": "", "original": ""}
			c.ServeJSON()
		} else {
			// filesize, _ = FileSize(path)
			// filesize = filesize / 1000.0
			// _, err = models.AddUserAvator(user.Id, Url+newname)
			// if err != nil {
			// 	beego.Error(err)
			// }"link": Url + newname,
			// c.Ctx.WriteString(wxsite + Url + newname)
			c.Data["json"] = map[string]interface{}{"errNo": 0, "state": "SUCCESS", "link": Url + newname, "title": newname, "original": Url + newname}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"errNo": 0, "state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// 添加wiki里的图片上传
func (c *FroalaController) UploadWikiImg() {
	//保存上传的图片
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	// var filesize int64
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
	Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	// filesize, _ = FileSize(path1)
	// filesize = filesize / 1000.0
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
	c.ServeJSON()
}

//下面这个保留
// func (c *FroalaController) UploadImg() { //对应这个路由 beego.Router("/controller", &controllers.FroalaController{}, "post:UploadImage")
// 	file, header, err := c.GetFile("file") // r.FormFile("upfile")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
// 	err = os.MkdirAll(path.Join("static", "upload"), 0775)
// 	if err != nil {
// 		panic(err)
// 	}
// 	outFile, err := os.Create(path.Join("static", "upload", filename))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer outFile.Close()
// 	io.Copy(outFile, file)
// 	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "/static/upload/" + filename, "title": "111", "original": "demo.jpg"}
// 	c.ServeJSON()
// }
