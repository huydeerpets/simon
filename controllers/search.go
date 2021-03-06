package controllers

import (
	"simon/models"
)

type SearchController struct {
	BaseController
}

// @router /search [get]
func (this *SearchController) Get() {
	//文章列表
	var pageSize int = 6
	page, err := this.GetInt("page") //获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetSearchArticles(this.GetString("s"), pageSize, (page-1)*pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/search?s="+this.GetString("s"))

	//模板变量
	this.Data["s"] = this.GetString("s")
	this.Data["article"] = articles
	this.Data["page"] = pages.Show()
	this.Layout = "frontend/layout/2columns-right.html"
	this.TplName = "frontend/article/list.html"
}
