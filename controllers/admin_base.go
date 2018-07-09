package controllers

import (
	"simon/models"

	"github.com/astaxie/beego"
)

type AdminBaseController struct {
	beego.Controller
	isAdminLogin bool
}

func (this *AdminBaseController) Prepare() {
	//后台登录信息
	var loginUser models.User
	admin_userLogin := this.GetSession("admin_userLogin")
	if admin_userLogin == nil {
		this.isAdminLogin = false
	} else {
		this.isAdminLogin = true
		loginUser = models.GetUserInfo(this.GetSession("admin_userId"))
	}

	//模板变量
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["isAdminLogin"] = this.isAdminLogin
	this.Data["loginUser"] = loginUser
}
