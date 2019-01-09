package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Category",
			Router: `/category/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/category/addcategory`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "CategoryTitle",
			Router: `/categorytitle`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "ErrLog",
			Router: `/errlog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "InfoLog",
			Router: `/infolog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticle",
			Router: `/addwxarticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticles",
			Router: `/addwxarticles/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticle",
			Router: `/getwxarticle/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticles",
			Router: `/getwxarticles`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticless",
			Router: `/getwxarticless/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGetDocTypeByName",
			Router: `/flowgetdoctypebyname`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGetDocTypeByName1",
			Router: `/flowgetdoctypebyname1`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "WorkFlow",
			Router: `/workflow`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "WorkFlowData",
			Router: `/workflowdata`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "WorkFlowData1",
			Router: `/workflowdata1`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImg",
			Router: `/uploadwximg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImgs",
			Router: `/uploadwximgs/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:LoginController"],
		beego.ControllerComments{
			Method: "WxLogin",
			Router: `/wxlogin/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxPdf",
			Router: `/wxpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxStandardPdf",
			Router: `/wxstandardpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:RegistController"],
		beego.ControllerComments{
			Method: "WxRegist",
			Router: `/wxregist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxLike",
			Router: `/addwxlike/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxRelease",
			Router: `/addwxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "DeleteWxRelease",
			Router: `/deletewxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:SearchController"],
		beego.ControllerComments{
			Method: "SearchWxDrawings",
			Router: `/searchwxdrawings`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/meritms/controllers:StandardController"],
		beego.ControllerComments{
			Method: "SearchWxStandards",
			Router: `/searchwxstandards`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
