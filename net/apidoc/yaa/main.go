package main

import (
	"github.com/kataras/iris"

	"github.com/betacraft/yaag/yaag"
	"github.com/betacraft/yaag/irisyaag"
)

type myXML struct {
	Result string `zml:"result"`
}

func main()  {
	app := iris.New()

	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apodoc.html",
		BaseUrls: map[string]string{"Production":"", "Staging": ""},
	})
	app.Use(irisyaag.New())

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"result": "Hello single dog!"})
	})

	app.Get("/plain", func(ctx iris.Context) {
		ctx.Text("Hello all world!")
	})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(myXML{Result: "this is a xml"})
	})

	app.Get("/complex", func(ctx iris.Context) {
		value := ctx.URLParam("key")
		ctx.JSON(iris.Map{"value": value})
	})

	app.Run(iris.Addr(":8080"))
}
