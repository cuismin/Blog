package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//mylsq 操作
func RegisterDB() {
	user := beego.AppConfig.String("user")
	pwd := beego.AppConfig.String("pwd")
	url := beego.AppConfig.String("url")
	database := beego.AppConfig.String("database")
	prot := beego.AppConfig.String("prot")
	maxIdleConns ,_ := beego.AppConfig.Int("maxIdleConns")
	maxOpenConns ,_ := beego.AppConfig.Int("maxOpenConns")
	createDataBase ,_ := beego.AppConfig.Bool("createDataBase")
	//createDataBase = true
	//注册驱动
	orm.RegisterDriver("mysql",orm.DRMySQL);
	//设置连接
	orm.RegisterDataBase("default", "mysql", user+":"+pwd+"@tcp("+url+":"+prot+")/"+database+"?charset=utf8", maxIdleConns, maxOpenConns);
	//创建表  force 如果存在更新  verbose 如果不存在就创建
	orm.RunSyncdb("default",createDataBase,true);
}