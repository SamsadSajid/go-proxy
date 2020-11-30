package main

import (
	"github.com/SamsadSajid/go-proxy-v1/proxy/pay_bill_controller"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/api/pay_bill", pay_bill_controller.PayBillController)

	router.Run("127.0.0.1:8080")
}
