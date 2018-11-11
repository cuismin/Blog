package main

import (
	_ "Blog/routers"
	"github.com/astaxie/beego"
	"Blog/models"
	"github.com/astaxie/beego/orm"
)

//引入数据模型
func init(){
	// 注册数据库
	models.RegisterDB()
	SetUpStaticFileAddress()
}

//设置静态文件地址
func SetUpStaticFileAddress()  {
	folderName := beego.AppConfig.String("folderName")
	url := beego.AppConfig.String("folderDir")
	beego.SetStaticPath(folderName,url)
}

func main() {
	showSql ,_ := beego.AppConfig.Bool("showSql");
	orm.Debug = showSql
	beego.Run()
}

