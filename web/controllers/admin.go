package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/services"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	UserdayService services.UserdayService
	BlackIpService services.BlackIpService
	ResultService  services.ResultService
	GiftService    services.GiftService
	CodeService    services.CodeService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}

}
