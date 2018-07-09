package controllers

import (
	"simon/models"
	"simon/utils"
)

type AdminConfigController struct {
	AdminBaseController
}

// @router /admin/config [get,post]
func (this *AdminConfigController) ConfigList() {
	if this.Ctx.Input.Method() == "GET" {
		//配置列表
		configs := models.GetListConfig()

		//模板变量
		this.Data["configs"] = configs
		this.Layout = "admin/layout/2columns-left.html"
		this.TplName = "admin/config.html"
	} else {
		config := &models.Config{}
		if err := this.ParseForm(config); err != nil {
			return
		}

		if _, err := models.InsertConfig(config); err == nil {
			//删除配置的缓存
			utils.Rc.Delete("configs")
		}
		this.Redirect("/admin/config", 302)
	}
}

// @router /admin/config/:id [post,put,delete]
func (this *AdminConfigController) ConfigUpdate() {
	if this.GetString("_method") == "DELETE" {
		id, _ := this.GetInt64(":id")
		err := models.DeleteConfig(id)
		if err == nil {
			//删除配置的缓存
			utils.Rc.Delete("configs")

			this.Redirect("/admin/config", 302)
		}
	} else {
		id, _ := this.GetInt64(":id")
		params := make(map[string]string)
		params["name"] = this.GetString("name")
		params["path"] = this.GetString("path")
		params["value"] = this.GetString("value")
		err := models.UpdateConfig(id, params)
		if err == nil {
			//删除配置的缓存
			utils.Rc.Delete("configs")

			this.Redirect("/admin/config", 302)
		}
	}
}
