package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//添加数据
func SaveData(entity interface{}) bool {
	ormer := orm.NewOrm()
	_, err := ormer.Insert(entity)
	if err != nil {
		beego.Error("SaveDatas error:", err)
		return  false
	}
	return true
}

//更新数据
func UpdateData(entity interface{}) bool {
	ormer := orm.NewOrm()
	_, err := ormer.Update(entity)
	if err != nil {
		beego.Info("userUpdate err=", err)
		return  false
	}
	return true
}

//查询多条记录
func QueryAll(tableName string, modules interface{}) Result {
	ormer := orm.NewOrm()
	_, err := ormer.QueryTable(tableName).All(modules)
	if err != nil {
		beego.Error("QueryAll err:", err)
		return Result{Message:"失败",Data:nil,State:404,Success:false}
	}
	return Result{Message:"成功",Data:modules,State:200,Success:true}
}
//查询一条数据
func QueryObject( conditionModule interface{}) {
	ormer := orm.NewOrm()
	error := ormer.Read(conditionModule)
	if error != nil{
		beego.Error(error)
	}
}

func Delete(uuid interface{})bool {
	ormer := orm.NewOrm()
	if _, err := ormer.Delete(uuid); err == nil {
		return  false
	}
	return true
}

