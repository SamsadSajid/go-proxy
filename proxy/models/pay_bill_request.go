package models

type PayBillRequest struct {
	UserId string `json:"userId"`
	Amount string `json:"amount"`
}
