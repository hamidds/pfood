package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
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

func CustomerPhoneCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	fmt.Println(params["number"])
	_, err := CustomerStore.GetByPhone(params["number"])
	if err != nil {
		writer.WriteHeader(http.StatusOK)
		return
	} else{
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("phonenumber already exist")))
		return
	}

}

func CustomerSignUp(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Customer sign up")
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
	fmt.Println( "customer phone number: " + customer.PhoneNumber)
	fmt.Println( "customer Password: " + customer.Password)

	// Errors which are related to customer fields
	err = CustomerValidate(customer)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}
	fmt.Println("Customer Fields Validated")

	// Username already exists!
	if _, err := CustomerStore.GetByPhone(customer.PhoneNumber); err == nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("phonenumber already exists")))
		fmt.Println("phonenumber already exists")
		return
	}
	fmt.Println("Username Checked")

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
	fmt.Println("Response Sent")
}

func UpdateCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var customer model.Customer
	err = json.Unmarshal(reqBody, &customer)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	currentCustomer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("customer doesn't exist")))
		return
	}

	//if _, err := CustomerStore.GetByPhone(customer.PhoneNumber); err == nil && params["phone_number"] != customer.PhoneNumber {
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	json.NewEncoder(writer).Encode(model.NewError(errors.New("phonenumber is already taken")))
	//	return
	//}

	err = CustomerStore.UpdateProfile(currentCustomer, params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	//json.NewEncoder(writer).Encode(response)
}
