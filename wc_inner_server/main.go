package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wc_inner_server/middleware"

	"wc_inner_server/routes"
)

var (
	engine = gin.Default()
)

func main() {
	fmt.Println("gin.Version: ", gin.Version)
	engine.Use(middleware.WebHeaderMiddle())
	routes.Router(engine)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	if err := engine.Run(":8060");err != nil {
		fmt.Println("启动服务失败", err.Error())
	}
}