package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"iris-demo/web/route"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	
	mvc.Configure(app.Party("/index"), route.Index)

	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	app.StaticWeb("/public", "./web/public")

	app.Run(iris.Addr(":8080"))
}
