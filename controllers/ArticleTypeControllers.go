package controllers

import (
	"Blog/models"
)

var articleTypeInterface = new (models.ArticleType)

type ArticleTypeController struct {
	BaseController
}

//设置方法
func (this *ArticleTypeController) URLMapping() {
	this.Mapping("QueryArticleType", this.QueryArticleType)
	this.Mapping("GetArticleTypeInfo", this.GetArticleTypeInfo)
	this.Mapping("DeleteArticleType", this.DeleteArticleType)
	this.Mapping("UpdateArticleType", this.UpdateArticleType)
	this.Mapping("SaveArticleType", this.SaveArticleType)


}

// @router /ArticleType/query [get]
func (this *ArticleTypeController) QueryArticleType() {
	result := articleTypeInterface.ListArticleType(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

// @router /ArticleType/getArticleType/:uuid [get]
func (this *ArticleTypeController) GetArticleTypeInfo() {
	uuid := this.Ctx.Input.Param(":uuid")//拿到当前页面
	this.SetData(articleTypeInterface.GetArticleTypeInfo(uuid))
	this.ServeJSON()
}

//根据ID查询文档信息
// @router /ArticleType/del [post]
func (this *ArticleTypeController) DeleteArticleType() {
	result := articleTypeInterface.Delete(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

//修改数据
// @router /ArticleType/update [post]
func (this *ArticleTypeController) UpdateArticleType() {
	result := articleTypeInterface.Update(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

//添加数据
// @router /ArticleType/save [post]
func (this *ArticleTypeController) SaveArticleType() {
	result := articleTypeInterface.Save(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}


