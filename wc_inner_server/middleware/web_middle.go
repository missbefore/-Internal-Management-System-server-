package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wc_inner_server/common"
	"wc_inner_server/common/code"
)

var sendError = common.SendError

func WebHeaderMiddle() gin.HandlerFunc {

	return func(c *gin.Context) {
		if method := c.Request.Method;method=="OPTIONS" {
			c.Status(http.StatusOK)
		}
		origin := c.Request.Header.Get("Origin")        //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "http://do.com:3000")
			c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Authorization , Access-Control-Request-Headers")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Expose-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method ,Access-Control-Request-Headers")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Header("content-type", "application/json")       // 设置返回格式是json
		}

		c.Next()
	}
}

func ValidateHttpHeaderToken() gin.HandlerFunc  {
	
	return func(c *gin.Context) {
		if token := c.Request.Header.Get("Authorization");token != "" {
			tokenJwt, err := common.ParseToken(token)
			if err != nil {
				fmt.Println(err.Error())
				sendError(err.Error(), code.ErrorCode.TokeNotValid, c)
				return
			}
			if tokenJwt.Valid {
				fmt.Println("居然特么通过了")
				c.Next()
			} else {
				sendError("have no Authorization",  c)
				return
			}

		} else {
			sendError("have no Authorization", c)
			return
		}
	}
}

