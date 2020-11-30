package pay_bill_controller

import (
	"github.com/SamsadSajid/go-proxy-v1/proxy/models"
	"github.com/SamsadSajid/go-proxy-v1/proxy/pay_bill_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PayBillController(c *gin.Context)  {
	/*
	* gin is based on net/http package, and that serves each request by an individual goroutine.
	* You do not need to worry about it.
	*/

	reqBody := models.PayBillRequest{}
	err := c.BindJSON(&reqBody)

	if err != nil {
		log.Panic("Bind json err is ", err)
	}

	payBillApiResponse := pay_bill_service.MakePayBill(reqBody)

	c.JSON(http.StatusOK, gin.H {
		"data": payBillApiResponse,
	})
}
