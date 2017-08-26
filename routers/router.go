// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"meritms/controllers"
)

func init() {
	//1.首页index
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/1/slide", &controllers.MainController{}, "*:Slide")
	beego.Router("/zsj", &controllers.MainController{}, "*:ZSJ")
	// beego.Router("/postdata", &controllers.MainController{}, "*:Postdata")
	//显示首页
	beego.Router("/", &controllers.IndexController{}, "*:GetIndex")
	beego.Router("/index", &controllers.IndexController{}, "*:GetIndex")
	//显示右侧页面框架
	beego.Router("/index/user", &controllers.IndexController{}, "*:GetUser")
	//这里显示用户查看主人日程
	beego.Router("/calendar", &controllers.IndexController{}, "*:Calendar")
	//会议室预定日程
	beego.Router("/meetingroom", &controllers.IndexController{}, "*:MeetingroomCalendar")
	//车辆预定日程
	beego.Router("/car", &controllers.IndexController{}, "*:GetCarCalendar")
	//首页搜索项目或成果
	beego.Router("/index/searchproject", &controllers.SearchController{}, "*:SearchProject")
	beego.Router("/index/searchproduct", &controllers.SearchController{}, "*:SearchProduct")

	//管理员
	beego.Router("/admin", &controllers.AdminController{}) //admin-get()
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//批量添加首页轮播图片
	beego.Router("/admin/base/addcarousel", &controllers.AdminController{}, "*:AddCarousel")
	//获取首页轮播图片填充表格
	beego.Router("/admin/base/carousel", &controllers.AdminController{}, "*:Carousel")

	//根据数字id查询类别或目录分级表
	beego.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表
	beego.Router("/admin/categorytitle", &controllers.AdminController{}, "*:CategoryTitle")
	//添加目录类别
	beego.Router("/admin/category/addcategory", &controllers.AdminController{}, "*:AddCategory")
	//修改目录类别
	beego.Router("/admin/category/updatecategory", &controllers.AdminController{}, "*:UpdateCategory")
	//删除目录类
	beego.Router("/admin/category/deletecategory", &controllers.AdminController{}, "*:DeleteCategory")

	//显示项目目录:注意这个方法放在projcontroller中
	beego.Router("/admin/project/getprojectcate/:id:string", &controllers.ProjController{}, "*:GetProjectCate")
	//添加子节点
	beego.Router("/admin/project/addprojectcate", &controllers.ProjController{}, "*:AddProjectCate")
	//修改节点名
	beego.Router("/admin/project/updateprojectcate", &controllers.ProjController{}, "*:UpdateProjectCate")
	//删除节点和其下子节点——和删除项目一样
	beego.Router("/admin/project/deleteprojectcate", &controllers.ProjController{}, "*:DeleteProject")

	//根据项目id查询项目同步ip
	beego.Router("/admin/project/synchip/:id:string", &controllers.AdminController{}, "*:SynchIp")
	//添加项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/addsynchip", &controllers.AdminController{}, "*:AddsynchIp")
	//修改项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/updatesynchip", &controllers.AdminController{}, "*:UpdatesynchIp")
	//删除项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/deletesynchip", &controllers.AdminController{}, "*:DeletesynchIp")

	//填充部门表格数据
	beego.Router("/admin/department", &controllers.AdminController{}, "*:Department")
	//根据数字id查询类别或目录分级表
	beego.Router("/admin/department/?:id:string", &controllers.AdminController{}, "*:Department")
	//根据名字查询目录分级表
	beego.Router("/admin/departmenttitle", &controllers.AdminController{}, "*:DepartmentTitle")
	//添加目录类别
	beego.Router("/admin/department/adddepartment", &controllers.AdminController{}, "*:AddDepartment")
	//修改目录类别
	beego.Router("/admin/department/updatedepartment", &controllers.AdminController{}, "*:UpdateDepartment")
	//删除目录类
	beego.Router("/admin/department/deletedepartment", &controllers.AdminController{}, "*:DeleteDepartment")
	//*******价值****
	//填充MERIT表格数据
	beego.Router("/admin/merit", &controllers.AdminController{}, "*:Merit")
	//根据数字id查询这个分类下的价值
	beego.Router("/admin/merit/?:id:string", &controllers.AdminController{}, "*:Merit")
	//根据数字id查询这个科室id下的价值分类
	beego.Router("/admin/merit/secoffice/?:id:string", &controllers.AdminController{}, "*:SecofficeMerit")
	//向科室里添加价值分类
	beego.Router("/admin/merit/secoffice/addsecofficemerit", &controllers.AdminController{}, "*:AddSecofficeMerit")

	//添加
	beego.Router("/admin/merit/addmerit", &controllers.AdminController{}, "*:AddMerit")
	//修改
	beego.Router("/admin/merit/updatemerit", &controllers.AdminController{}, "*:UpdateMerit")
	//删除
	beego.Router("/admin/merit/deletemerit", &controllers.AdminController{}, "*:DeleteMerit")

	//查询所有ip地址段
	beego.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")
	//添加IP地址段
	beego.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	beego.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//删除IP地址段
	beego.Router("/admin/ipsegment/deleteipsegment", &controllers.AdminController{}, "*:DeleteIpsegment")
	//编辑成果类型和折标系数表
	beego.Router("/admin/achievcategory", &controllers.Achievement{}, "get:Achievcategory")
	beego.Router("/admin/achievcategory/addachievcategory", &controllers.Achievement{}, "post:AddAchievcategory")
	beego.Router("/admin/achievcategory/updateachievcategory", &controllers.Achievement{}, "post:UpdateAchievcategory")
	beego.Router("/admin/achievcategory/deleteachievcategory", &controllers.Achievement{}, "post:DeleteAchievcategory")

	// beego.Router("/jsoneditor", &controllers.AdminController{}, "get:Jsoneditor")
	//如果后面不带id，则显示所有用户
	beego.Router("/admin/user/?:id:string", &controllers.UserController{}, "*:User")
	//添加用户
	beego.Router("/admin/user/adduser", &controllers.UserController{}, "*:AddUser")
	//导入用户
	beego.Router("/admin/user/importusers", &controllers.UserController{}, "*:ImportUsers")

	//修改用户
	beego.Router("/admin/user/updateuser", &controllers.UserController{}, "*:UpdateUser")
	//删除用户
	beego.Router("/admin/user/deleteuser", &controllers.UserController{}, "*:DeleteUser")
	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")
	//用户产看自己的table中数据填充
	beego.Router("/usermyself", &controllers.UserController{}, "get:Usermyself")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	beego.Router("/roleerr", &controllers.UserController{}, "*:Roleerr") //显示权限不够

	//2.2首页进入成果登记
	beego.Router("/achievement", &controllers.Achievement{}, "get:GetAchievement")
	//这个同上面一样
	beego.Router("/getachievement", &controllers.Achievement{}, "get:GetAchievement")
	//个人在线添加成果
	beego.Router("/achievement/addcatalog", &controllers.Achievement{}, "post:AddCatalog")
	//个人在线直接提交成果
	beego.Router("/achievement/sendcatalog", &controllers.Achievement{}, "post:SendCatalog")
	//在线退回成果
	beego.Router("/achievement/downsendcatalog", &controllers.Achievement{}, "post:DownSendCatalog")
	//个人在线修改保存成果
	beego.Router("/achievement/modifycatalog", &controllers.Achievement{}, "post:ModifyCatalog")
	//如果一个成果的附件多于1条，则查看一个成果的附件列表
	beego.Router("/achievement/catalog/attachment/:id:string", &controllers.Achievement{}, "get:CatalogAttachment")
	//修改附件
	beego.Router("/achievement/catalog/modifylink", &controllers.Achievement{}, "post:ModifyLink")
	//查看一个成果的校审意见列表
	beego.Router("/achievement/catalog/content/:id:string", &controllers.Achievement{}, "get:CatalogContent")
	//修改校审意见
	beego.Router("/achievement/catalog/modifycontent", &controllers.Achievement{}, "post:ModifyContent")

	//个人在线删除一条成功
	beego.Router("/achievement/delete", &controllers.Achievement{}, "post:DeleteCatalog")
	//测试某个专业下总成本分布情况
	beego.Router("/achievement/specialty", &controllers.Achievement{}, "get:Specialty")
	//个人当月成果类型组成
	beego.Router("/achievement/echarts", &controllers.Achievement{}, "get:Echarts")
	//beego.Router("/achievement/echarts1", &controllers.Achievement{}, "get:Echarts1")
	//个人当年成果类型组成
	beego.Router("/achievement/echarts2", &controllers.Achievement{}, "get:Echarts2")
	//项目阶段专业，全年成果类型组成
	beego.Router("/achievement/echarts3", &controllers.Achievement{}, "get:Echarts3")

	//点击个人参与的项目，弹出模态框，显示这个项目所有成果
	beego.Router("/achievement/projectachievement", &controllers.Achievement{}, "*:ProjectAchievement")
	//一年来一个项目的每个人的贡献率
	beego.Router("/achievement/projectuserparticipate", &controllers.Achievement{}, "*:ProjectUserParticipate")

	beego.Router("/test", &controllers.TestController{}, "get:Test")
	// beego.Router("/test1", &controllers.AdminController{}, "get:Test1")

	//这个是ue编辑器用的
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")

	//2.1首页进入价值——根据登陆者的权限，显示对应的侧栏：普通用户直接进入自己的价值侧栏
	beego.Router("/merit", &controllers.MeritController{}, "get:GetMerit")
	// 主页里显示iframe——科室总体情况
	beego.Router("/merit/secofficeshow", &controllers.MeritController{}, "get:Secofficeshow")
	// 填充科室总体情况数据
	beego.Router("/merit/secofficedata", &controllers.MeritController{}, "get:SecofficeData")

	// 2.1.1用户进入价值，侧栏右边的iframe
	beego.Router("/merit/myself", &controllers.MeritController{}, "get:Myself")
	// 2.1.1用户点击左侧价值，侧栏右边的iframe显示这个价值的内容列表
	//根据id=1,2,3分别显示准备提交，已经提交，已经完成
	beego.Router("/merit/send/:id:int", &controllers.MeritController{}, "get:MeritSend")
	//管理人员登录看到需要处理审核的价值内容
	beego.Router("/merit/examined", &controllers.MeritController{}, "get:MeritExamined")

	// 添加merit
	beego.Router("/merit/addmerit", &controllers.MeritController{}, "post:AddMerit")
	// 删除merittopic
	beego.Router("/merit/delete", &controllers.MeritController{}, "post:Delete")

	// 传递merit
	beego.Router("/merit/sendmerit", &controllers.MeritController{}, "post:SendMerit")
	beego.Router("/merit/downsendmerit", &controllers.MeritController{}, "post:DownSendMerit")

	// 修改merit
	beego.Router("/merit/updatemerit", &controllers.MeritController{}, "post:UpdateMerit")

	//显示所有人价值排名
	// beego.Router("/getperson", &controllers.JsonController{}, "get:GetPerson")
	//进行价值结构编辑
	// beego.Router("/get", &controllers.JsonController{})
	// beego.Router("/json", &controllers.JsonController{}) //这个和上面等价
	// beego.Router("/importjson", &controllers.JsonController{}, "post:ImportJson")

	// beego.Router("/addjson", &controllers.JsonController{}, "post:Addjson")
	// beego.Router("/modifyjson", &controllers.JsonController{}, "get:Modifyjson")          //显示修改页面
	// beego.Router("/modifyjsonpost", &controllers.JsonController{}, "post:ModifyjsonPost") //提交修改
	// beego.Router("/deletejson", &controllers.JsonController{}, "get:Deletejson")

	// beego.Router("/merit/add", &controllers.MeritTopicController{}, "get:MeritAdd")
	// beego.Router("/AddMeritTopic", &controllers.MeritTopicController{}, "post:AddMeritTopic")
	// beego.Router("/view", &controllers.MeritTopicController{}, "get:ViewMeritTopic")
	// beego.Router("/modify", &controllers.MeritTopicController{}, "get:ModifyMeritTopic")
	// beego.Router("/ModifyPost", &controllers.MeritTopicController{}, "post:ModifyPost")
	// beego.Router("/delete", &controllers.MeritTopicController{}, "get:DeleteMeritTopic")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")

	beego.Router("/regist", &controllers.RegistController{})
	// beego.Router("/registerr", &controllers.RegistController{}, "get:RegistErr")
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "post:GetUname")
	//get方法用于x-editable的select2方法_作废，select2不必须要动态数据
	beego.Router("/regist/getuname1", &controllers.RegistController{}, "get:GetUname1")

	//成果登记系统
	//成果登记表导入数据库
	beego.Router("/achievement/import_xls_catalog", &controllers.Achievement{}, "post:Import_Xls_Catalog")
	// 主页里显示iframe——科室总体情况
	beego.Router("/achievement/secofficeshow", &controllers.Achievement{}, "get:Secofficeshow")
	// 填充科室总体情况数据
	beego.Router("/achievement/secofficedata", &controllers.Achievement{}, "get:SecofficeData")

	//根据id=1,2,3,4分别显示准备提交，设计，校核，审查;5已经提交,6已经完成
	beego.Router("/achievement/send/:id:int", &controllers.Achievement{}, "get:AchievementSend")
	//用户在线登记时，自己发起的成果，还未提交
	// beego.Router("/achievement/myself", &controllers.Achievement{}, "get:Myself")
	//自己发起的成果，已经提交
	// beego.Router("/achievement/running", &controllers.Achievement{}, "get:Running")
	//别人传来，自己处于设计位置
	// beego.Router("/achievement/designd", &controllers.Achievement{}, "get:Designd")
	//别人传来，自己处于校核位置
	// beego.Router("/achievement/checked", &controllers.Achievement{}, "get:Checked")
	//别人传来，自己处于审查位置
	// beego.Router("/achievement/examined", &controllers.Achievement{}, "get:Examined")
	//查看用户个人时，获取已经完成的数据
	// beego.Router("/achievement/completed", &controllers.Achievement{}, "get:Completed")
	//获取自己参与的项目列表
	beego.Router("/achievement/participate", &controllers.Achievement{}, "get:Participate")
	//获取科室的项目列表
	beego.Router("/achievement/secparticipate", &controllers.Achievement{}, "get:SecParticipate")
	//获取科室的当月成果列表
	beego.Router("/achievement/secprojectachievement", &controllers.Achievement{}, "get:SecProjectAchievement")

	//人员管理
	// beego.Router("/user/AddUser", &controllers.UserController{}, "*:AddUser")
	// beego.Router("/user/UpdateUser", &controllers.UserController{}, "*:UpdateUser")
	// beego.Router("/user/deluser", &controllers.UserController{}, "*:DelUser")
	// beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	// //管理员修改用户资料
	// beego.Router("/user/view", &controllers.UserController{}, "get:View")
	// beego.Router("/user/view/*", &controllers.UserController{}, "get:View")
	// beego.Router("/user/importexcel", &controllers.UserController{}, "post:ImportExcel")

	// //用户修改自己密码
	// beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	// //用户登录后查看自己的资料
	// beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")

	// beego.Router("/role/AddAndEdit", &controllers.RoleController{}, "*:AddAndEdit")
	// beego.Router("/role/DelRole", &controllers.RoleController{}, "*:DelRole")
	// // beego.Router("/role/AccessToNode", &controllers.RoleController{}, "*:AccessToNode")
	// // beego.Router("/role/AddAccess", &controllers.RoleController{}, "*:AddAccess")
	// beego.Router("/role/RoleToUserList", &controllers.RoleController{}, "*:RoleToUserList")
	// beego.Router("/role/AddRoleToUser", &controllers.RoleController{}, "*:AddRoleToUser")
	// beego.Router("/role/Getlist", &controllers.RoleController{}, "*:Getlist")
	// beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	// beego.Router("/roleerr", &controllers.RoleController{}, "*:Roleerr") //显示权限不够
	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")
	//修改项目
	beego.Router("/project/updateproject", &controllers.ProjController{}, "*:UpdateProject")
	//删除项目
	beego.Router("/project/deleteproject", &controllers.ProjController{}, "*:DeleteProject")
	//项目时间轴
	beego.Router("/project/:id([0-9]+)/gettimeline", &controllers.ProjController{}, "get:ProjectTimeline")
	beego.Router("/project/:id([0-9]+)/timeline", &controllers.ProjController{}, "get:Timeline")

	//根据项目id进入一个具体项目的侧栏
	beego.Router("/project/:id([0-9]+)", &controllers.ProjController{}, "*:GetProject")
	//进入项目日历
	beego.Router("/project/:id([0-9]+)/getcalendar", &controllers.ProjController{}, "*:GetCalendar")
	//取得日历数据
	beego.Router("/project/:id([0-9]+)/calendar", &controllers.ProjController{}, "*:Calendar")
	//添加日历
	beego.Router("/project/:id([0-9]+)/calendar/addcalendar", &controllers.ProjController{}, "*:AddCalendar")
	//修改
	beego.Router("/project/calendar/updatecalendar", &controllers.ProjController{}, "*:UpdateCalendar")
	//删除
	beego.Router("/project/calendar/deletecalendar", &controllers.ProjController{}, "*:DeleteCalendar")
	//拖曳事件
	beego.Router("/project/calendar/dropcalendar", &controllers.ProjController{}, "*:DropCalendar")
	//resize事件
	beego.Router("/project/calendar/resizecalendar", &controllers.ProjController{}, "*:ResizeCalendar")
	//日历中传照片
	beego.Router("/project/calendar/uploadimage", &controllers.ProjController{}, "*:UploadImage")

	//点击侧栏，根据id返回json数据给导航条
	beego.Router("/project/navbar/:id:string", &controllers.ProjController{}, "*:GetProjNav")

	//根据项目侧栏id显示这个id下的成果界面——作废，用上面GetProject界面
	beego.Router("/project/:id:string/:id:string", &controllers.ProdController{}, "*:GetProjProd")

	//给上面那个页面提供table所用的json数据
	beego.Router("/project/products/:id:string", &controllers.ProdController{}, "*:GetProducts")
	//点击项目名称，显示这个项目下所有成果
	beego.Router("/project/allproducts/:id:string", &controllers.ProjController{}, "*:GetProjProducts")
	//这个对应上面的页面进行表格内数据填充，用的是prodcontrollers中的方法
	beego.Router("/project/products/all/:id:string", &controllers.ProdController{}, "*:GetProjProducts")
	//给上面那个页面提供项目同步ip的json数据
	beego.Router("/project/synchproducts/:id:string", &controllers.ProdController{}, "*:GetsynchProducts")

	//对外提供同步成果接口json数据
	beego.Router("/project/providesynchproducts", &controllers.ProdController{}, "*:ProvidesynchProducts")

	//向项目侧栏id下添加成果_这个没用到，用下面addattachment
	// beego.Router("/project/product/addproduct", &controllers.ProdController{}, "post:AddProduct")
	//编辑成果信息
	beego.Router("/project/product/updateproduct", &controllers.ProdController{}, "post:UpdateProduct")
	//删除成果
	beego.Router("/project/product/deleteproduct", &controllers.ProdController{}, "post:DeleteProduct")
	//取得成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/attachment/:id:string", &controllers.AttachController{}, "*:GetAttachments")
	//取得同步成果中所有附件中的非pdf附件列表
	beego.Router("project/product/synchattachment", &controllers.AttachController{}, "*:GetsynchAttachments")
	//提供接口：同步成果中所有附件中的非pdf附件列表
	beego.Router("project/product/providesynchattachment", &controllers.AttachController{}, "*:ProvideAttachments")

	//取得成果中所有附件中——包含pdf，非pdf，文章
	beego.Router("/project/product/allattachments/:id:string", &controllers.AttachController{}, "*:GetAllAttachments")

	//取得成果中所有附件的pdf列表
	beego.Router("/project/product/pdf/:id:string", &controllers.AttachController{}, "*:GetPdfs")
	//取得同步成果中所有pdf列表
	beego.Router("/project/product/synchpdf", &controllers.AttachController{}, "*:GetsynchPdfs")
	//提供同步成果中所有pdf列表
	beego.Router("/project/product/providesynchpdf", &controllers.AttachController{}, "*:ProvidePdfs")

	//根据文章id取得成果中的文章
	beego.Router("/project/product/article/:id:string", &controllers.ArticleController{}, "*:GetArticle")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/articles/:id:string", &controllers.ArticleController{}, "*:GetArticles")
	//取得同步成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/syncharticles", &controllers.ArticleController{}, "*:GetsynchArticles")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/providesyncharticles", &controllers.ArticleController{}, "*:ProvideArticles")
	//文章进行编辑
	beego.Router("/project/product/updatearticle", &controllers.ArticleController{}, "*:UpdateArticle")
	//删除文章
	beego.Router("/project/product/deletearticle", &controllers.ArticleController{}, "*:DeleteArticle")

	//查看一个成果
	// beego.Router("/project/product/?:id:string"

	//向成果里添加附件：批量一对一模式
	beego.Router("/project/product/addattachment", &controllers.AttachController{}, "post:AddAttachment")
	//向成果里添加附件：多附件模式
	beego.Router("/project/product/addattachment2", &controllers.AttachController{}, "post:AddAttachment2")
	//编辑附件列表：向成果里追加附件：多附件模式
	beego.Router("/project/product/updateattachment", &controllers.AttachController{}, "post:UpdateAttachment")
	//删除附件列表中的附件
	beego.Router("/project/product/deleteattachment", &controllers.AttachController{}, "post:DeleteAttachment")

	//向成果里添加文章
	beego.Router("/project/product/addarticle", &controllers.ArticleController{}, "post:AddArticle")

	//项目进度甘特图
	beego.Router("/projectgant", &controllers.ProjGantController{}, "get:Get")
	//数据填充
	// beego.Router("/projectgant/getprojgants", &controllers.ProjGantController{}, "get:GetProjGants")

	//添加项目进度
	beego.Router("/projectgant/addprojgant", &controllers.ProjGantController{}, "post:AddProjGant")
	//导入项目进度数据
	beego.Router("/projectgant/importprojgant", &controllers.ProjGantController{}, "post:ImportProjGant")

	//修改项目进度
	beego.Router("/projectgant/updateprojgant", &controllers.ProjGantController{}, "post:UpdateProjGant")
	//删除项目进度
	beego.Router("/projectgant/deleteprojgant", &controllers.ProjGantController{}, "post:DeleteProjGant")
	//关闭项目进度
	beego.Router("/projectgant/closeprojgant", &controllers.ProjGantController{}, "post:CloseProjGant")

	//附件下载"/attachment/*", &controllers.AttachController{}
	// beego.InsertFilter("/attachment/*", beego.BeforeRouter, controllers.ImageFilter)
	beego.Router("/attachment/*", &controllers.AttachController{}, "get:DownloadAttachment")
	//上面用attachment.ImageFilter是不行的，必须是package.func
	//首页轮播图片的权限
	beego.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")

	//添加日历
	beego.Router("/index/carcalendar/addcalendar", &controllers.IndexController{}, "*:AddCarCalendar")
	//查询
	beego.Router("/index/carcalendar", &controllers.IndexController{}, "*:CarCalendar")
	//修改
	beego.Router("/index/carcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateCarCalendar")
	//删除
	beego.Router("/index/carcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteCarCalendar")
	//拖曳事件
	beego.Router("/index/carcalendar/dropcalendar", &controllers.IndexController{}, "*:DropCarCalendar")
	//resize事件
	beego.Router("/index/carcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeCarCalendar")

	//添加日历
	beego.Router("/index/meetingroomcalendar/addcalendar", &controllers.IndexController{}, "*:AddMeetCalendar")
	//查询
	beego.Router("/index/meetingroomcalendar", &controllers.IndexController{}, "*:MeetCalendar")
	//修改
	beego.Router("/index/meetingroomcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateMeetCalendar")
	//删除
	beego.Router("/index/meetingroomcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteMeetCalendar")
	//拖曳事件
	beego.Router("/index/meetingroomcalendar/dropcalendar", &controllers.IndexController{}, "*:DropMeetCalendar")
	//resize事件
	beego.Router("/index/meetingroomcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeMeetCalendar")

	//获得ecms提交过来的成果清单
	beego.Router("/getecmspost", &controllers.EcmsController{}, "post:GetEcmsPost")

	// beego.SetStaticPath("/views/cms", "views/cms")
	// beego.SetStaticPath("/res", "views/cms/res")
}
