package controllers

import (
	"encoding/json"
	"simon/models"
	"simon/utils"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	config map[string]string
}

func (this *BaseController) Prepare() {

	//redis cache client

	//配置信息
	var cache_time int64
	config := make(map[string]string)
	if utils.Rc.IsExist("configs") {
		cache_content := string(utils.Rc.Get("configs").([]uint8))
		json.Unmarshal([]byte(cache_content), &config)
		cache_time = utils.StringToInt64(config["web_cache_time"])
	} else {
		config = models.GetConfigs()
		cache_time = utils.StringToInt64(config["web_cache_time"])
		if str, err := json.Marshal(config); err == nil {
			utils.Rc.Put("configs", string(str), time.Duration(cache_time)*time.Second)
		}
	}
	this.config = config

	//分类列表
	var allCategory []models.Category
	if utils.Rc.IsExist("allCategory") {
		cache_content := string(utils.Rc.Get("allCategory").([]uint8))
		json.Unmarshal([]byte(cache_content), &allCategory)
	} else {
		allCategory = models.GetCategoryList()
		if str, err := json.Marshal(allCategory); err == nil {
			utils.Rc.Put("allCategory", string(str), time.Duration(cache_time)*time.Second)
		}
	}

	//侧边栏
	var latest, hot []models.Article
	if utils.Rc.IsExist("latest") {
		cache_content := string(utils.Rc.Get("latest").([]uint8))
		json.Unmarshal([]byte(cache_content), &latest)
	} else {
		latest, _ = models.GetLatestArticles(8, 0)
		if str, err := json.Marshal(latest); err == nil {
			utils.Rc.Put("latest", string(str), time.Duration(cache_time)*time.Second)
		}
	}

	if utils.Rc.IsExist("hot") {
		cache_content := string(utils.Rc.Get("hot").([]uint8))
		json.Unmarshal([]byte(cache_content), &hot)
	} else {
		hot = models.GetTopViewArticles()
		if str, err := json.Marshal(hot); err == nil {
			utils.Rc.Put("hot", string(str), time.Duration(cache_time)*time.Second)
		}
	}

	var tags []map[string]int64
	if utils.Rc.IsExist("tags") {
		cache_content := string(utils.Rc.Get("tags").([]uint8))
		json.Unmarshal([]byte(cache_content), &tags)
	} else {
		tags = models.GetArticleTags()
		if str, err := json.Marshal(tags); err == nil {
			utils.Rc.Put("tags", string(str), time.Duration(cache_time)*time.Second)
		}
	}

	//模板变量
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["category"] = allCategory
	this.Data["latest"] = latest
	this.Data["hot"] = hot
	this.Data["tags"] = tags
	this.Data["configs"] = config
}

//是否需要模板
func (c *BaseController) IsNeedTemplate() {
	pushstate := c.GetString("pushstate")
	if pushstate != "1" {
		c.Layout = "layout/layout.html"
	}
}
