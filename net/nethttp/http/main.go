package main

import (
	"github.com/kataras/iris"

	"net/http"
)

func main()  {
	app := iris.New()
	irisMiddleware := iris.FromStd(nativeTestMiddleware)
	app.Use(irisMiddleware)


	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Home</h1>")
	})

	app.Get("/ok", func(ctx iris.Context) {
		ctx.HTML("<b>Hello all friends.</b>")
	})

	app.Run(iris.Addr(":8080"))
}

func nativeTestMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)  {
	println("Request path: " + r.URL.Path)
	next(w, r)
}
