package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"math"
	"strings"
	"encoding/json"
	"github.com/satori/go.uuid"
)

func init() {
	//注册表
	orm.RegisterModel(new(Article))
}
type ArticleInterface interface {
	GetTableName()string
	GetArticleInfo(uuid string) Article
	ListArticle(condition map[string]string) map[string]interface{}
	Delete(data map[string]string) Result
	Update(data map[string]string) Result
	Save(data map[string]string)Result
}


//文章
type Article struct {
	Uuid         string       `orm:"size(36);pk"` //序号
	Title        string       `orm:"size(200)"`   //标题
	Content      string       `orm:"type(text)"`  //内容
	CreateDate   string       `orm:"size(50)"`    //创建时间
	Publisher    string       `orm:"size(32)"`    //发布者
	Forward      int          //转发次数
	ArticleType  *ArticleType `orm:"rel(fk)"`   //文章类型
	CreateUser   string       `orm:"size(100)"` //创建人
	SetTop       bool         //是否置顶
	PrioritySort string       `orm:"size(10)"` //优先排序输出 在不置顶的时候有用 或者大家都在置顶的时候安装 当前输出
	User         *User        `orm:"rel(fk)"`  //设置一对多关系
}

func (this *Article) GetTableName()string{
	return "Article";
}

func (this *Article) GetArticleInfo(uuid string) Article  {
	article := Article{Uuid:uuid}
	QueryObject(&article);
	return article;
}

func(this *Article) ListArticle(condition map[string]string) map[string]interface{} {
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
	var articles []Article
	count ,_ := qs.SetCond(cond).Count()
	num, _ := qs.SetCond(cond).RelatedSel().OrderBy("-CreateDate").Limit(offset, start).All(&articles)
	result["limt"] = page
	result["pageSize"] = num
	result["count"] = count
	result["countPage"] = int(math.Ceil(float64(count) / float64(offset)))
	result["data"] = SetResult("成功",articles,true,200,"")
	return result
}

func (this *Article) Delete(data map[string]string) Result  {
	uuids := data["uuids"]
	if len(uuids) > 0 {
		field := strings.Split(uuids,",")
		for _ , data := range field {
			article := Article{Uuid:data}
			Delete(&article)
		}
		return SetResult("删除成功","",true,200,"")
	}
	return SetResult("删除失败","",false,405,"")
}



func (this *Article) Update(data map[string]string) Result  {
	article := this.getArticleData(data)
	if UpdateData(&article){
		return SetResult("修改成功","",true,200,"")
	}
	return SetResult("修改失败","",false,400,"")
}

//添加文章
func (this *Article) Save(data map[string]string)Result {
	article := this.getArticleData(data)
	if SaveData(&article){
		return SetResult("添加成功","",true,200,"")
	}
	return SetResult("添加失败","",false,400,"")
}

//获取添加或者修改数据 组装成Article对象
func (this *Article) getArticleData(data map[string]string)Article{
	udata := data["data"]
	articleData := Article{}
	var resultArticle Article
	if len(udata) > 0 {
		json.Unmarshal([]byte(udata) , &articleData);
		if len(articleData.Uuid) > 0 {
			resultArticle.Uuid = articleData.Uuid
		}else{
			uuid, _ := uuid.NewV4()
			resultArticle.Uuid  = uuid.String()
		}
		resultArticle.Title = articleData.Title
		resultArticle.Content = articleData.Content
		resultArticle.CreateDate = GetCurrentDate()
		resultArticle.Publisher  = articleData.Publisher
		resultArticle.Forward    = 0
		resultArticle.ArticleType = &ArticleType{Uuid:articleData.ArticleType.Uuid}
		resultArticle.CreateUser = articleData.CreateUser
		resultArticle.User = &User{Uuid:articleData.User.Uuid}
		return resultArticle
	}
	return resultArticle
}
