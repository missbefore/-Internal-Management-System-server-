package routes

import (
	"github.com/gin-gonic/gin"
	"wc_inner_server/controller/merchant"
	"wc_inner_server/controller/payment"
	"wc_inner_server/controller/public"
	"wc_inner_server/controller/schedules"
	"wc_inner_server/middleware"
)

func Router(router *gin.Engine) {

	router.POST("/login",  public.LoginToSystem)
	router.POST("/auth", public.GetAuthInfo)
	merchantGroup := router.Group("/merchants", middleware.ValidateHttpHeaderToken())
	{
		merchantGroup.GET("/list", merchant.List)
		merchantGroup.GET("/options", merchant.GetOptionsData)
		merchantGroup.GET("/schedules", schedules.List)
	}


	paymentGroup := router.Group("/payment", middleware.ValidateHttpHeaderToken())
	{
		paymentGroup.GET("/list", payment.List)
		paymentGroup.GET("/options", payment.GetOptionsData)
	}

}