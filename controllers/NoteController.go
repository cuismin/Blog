package controllers

import "Blog/models"

var noteInterface  = new (models.Note)
type NoteController struct {
	BaseController
}

//设置方法
func (this *NoteController) URLMapping() {
	this.Mapping("QueryNote", this.QueryNote)
	this.Mapping("NoteInfo",this.NoteInfo)
	this.Mapping("DeleteNote",this.DeleteNote)
	this.Mapping("GetNoteInfo",this.GetNoteInfo)
	this.Mapping("UpdateNote",this.UpdateNote)
}
// @router /Note/query [get]
func (this *NoteController) QueryNote() {
	result := noteInterface.ListNote(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}
//根据ID查询文档信息
// @router /Note/info/:uuid [get]
func (this *NoteController) NoteInfo() {
	uuid := this.Ctx.Input.Param(":uuid")//拿到当前页面
	this.Data["content"] = noteInterface.GetNoteInfo(uuid)
	this.TplName = "home/note.html";
}

// @router /Note/getNote/:uuid [get]
func (this *NoteController) GetNoteInfo() {
	uuid := this.Ctx.Input.Param(":uuid")//拿到当前页面
	this.SetData(noteInterface.GetNoteInfo(uuid))
	this.ServeJSON()
}

//根据ID查询文档信息
// @router /Note/del [post]
func (this *NoteController) DeleteNote() {
	result := noteInterface.Delete(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

//修改数据
// @router /Note/update [post]
func (this *NoteController) UpdateNote() {
	result := noteInterface.Update(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

//添加数据
// @router /Note/save [post]
func (this *NoteController) SaveNote() {
	result := noteInterface.Save(this.GetParameters());
	this.SetData(result)
	this.ServeJSON()
}

