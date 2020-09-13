// @APIVersion 1.0.0
// @Title MeritMS API
// @Description MeritMS has every tool to get any job done, so codename for the new MeritMS APIs.
// @Contact 504284@qq.com
package routers

import (
	"github.com/3xxx/meritms/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//自动化文档
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/admin",
				beego.NSInclude(
					&controllers.AdminController{},
					// &controllers.CustomerCookieCheckerController{},
				),
			),
			beego.NSNamespace("/wx",
				beego.NSInclude(
					&controllers.FroalaController{},
					&controllers.LoginController{},
					&controllers.RegistController{},
					&controllers.MeritController{},
				),
			),
			beego.NSNamespace("/adminlog",
				beego.NSInclude(
					&controllers.AdminLogController{},
				),
			),

			// beego.NSNamespace("/cms",
			// 	beego.NSInclude(
			// 		&controllers.CMSController{},
			// 	),
			// ),
			// beego.NSNamespace("/suggest",
			// 	beego.NSInclude(
			// 		&controllers.SearchController{},
			// 	),
			// ),
		)
	beego.AddNamespace(ns)
	// beego.Router("/.well-known/pki-validation/AC9A20F9BD09F18D247337AABC67BC06.txt", &controllers.AdminController{}, "*:Testdown")
	beego.Router("/.well-known/pki-validation/*", &controllers.AdminController{}, "*:Testdown")
	beego.Router("/role/test", &controllers.RoleController{}, "*:Test")

	//api接口
	//显示首页
	beego.Router("/index", &controllers.IndexController{}, "*:GetIndex")
	//首页放到onlyoffice
	// beego.Router("/", &controllers.OnlyController{}, "get:Get")
	//显示右侧页面框架
	beego.Router("/index/user", &controllers.IndexController{}, "*:GetUser")

	//后台
	beego.Router("/admin", &controllers.AdminController{})
	// beego.Router("/admincategory", &controllers.AdminController{}, "*:GetAdminCategory")
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//批量添加首页轮播图片
	beego.Router("/admin/base/addcarousel", &controllers.AdminController{}, "*:AddCarousel")
	//获取首页轮播图片填充表格
	beego.Router("/admin/base/carousel", &controllers.AdminController{}, "*:Carousel")

	//根据数字id查询类别或目录分级表
	beego.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表_这里应该放多一个/category路径下
	beego.Router("/admin/categorytitle", &controllers.AdminController{}, "*:CategoryTitle")
	//添加目录类别
	beego.Router("/admin/category/addcategory", &controllers.AdminController{}, "*:AddCategory")
	//修改目录类别
	beego.Router("/admin/category/updatecategory", &controllers.AdminController{}, "*:UpdateCategory")
	//删除目录类
	beego.Router("/admin/category/deletecategory", &controllers.AdminController{}, "*:DeleteCategory")

	//添加IP地址段
	beego.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	beego.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//删除IP地址段
	beego.Router("/admin/ipsegment/deleteipsegment", &controllers.AdminController{}, "*:DeleteIpsegment")

	//查询所有ip地址段
	beego.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")

	//根据项目id查询项目同步ip
	beego.Router("/admin/project/synchip/:id:string", &controllers.AdminController{}, "*:SynchIp")
	//添加项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/addsynchip", &controllers.AdminController{}, "*:AddsynchIp")
	//修改项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/updatesynchip", &controllers.AdminController{}, "*:UpdatesynchIp")
	//删除项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/deletesynchip", &controllers.AdminController{}, "*:DeletesynchIp")
	//后台部门结构
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
	beego.Router("/admin/merit", &controllers.AdminMeritController{}, "*:Merit")
	//根据数字id查询这个分类下的价值
	beego.Router("/admin/merit/?:id:string", &controllers.AdminMeritController{}, "*:Merit")
	//根据数字id查询这个科室id下的价值分类
	beego.Router("/admin/merit/secoffice/?:id:string", &controllers.AdminMeritController{}, "*:SecofficeMerit")
	//向科室里添加价值分类
	beego.Router("/admin/merit/secoffice/addsecofficemerit", &controllers.AdminMeritController{}, "*:AddSecofficeMerit")

	//添加
	beego.Router("/admin/merit/addmerit", &controllers.AdminMeritController{}, "*:AddMerit")
	//修改
	beego.Router("/admin/merit/updatemerit", &controllers.AdminMeritController{}, "*:UpdateMerit")
	//删除
	beego.Router("/admin/merit/deletemerit", &controllers.AdminMeritController{}, "*:DeleteMerit")

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

	//新建角色
	beego.Router("/admin/role/post", &controllers.RoleController{}, "post:Post")
	beego.Router("/admin/role/update", &controllers.RoleController{}, "put:Update")
	beego.Router("/admin/role/delete", &controllers.RoleController{}, "post:Delete")
	beego.Router("/admin/role/get/?:id:string", &controllers.RoleController{}, "get:Get")
	//添加用户角色
	beego.Router("/admin/role/userrole", &controllers.RoleController{}, "post:UserRole")
	//添加角色对项目目录文件操作权限
	//查询角色对项目目录文件操作的权限
	beego.Router("/admin/role/getpermission", &controllers.RoleController{}, "get:GetRolePermission")

	//角色——成果统计、价值统计查看权限
	beego.Router("/admin/role/roleachieve", &controllers.RoleController{}, "post:RoleAchieve")
	//查询角色具有的科室成果、科室价值查看权限
	beego.Router("/admin/role/getroleachieve", &controllers.RoleController{}, "get:GetRoleAchieve")

	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")
	//用户产看自己的table中数据填充
	beego.Router("/usermyself", &controllers.UserController{}, "get:Usermyself")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	//页面登录提交用户名和密码
	beego.Router("/post", &controllers.LoginController{}, "post:Post")
	//弹框登录提交用户名和密码
	beego.Router("/loginpost", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
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

	beego.Router("/regist", &controllers.RegistController{})
	// beego.Router("/registerr", &controllers.RegistController{}, "get:RegistErr")
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "*:GetUname")
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

	//根据附件绝对地址下载
	beego.Router("/attachment/*", &controllers.AttachController{}, "get:Attachment")
	beego.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")
	//这个有哦何用？
	beego.SetStaticPath("/attachment/wiki", "attachment/wiki")
	beego.SetStaticPath("/swagger", "swagger")

	//获得ecms提交过来的成果清单
	beego.Router("/getecmspost", &controllers.EcmsController{}, "post:GetEcmsPost")

}
