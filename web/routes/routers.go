package routes

import (
	"github.com/kataras/iris/v12/_examples/mvc/login/web/middleware"
	"github.com/kataras/iris/v12/mvc"
	"lottery/bootstrap"
	"lottery/services"
	"lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	blackIpService := services.NewBlackIpService()
	resultService := services.NewResultService()
	userDayService := services.NewUserdayService()

	index := mvc.New(b.Party("/"))
	index.Register(userService,
		giftService,
		codeService,
		blackIpService,
		resultService,
		userDayService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService,
		giftService,
		codeService,
		blackIpService,
		resultService,
		userDayService)
	admin.Handle(new(controllers.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))

	adminCode := admin.Party("/code")
	adminCode.Register(codeService)
	adminCode.Handle(new(controllers.AdminCodeController))
}
