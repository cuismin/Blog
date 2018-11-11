package controllers

import (
	"Blog/models"
)

type UserController struct {
	BaseController
}


var userInterface = new (models.User)

//设置方法
func (this *UserController) URLMapping() {
	this.Mapping("QueryUser", this.QueryUser)
	this.Mapping("Login", this.Login)
	this.Mapping("Admin", this.Admin)
}

// @router /user/query [get]
func (this *UserController) QueryUser() {
	result := userInterface.ListUser(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}
// @router /user/login [post]
func (this *UserController)Login(){
	parameters := this.GetParameters()
	userName := parameters["UserName"]
	Password := parameters["Password"]
	user := userInterface.GetUserInfo(userName);
	var result models.Result
	if user.Password != models.Md5(Password) {
		result = models.SetResult("用户名或密码错误",nil,false,401,"")
	}else{
		sid := models.Md5(user.Uuid+user.Password)
		result = models.SetResult("登录成功",user,true,200,sid)
	}
	this.SetData(result)
	this.ServeJSON()
}

// @router /cuism [get]
func (this *UserController) Admin() {
	this.TplName = "login.html";
}
