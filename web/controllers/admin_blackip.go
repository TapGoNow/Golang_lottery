package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/comm"
	"lottery/services"
)

type AdminBlackIpController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	UserdayService services.UserdayService
	BlackIpService services.BlackIpService
	ResultService  services.ResultService
	GiftService    services.GiftService
	CodeService    services.CodeService
}

func (c *AdminBlackIpController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""
	//数据列表
	datalist := c.BlackIpService.GetAll(page, size)
	total := (page-1)*size + len(datalist)
	// 数据总数
	if len(datalist) >= size {
		total = int(c.BlackIpService.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name: "admin/blackip.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "blackip",
			"Datalist": datalist,
			"Total":    total,
			"Now":      comm.NowUnix(),
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminBlackIpController) GetBlack() mvc.Result {

}
