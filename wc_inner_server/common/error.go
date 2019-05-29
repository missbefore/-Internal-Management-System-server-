package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wc_inner_server/common/code"
)


func SendError(msg string, args ...interface{})  {
	if len(args) == 0 {
		panic("缺少 *gin.context")
	}

	var c *gin.Context

	var errNo = code.ErrorCode.ERROR
	if len(args) == 1 {
		ctx, ok := args[0].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}

		c = ctx
	}

	if len(args) == 2 {
		errorNo, ok := args[0].(int)
		if !ok {
			panic("error No. 错误")
		}

		errNo = errorNo
		ctx, ok := args[1].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}
		c = ctx
	}

	if c !=nil {
		c.JSON(http.StatusOK, gin.H{
			"errorCode": errNo,
			"msg" :  msg,
			"data": gin.H{},
		})

		c.Abort()
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicing %s\n", err)
		} else {
			fmt.Println("loading......")
		}
	}()
}