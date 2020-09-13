//成果里的附件操作
package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type AttachController struct {
	beego.Controller
}

//目前有文章中的图片、成果中文档的预览、onlyoffice中的文档协作、pdf中的附件路径等均采用绝对路径型式
//文章中的附件呢？
//default中的pdf页面中的{{.pdflink}}，绝对路径
// type Session struct {
// 	Session int
// }
//attachment/路径/附件名称
func (c *AttachController) Attachment() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(c.Ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	if strings.Contains(filePath, "?hotqinsessionid=") {
		filePathtemp := strings.Split(filePath, "?")
		filePath = filePathtemp[0]
		beego.Info(filePath)
	}
	// beego.Info(filePath)
	fileext := path.Ext(filePath)
	// filepath1 := path.Dir(filePath)
	// array := strings.Split(filepath1, "/")

	switch fileext {
	case ".JPG", ".jpg", ".png", ".PNG", ".bmp", ".BMP", ".mp4", ".MP4":
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	default:
		// beego.Info(useridstring)
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//首页轮播图片给予任何权限
func (c *AttachController) GetCarousel() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
}

//返回文件大小
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
