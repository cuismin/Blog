package controllers

import (
	"github.com/astaxie/beego"
)


type BaseController struct {
	beego.Controller
}

//获取app.conf 配置文件数据
func (this *BaseController) GetConfigParameter(key string) string{
	return beego.AppConfig.String(key)
}

//获取请求参数
func (this *BaseController) GetParameters() map[string]string{
    parameter := make(map[string]string);
	this.Ctx.Request.ParseForm()
	if len(this.Ctx.Request.Form) > 0 {
		for key, _ := range this.Ctx.Request.Form {
			parameter[key] = this.Ctx.Request.Form.Get(key)
		}
	}
	return parameter
}



