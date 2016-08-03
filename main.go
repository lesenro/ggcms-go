package main

import (
	"fmt"
	_ "ggcms/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	connDbInit()
}

//数据库初始化
func connDbInit() {
	dbuser := beego.AppConfig.String("mysqluser")
	dbpass := beego.AppConfig.String("mysqlpass")
	dburl := beego.AppConfig.String("mysqlurls")
	dbname := beego.AppConfig.String("mysqldb")
	dbport, _ := beego.AppConfig.Int("mysqlport")

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbuser, dbpass, dburl, dbport, dbname)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN)
	orm.Debug=true
}
func main() {
	beego.Run()
}
