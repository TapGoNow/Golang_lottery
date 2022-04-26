package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"lottery/comm"
	"lottery/models"
	"lottery/services"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	UserdayService services.UserdayService
	BlackIpService services.BlackIpService
	ResultService  services.ResultService
	GiftService    services.GiftService
	CodeService    services.CodeService
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "Welcome to Go抽奖系统, <a href='/public/index.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.GiftService.GetAll(true)
	list := make([]models.LtGift, 0)
	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

func (c *IndexController) GetNewprice() map[string]interface{} {
	rs := make(map[string]interface{})
	//todo

	return rs
}

func (c *IndexController) GetLogin() {
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=login")
}

func (c *IndexController) GetLogout() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=logout")
}
