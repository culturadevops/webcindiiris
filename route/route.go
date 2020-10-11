package route

import (
	commons "inkafarma/webcindi/common"
	controllers "inkafarma/webcindi/controller"
	"inkafarma/webcindi/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Routes(app *iris.Application) {
	//登录路由
	mvc.New(app.Party("/")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.IndexController))

	//登录路由
	mvc.New(app.Party("/login")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.LoginController))

	//系统路由
	mvc.New(app.Party("/system", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.SystemController))
	//管理员管理
	mvc.New(app.Party("/administrators", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.AdministratorsController))
	mvc.New(app.Party("/bastion", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.BastionController))
	mvc.New(app.Party("/runbook", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.RunbooksController))
}
