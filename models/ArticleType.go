package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/json"
	"github.com/satori/go.uuid"
	"strconv"
	"github.com/astaxie/beego"
	"math"
)

func init() {
	//注册表
	orm.RegisterModel(new(ArticleType))
}

//文章
type ArticleType struct {
	Uuid       string     `orm:"size(36);pk"`   //序号
	TypeName   string     `orm:"size(50)"`      //类型名称
	CreateDate string     `json:"createDate"`   //创建时间
	Article    []*Article `orm:"reverse(many)"` //设置一对多方向关系
	Note       []*Note `orm:"reverse(many)"` //设置一对多方向关系
}

type ArticleTypeInterface interface {
	//获取表名
	GetTableName() string
	//查询所有文章类型信息
	GetArticleTypeAll() Result
	GetArticleTypeInfo(uuid string) ArticleType
	Delete(data map[string]string) Result
	Update(data map[string]string) Result
	Save(data map[string]string)   Result
	ListArticleType(condition map[string]string) map[string]interface{}

}

func (this *ArticleType) GetTableName() string {
	return "ArticleType"
}

func (this *ArticleType) GetArticleTypeAll() Result {

   return  QueryAll(this.GetTableName(),&[]ArticleType{});
}

func (this *ArticleType) GetArticleTypeInfo(uuid string) ArticleType  {
	articleType := ArticleType{Uuid:uuid}
	QueryObject(&articleType);
	return articleType;
}
func (this *ArticleType) Delete(data map[string]string) Result  {
	uuids := data["uuids"]
	if len(uuids) > 0 {
		field := strings.Split(uuids,",")
		for _ , data := range field {
			articleType := ArticleType{Uuid:data}
			Delete(&articleType)
		}
		return SetResult("删除成功","",true,200,"")
	}
	return SetResult("删除失败","",false,405,"")
}

func (this *ArticleType) Update(data map[string]string) Result  {
	note := this.getNoteData(data)
	if UpdateData(&note){
		return SetResult("修改成功","",true,200,"")
	}
	return SetResult("修改失败","",false,400,"")
}


func (this *ArticleType) Save(data map[string]string) Result {
	articleType := this.getNoteData(data)
	if SaveData(&articleType){
		return SetResult("添加成功","",true,200,"")
	}
	return SetResult("添加失败","",false,400,"")
}

func(this *ArticleType) ListArticleType(condition map[string]string) map[string]interface{} {
	result := make(map[string]interface{});
	o := orm.NewOrm()
	qs := o.QueryTable(this.GetTableName())
	cond := orm.NewCondition()
	page ,_ := strconv.Atoi(condition["page"])
	if page < 1 {
		page = 1
	}
	if len(condition["Title"]) > 0 {
		//like %Title%
		cond = cond.And("TypeName__icontains",condition["Title"])
	}
	offset ,_ := beego.AppConfig.Int("pageSize")
	start := (page - 1) * offset
	var articleType []ArticleType
	count ,_ := qs.SetCond(cond).Count()
	num, _ := qs.SetCond(cond).RelatedSel().OrderBy("-CreateDate").Limit(offset, start).All(&articleType)
	result["limt"] = page
	result["pageSize"] = num
	result["count"] = count
	result["countPage"] = int(math.Ceil(float64(count) / float64(offset)))
	result["data"] = SetResult("成功",articleType,true,200,"")
	return result
}

func (this *ArticleType) getNoteData(data map[string]string) ArticleType{
	udata := data["data"]
	articleTypeData := ArticleType{}
	var resultArticleType ArticleType
	if len(udata) > 0 {
		json.Unmarshal([]byte(udata) , &articleTypeData);
		if len(articleTypeData.Uuid) > 0 {
			resultArticleType.Uuid = articleTypeData.Uuid
		}else{
			uuid, _ := uuid.NewV4()
			resultArticleType.Uuid  = uuid.String()
		}
		resultArticleType.TypeName = articleTypeData.TypeName
		resultArticleType.CreateDate = GetCurrentDate()
		return resultArticleType
	}
	return resultArticleType
}