package routes

import (
	"LuckyDraw/project/bootstrap"
	"LuckyDraw/project/services"
	"LuckyDraw/project/web/controller"
	"github.com/kataras/iris/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	index := mvc.New(b.Party("/"))
	index.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	index.Handle(new(controller.IndexController))
}
