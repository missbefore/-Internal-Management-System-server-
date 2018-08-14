package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
)

func main()  {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/",  func(ctx iris.Context) {
		ctx.HTML("<h1>解析的html</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("这个是直接返回字符串")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

/*
尼玛这是开始访问页面
*/