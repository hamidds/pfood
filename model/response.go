package model

import "github.com/hamidds/pfood/utils"

type CustomerResponse struct {
	Customer struct {
		PhoneNumber  string     `json:"phone_number"   bson:"_id"`
		Name         string     `json:"name"           bson:"name"`
		Credit       int        `json:"credit"         bson:"credit"`
		District     int        `json:"district"       bson:"district"`
		Address      string     `json:"address"        bson:"address"`
		OrderHistory []*Order   `json:"order_history"  bson:"order_history"`
		Favorites    []*Food    `json:"favorites"      bson:"favorites"`
		Comments     []*Comment `json:"comments"       bson:"comments"`
		Token        string     `json:"token"`
	} `json:"user"`
}

func NewUserResponse(customer *Customer) *CustomerResponse {
	response := new(CustomerResponse)
	response.Customer.PhoneNumber = customer.PhoneNumber
	response.Customer.Name = customer.Name
	response.Customer.Credit = customer.Credit
	response.Customer.District = customer.District
	response.Customer.Address = customer.Address
	response.Customer.OrderHistory = customer.OrderHistory
	response.Customer.Favorites = customer.Favorites
	response.Customer.Comments = customer.Comments
	response.Customer.Token = utils.GenerateJWT("phone_number", customer.PhoneNumber)
	return response
}