package controllers

import (
	"simon/models"
	"simon/utils"
)

type MainController struct {
	BaseController
}

// @router / [get]
func (this *MainController) Get() {
	//文章列表
	pageSize := utils.StringToInt(this.config["web_perpage"])
	page, err := this.GetInt("page") //获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetLatestArticles(pageSize, (page-1)*pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/")

	//模板变量
	this.Data["article"] = articles
	this.Data["page"] = pages.Show()
	this.Layout = "frontend/layout/2columns-right.html"
	this.TplName = "frontend/article/list.html"
}
