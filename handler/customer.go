package handler

import (
	"encoding/json"
	"errors"
	"github.com/hamidds/pfood/model"
	"github.com/hamidds/pfood/store"
	"io/ioutil"
	"net/http"
)

var CustomerStore *store.CustomerStore

//var CustomerSignUp1 = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	reqBody, err := ioutil.ReadAll(request.Body)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	var customer *model.Customer
//
//	// Errors which are related to Json encoding
//	err = json.Unmarshal(reqBody, &customer)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	// Errors which are related to customer fields
//	err = CustomerValidate(customer)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	// Username already exists!
//	if _, err := CustomerStore.GetByPhone(customer.PhoneNumber); err == nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(errors.New("username already exists")))
//		return
//	}
//
//	//newCustomer := model.(wallet.Name, 0.0, []model.Coin{}, time.Now())
//	err = CustomerStore.Create(customer)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//	response := model.NewUserResponse(customer)
//	json.NewEncoder(writer).Encode(response)
//})

var CustomerLogin = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var c *model.Customer

	// Errors which are related to Json encoding
	err = json.Unmarshal(reqBody, &c)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = CustomerValidate(c)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	customer, err := CustomerStore.GetByPhone(c.PhoneNumber)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	if customer == nil {
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("wrong phonenumber")))
		return
	}

	if !customer.CheckPassword(customer.Password) {
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("wrong password")))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewUserResponse(customer)
	json.NewEncoder(writer).Encode(response)
})


func GetCustomers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var customers *[]model.Customer

	customers, err := CustomerStore.GetCustomersList()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(customers)
}




func CustomerSignUp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var customer *model.Customer

	// Errors which are related to Json encoding
	err = json.Unmarshal(reqBody, &customer)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = CustomerValidate(customer)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Username already exists!
	if _, err := CustomerStore.GetByPhone(customer.PhoneNumber); err == nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("username already exists")))
		return
	}

	//newCustomer := model.(wallet.Name, 0.0, []model.Coin{}, time.Now())
	err = CustomerStore.Create(customer)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewUserResponse(customer)
	json.NewEncoder(writer).Encode(response)
}


