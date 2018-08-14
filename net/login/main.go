package main

import (
	"github.com/kataras/iris"
	"net/login/datasource"
	"net/login/repositories"
	"net/login/services"
	"github.com/kataras/iris/mvc"
	"net/login/web/middleware"
	"github.com/kataras/iris/sessions"
	"time"

	"net/login/web/controller"
)

func main()  {
	app := iris.New()

	app.Logger().SetLevel("debug")

	template := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(template)

	app.StaticWeb("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "这里的基本上不是你想看到的主页"))
		ctx.View("shared/error.html")
	})

	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		app.Logger().Fatal("error while loading the users: %v", err)
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	users := mvc.New(app.Party("/users"))
	users.Router.Use(middleware.BasicAuth)
	users.Register(userService)
	users.Handle(new(controller.UserController))

	sessionManager := sessions.New(sessions.Config{
		Cookie: "session_cookie_name",
		Expires: 24 * time.Hour,
	})
	user :=  mvc.New(app.Party("/user"))
	user.Register(
		userService,
		sessionManager.Start,
	)
	user.Handle(new(controller.UserController))

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}

