package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name     string `form:"name"`
	Address  string `from:"address"`
}

func main()  {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context)  {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("===== Only Bind By Query String =====")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(http.StatusOK, "Success")
}

