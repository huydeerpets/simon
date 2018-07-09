package controllers

import (
	"io/ioutil"

	"github.com/astaxie/beego"
)

type AvController struct {
	BaseController
}

// @router /av/getav [get]
func (c *AvController) GetAv() {
	video, err := ioutil.ReadFile("./static/av/docker.mp4")
	if err != nil {
		beego.Info("======>", err)

	}
	c.Ctx.Output.Header("Content-Type", "video/mp4")
	c.Data["json"] = map[string]interface{}{"ret": 200, "msg": "get av", "data": video}
	c.ServeJSON()
}
