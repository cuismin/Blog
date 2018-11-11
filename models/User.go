package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"math"
)

func init() {
	//注册表
	orm.RegisterModel(new(User))
}

type User struct {
	Uuid         string     `orm:"size(36);pk"` //序号
	UserName     string     `orm:"size(100)"`   //用户名
	Password     string     `orm:"size(100)"`   //用户密码
	RealName     string     `orm:"size(100)"`   //真实姓名
	Nickname     string     `orm:"size(100)"`   //昵称
	HeadPortrait string     `orm:"size(200)"`   //头像
	Sex          byte       // 0|1 ? 男 | 女
	Age          int        //年龄
	CreateDate   string     `orm:"size(50)"`      //创建时间
	Article      []*Article `orm:"reverse(many)"` //设置一对多方向关系
}

type UserInterface interface {
	GetTableName() string
	GetUserInfo(uuid string) User
	ListUser(condition map[string]string) map[string]interface{}
}

func (this *User) GetTableName() string {
	return "User"
}

func (this *User) GetUserInfo(uuid string) User  {
	user := User{Uuid:uuid}
	QueryObject(&user);
	return user;
}

func(this *User) ListUser(condition map[string]string) map[string]interface{} {
	result := make(map[string]interface{});
	o := orm.NewOrm()
	qs := o.QueryTable(this.GetTableName())
	cond := orm.NewCondition()
	page ,_ := strconv.Atoi(condition["page"])
	if page < 1 {
		page = 1
	}
	if len(condition["UserName"]) > 0 {
		//like %Title%
		cond = cond.And("UserName__icontains",condition["UserName"])
	}
	offset ,_ := beego.AppConfig.Int("pageSize")
	start := (page - 1) * offset
	var users []User
	count ,_ := qs.SetCond(cond).Count()
	num, _ := qs.SetCond(cond).RelatedSel().Limit(offset, start).All(&users)
	result["limt"] = page
	result["pageSize"] = num
	result["count"] = count
	result["countPage"] = int(math.Ceil(float64(count) / float64(offset)))
	result["data"] = SetResult("成功",users,true,200,"")
	return result
}
