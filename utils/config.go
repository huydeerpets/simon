package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	RunMode               string
	MYSQL_URL             string //数据库连接
	MGO_DB                string
	AdapterName           string
	BEEGO_CACHE           string      //缓存地址
	Rc                    cache.Cache //redis缓存
	Re                    error       //redis错误
	LoginVerifyCodePrefox string      // google验证码前缀
)

func init() {
	RunMode = beego.AppConfig.String("run_mode")
	config, err := beego.AppConfig.GetSection(RunMode)
	if err != nil {
		panic("配置文件读取错误 " + err.Error())
	}
	beego.Info(RunMode + "模式")
	MYSQL_URL = config["mysql_url"]
	AdapterName = config["adapter_name"]
	BEEGO_CACHE = config["beego_cache"]
	MGO_DB = config["mgo_db"]

	Rc, Re = cache.NewCache(AdapterName, BEEGO_CACHE) //初始化缓存
	if Re != nil {
		beego.Info("Redis连接失败~", "ERROR:", Re.Error())
	} else {
		beego.Info("连接:", AdapterName)
	}
}
