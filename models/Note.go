package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/satori/go.uuid"
	"strings"
	"strconv"
	"github.com/astaxie/beego"
	"math"
)

func init() {
	//注册表
	orm.RegisterModel(new(Note))
}
type NoteInterface interface {
	GetTableName() string
	ListNote(condition map[string]string) map[string]interface{}
	Delete(data map[string]string) Result
	GetNoteInfo(uuid string) Note
	Update(data map[string]string) Result
	Save(data map[string]string)Result
}
//文章
type Note struct {
	Uuid         string       `orm:"size(36);pk"` //序号
	Title        string       `orm:"size(200)"`   //标题
	Content      string       `orm:"type(text)"`  //内容
	CreateDate   string       `orm:"size(50)"`    //创建时间
	ArticleType  *ArticleType `orm:"rel(fk)"`   //文章类型
	CreateUser   string       `orm:"size(100)"` //创建人
	User         *User        `orm:"rel(fk)"`  //设置一对多关系
}
func (this *Note)GetTableName() string  {

	return "Note"
}

func(this *Note) ListNote(condition map[string]string) map[string]interface{} {
	result := make(map[string]interface{});
	o := orm.NewOrm()
	qs := o.QueryTable(this.GetTableName())
	cond := orm.NewCondition()
	page ,_ := strconv.Atoi(condition["page"])
	if page < 1 {
		page = 1
	}
	if len(condition["Title"]) > 0 {
		cond = cond.And("Title__icontains",condition["Title"])
	}
	if len(condition["ArticleType"]) > 0{
		cond = cond.And("article_type_id",condition["ArticleType"])
	}
	if len(condition["User"]) > 0 {
		cond = cond.And("create_user",condition["CreateUser"])
	}
	offset ,_ := beego.AppConfig.Int("pageSize")
	start := (page - 1) * offset
	var note []Note
	count ,_ := qs.SetCond(cond).Count()
	num, _ := qs.SetCond(cond).RelatedSel().OrderBy("-CreateDate").Limit(offset, start).All(&note)
	result["limt"] = page
	result["pageSize"] = num
	result["count"] = count
	result["countPage"] = int(math.Ceil(float64(count) / float64(offset)))
	result["data"] = SetResult("成功",note,true,200,"")
	return result
}


func (this *Note) Delete(data map[string]string) Result  {
	uuids := data["uuids"]
	if len(uuids) > 0 {
		field := strings.Split(uuids,",")
		for _ , data := range field {
			note := Note{Uuid:data}
			Delete(&note)
		}
		return SetResult("删除成功","",true,200,"")
	}
	return SetResult("删除失败","",false,405,"")
}


func (this *Note) GetNoteInfo(uuid string) Note  {
	note := Note{Uuid:uuid}
	QueryObject(&note);
	return note;
}


func (this *Note) Update(data map[string]string) Result  {
	note := this.getNoteData(data)
	if UpdateData(&note){
		return SetResult("修改成功","",true,200,"")
	}
	return SetResult("修改失败","",false,400,"")
}

func (this *Note) Save(data map[string]string)Result {
	article := this.getNoteData(data)
	if SaveData(&article){
		return SetResult("添加成功","",true,200,"")
	}
	return SetResult("添加失败","",false,400,"")
}

func (this *Note) getNoteData(data map[string]string) Note{
	udata := data["data"]
	noteData := Note{}
	var resultNote Note
	if len(udata) > 0 {
		json.Unmarshal([]byte(udata) , &noteData);
		if len(noteData.Uuid) > 0 {
			resultNote.Uuid = noteData.Uuid
		}else{
			uuid, _ := uuid.NewV4()
			resultNote.Uuid  = uuid.String()
		}
		resultNote.Title = noteData.Title
		resultNote.Content = noteData.Content
		resultNote.CreateDate = GetCurrentDate()
		resultNote.ArticleType = &ArticleType{Uuid:noteData.ArticleType.Uuid}
		resultNote.CreateUser = noteData.CreateUser
		resultNote.User = &User{Uuid:noteData.User.Uuid}
		return resultNote
	}
	return resultNote
}