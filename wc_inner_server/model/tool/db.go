package tool

import (
	"blog_server/config"
	"github.com/jinzhu/gorm"

	"fmt"
	"os"
)

var DB *gorm.DB
var Associate map[string]string


func initDB()  {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	//打印错误信息
	if config.ServerConfig.Env == "Develop" {
		db.LogMode(true)
	}
	DB = db
}

func init()  {
	initDB()
}
