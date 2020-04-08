package controller

import (
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-demo/models"
	"iris-demo/services"
	"iris-demo/datasource"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.MovieService
}

func (c *IndexController) Get() mvc.Result {

	datalist := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

/**
 * 方法名带By，支持路径参数
 * Url：http://localhost:8080/index/movie/1
 * Type：GET
 **/
func (c *IndexController) GetMovieBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info":  data,
		},
	}
}

/**
 * Url：http://localhost:8080/index/search?country=指环王
 * Type：GET
 **/
func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.Instance().ClearCache(&models.Movie{})
	if err != nil {
		log.Fatal(err)
	}
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}