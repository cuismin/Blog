package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "DeleteArticle",
			Router: `/article/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetArticleInfo",
			Router: `/article/getArticle/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "ArticleInfo",
			Router: `/article/info/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "QueryArticle",
			Router: `/article/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "SaveArticle",
			Router: `/article/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "UpdateArticle",
			Router: `/article/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"],
		beego.ControllerComments{
			Method: "DeleteArticleType",
			Router: `/ArticleType/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"],
		beego.ControllerComments{
			Method: "GetArticleTypeInfo",
			Router: `/ArticleType/getArticleType/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"],
		beego.ControllerComments{
			Method: "QueryArticleType",
			Router: `/ArticleType/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"],
		beego.ControllerComments{
			Method: "SaveArticleType",
			Router: `/ArticleType/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["Blog/controllers:ArticleTypeController"],
		beego.ControllerComments{
			Method: "UpdateArticleType",
			Router: `/ArticleType/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:FileUploadController"] = append(beego.GlobalControllerRouter["Blog/controllers:FileUploadController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/file/upload?:guid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "DeleteNote",
			Router: `/Note/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "GetNoteInfo",
			Router: `/Note/getNote/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "NoteInfo",
			Router: `/Note/info/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "QueryNote",
			Router: `/Note/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "SaveNote",
			Router: `/Note/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["Blog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "UpdateNote",
			Router: `/Note/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:PageToolController"] = append(beego.GlobalControllerRouter["Blog/controllers:PageToolController"],
		beego.ControllerComments{
			Method: "Page",
			Router: `/page/:page`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:UserController"] = append(beego.GlobalControllerRouter["Blog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Admin",
			Router: `/cuism`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:UserController"] = append(beego.GlobalControllerRouter["Blog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Blog/controllers:UserController"] = append(beego.GlobalControllerRouter["Blog/controllers:UserController"],
		beego.ControllerComments{
			Method: "QueryUser",
			Router: `/user/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
