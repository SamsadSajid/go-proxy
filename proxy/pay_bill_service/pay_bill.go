package pay_bill_service

import (
	"bytes"
	"encoding/json"
	"github.com/SamsadSajid/go-proxy-v1/proxy/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func MakePayBill(reqBody models.PayBillRequest) models.PayBillApiResponse{
	log.Print("Pay Bill request body: ", reqBody)

	req, _ := getReqBody(reqBody)

	client := http.Client{Timeout: time.Second * 30}

	request, err := http.NewRequest(
		"POST",
		"http://127.0.0.1:8082/api/pay-bill",
		bytes.NewReader(req),
	)

	if err != nil {
		log.Panic("Error forming request")
	}

	request.Header.Add("Connection", "close")
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)

	if err != nil {
		log.Panic("Error from pay bill server is ", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Panic("Error in converting response to byte ", err)
	}

	payBillResponse := models.PayBillApiResponse{}
	err = json.Unmarshal(body, &payBillResponse)

	if err != nil {
		log.Panic("Error parsing response ", err)
	}

	log.Println("Res is ", payBillResponse)

	return payBillResponse
}

func getReqBody(body models.PayBillRequest) ([]byte, error) {
	reqBody := models.PayBillRequest{}

	reqBody = models.PayBillRequest{
		UserId: body.UserId,
		Amount: body.Amount,
	}

	return json.Marshal(reqBody)
}
