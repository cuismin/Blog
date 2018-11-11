package controllers

import (
	"strings"
)

type PageToolController struct {
	BaseController
}


func (this *PageToolController) URLMapping() {
	this.Mapping("Page", this.Page)
}
// @router /page/:page [get]
func (controller *PageToolController) Page() {
	 page := controller.Ctx.Input.Param(":page")//拿到当前页面
	 page = strings.Replace(page, "*", "/", -1);
	 controller.TplName = page;
}