package controllers

import (
	"Blog/models"
)

type ArticleController struct {
	BaseController
}

var articleInterface  = new (models.Article)

//设置方法
func (this *ArticleController) URLMapping() {
	this.Mapping("QueryArticle", this.QueryArticle)
	this.Mapping("ArticleInfo",this.ArticleInfo)
	this.Mapping("DeleteArticle",this.DeleteArticle)
	this.Mapping("GetArticleInfo",this.GetArticleInfo)
	this.Mapping("UpdateArticle",this.UpdateArticle)
}
// @router /article/query [get]
func (this *ArticleController) QueryArticle() {
	result := articleInterface.ListArticle(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}
//根据ID查询文档信息
// @router /article/info/:uuid [get]
func (this *ArticleController) ArticleInfo() {
	uuid := this.Ctx.Input.Param(":uuid")//拿到当前页面
	this.Data["content"] = articleInterface.GetArticleInfo(uuid).Content;
	this.TplName = "home/blog.html";
}

// @router /article/getArticle/:uuid [get]
func (this *ArticleController) GetArticleInfo() {
	uuid := this.Ctx.Input.Param(":uuid")//拿到当前页面
	this.SetData(articleInterface.GetArticleInfo(uuid))
	this.ServeJSON()
}

//根据ID查询文档信息
// @router /article/del [post]
func (this *ArticleController) DeleteArticle() {
	result := articleInterface.Delete(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}


//修改数据
// @router /article/update [post]
func (this *ArticleController) UpdateArticle() {
	result := articleInterface.Update(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

//添加数据
// @router /article/save [post]
func (this *ArticleController) SaveArticle() {
	result := articleInterface.Save(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}