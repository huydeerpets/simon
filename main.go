package main

import (
	_ "simon/routers"
	"simon/utils"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func main() {
	beego.Run()

}

func init() {
	orm.RegisterDataBase("default", "mysql", utils.MYSQL_URL)
	//Init template function
	beego.AddFuncMap("substring", utils.SubString)
	beego.AddFuncMap("is_active", utils.IsActive)
	beego.AddFuncMap("version", utils.GetStaticVersion)
	beego.AddFuncMap("base_url", utils.GetBaseUrl)
}
