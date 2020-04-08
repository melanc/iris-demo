package route

import (
	"iris-demo/web/controller"
	"github.com/kataras/iris/mvc"
	"iris-demo/web/middleware"
	"iris-demo/services"
)

//注意mvc.Application，它不是iris.Application。
func Index(app *mvc.Application) {
	//添加基本身份验证（admin：password）中间件
	app.Router.Use(middleware.BasicAuth)
	movieService := services.NewMovieService()
	app.Register(movieService)
	app.Handle(new(controller.IndexController))
}
