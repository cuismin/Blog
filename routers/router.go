package routers

import (
	"Blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//设置注解
	beego.Include(&controllers.PageToolController{})
	beego.Include(&controllers.ArticleController{})
	beego.Include(&controllers.ArticleTypeController{})
	beego.Include(&controllers.FileUploadController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.NoteController{})
}
