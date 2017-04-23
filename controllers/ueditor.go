package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	// "net/http"
	// "merit/models" //这里开始直接拷贝过来quick/models,程序总是提示出错：<orm.RegisterModel> table name `category` repeat register, must be unique
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UeditorController struct {
	beego.Controller
}

type UploadimageUE struct {
	url      string
	title    string
	original string
	state    string
}

func (c *UeditorController) ControllerUE() {
	op := c.Input().Get("action")
	key := c.Input().Get("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
	switch op {
	case "config": //这里还是要优化成conf/config.json
		file, err := os.Open("conf/config.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		fd, err := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/") //参考php的$CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		//将php中的正则移植到go中，需要将/ \/\*[\s\S]+?\*\/  /去掉前后的/，然后将\改成2个\\
		//参考//去除所有尖括号内的HTML代码，并换成换行符
		// re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		// src = re.ReplaceAllString(src, "\n")
		//当把<和>换成/*和*\时，斜杠/和*之间加双斜杠\\才行。
		src = re.ReplaceAllString(src, "")
		tt := []byte(src)

		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码
		c.Data["json"] = r
		c.ServeJSON()

	case "uploadimage", "uploadfile", "uploadvideo":

		if key == "diary" { //添加文章
			// categoryid := c.Input().Get("categoryid")
			//保存上传的图片
			_, h, err := c.GetFile("upfile")
			if err != nil {
				beego.Error(err)
			}
			// idNum, err := strconv.ParseInt(categoryid, 10, 64)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// category1, err := models.GetCategory(idNum) //2016-3-5这里修改为int64
			// if err != nil {
			// 	beego.Error(err)
			// 	return
			// }
			// var filesize int64
			path1 := h.Filename                 // 2016-3-5这里修改category1.DiskDirectory +
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			// filesize, _ = FileSize(path1)
			// filesize = filesize / 1000.0
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": h.Filename, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()
		} else {
			number := c.Input().Get("number")
			name := c.Input().Get("name")
			err := os.MkdirAll(".\\attachment\\"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}

			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			_, h, err := c.GetFile("upfile")
			// beego.Info(h)
			if err != nil {
				beego.Error(err)
			}

			path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			//如果扩展名为jpg
			// if strings.ToLower(path.Ext(h.Filename)) == ".jpg" {
			// }
			//如果包含jpg，则进行压缩——压缩导致UEditor里显示尺寸过大。
			// if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
			// 	// 随机名称
			// 	// to := path + random_name() + ".jpg"
			// 	origin := path1 //path + file.Name()
			// 	fmt.Println("正在处理" + origin + ">>>" + origin)
			// 	cmd_resize(origin, 2048, 0, origin)
			// 	//defer os.Remove(origin)//删除原文件
			// }
			// filesize, _ = FileSize(path1)
			// filesize = filesize / 1000.0
			// route = "/attachment/" + number + name + "/" + h.Filename
			// outFile, err := os.Create(path.Join(".\\attachment\\"+number+name+"\\", filename))
			// if err != nil {
			// 	beego.Error(err)
			// }
			// defer outFile.Close()
			// io.Copy(outFile, file)
			// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + filename, "title": filename, "original": filename}
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/attachment/" + number + name + "/" + h.Filename, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()

		}
	case "uploadscrawl":
		number := c.Input().Get("number")

		name := c.Input().Get("name")
		err := os.MkdirAll(".\\attachment\\"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		path1 := ".\\attachment\\" + number + name + "\\"
		//保存上传的图片
		//upfile为base64格式文件，转成图片保存
		ww := c.Input().Get("upfile")
		ddd, _ := base64.StdEncoding.DecodeString(ww)           //成图片文件并把文件写入到buffer
		newname := strconv.FormatInt(time.Now().Unix(), 10)     // + "_" + filename
		err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err != nil {
			beego.Error(err)
		}
		// var filesize int64
		// filesize, _ = FileSize(path1)
		// filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{
			"state":    "SUCCESS",
			"url":      "/attachment/" + number + name + "/" + newname + ".jpg",
			"title":    newname + ".jpg",
			"original": newname + ".jpg",
		}
		c.ServeJSON()
	case "listimage":
		type List struct {
			Url string `json:"url"`
			// Source string
			// State  string
		}

		type Listimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
			Start int    `json:"start"`
			Total int    `json:"total"`
		}

		list := []List{
			{"/static/upload/1.jpg"},
			{"/static/upload/2.jpg"},
		}

		listimage := Listimage{"SUCCESS", list, 1, 21}

		c.Data["json"] = listimage
		c.ServeJSON()

	case "catchimage":

		type List struct {
			Url    string `json:"url"`
			Source string `json:"source"`
			State  string `json:"state"`
		}

		type Catchimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
		}

		list := []List{
			{"/static/upload/1.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			{"/static/upload/2.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
		}
		catchimage := Catchimage{"SUCCESS", list}

		c.Data["json"] = catchimage
		c.ServeJSON()

		file, header, err := c.GetFile("source") // r.FormFile("upfile")
		beego.Info(header.Filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			panic(err)
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		io.Copy(outFile, file)

	}
	// c.Write(configJson)
}

// func UploadImage(w http.ResponseWriter, r *http.Request) { //这个没用
// 	file, header, err := r.FormFile("upfile")
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
// 	b, err := json.Marshal(map[string]string{
// 		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
// 		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
// 		"original": header.Filename,                            //原始文件名
// 		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(b))
// 	w.Write(b)
// }

// func (c *UeditorController) UploadImage() { //对应这个路由 beego.Router("/controller", &controllers.UeditorController{}, "post:UploadImage")

// 	file, header, err := c.GetFile("upfile") // r.FormFile("upfile")
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
// 	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + filename, "title": "111", "original": "demo.jpg"}
// 	c.ServeJSON()
// }
