package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hamidds/pfood/model"
	"io/ioutil"
	"net/http"
)

func AddOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var order *model.Order
	err = json.Unmarshal(reqBody, &order)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = OrderValidate(order)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Check if Restaurant doesn't exist
	currentRestaurant, err := RestaurantStore.GetByName(params["rname"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("restaurant doesn't exist")))
		return
	}

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)

	NewOrder := model.NewOrder()
	NewOrder.SetFields(order)
	NewOrder.Customer = customer
	customer.AddOrder(NewOrder)

	// Add order to User and Update Database
	err = CustomerStore.UpdateProfile(customer, params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Update Restaurant at Database
	currentRestaurant.AddOrder(NewOrder)
	err = RestaurantStore.UpdateRestaurant(currentRestaurant, currentRestaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	response := model.NewOrderResponse(NewOrder)
	json.NewEncoder(writer).Encode(response)
}

func UpdateOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var reply *model.CommentReply
	err = json.Unmarshal(reqBody, &reply)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = ReplyValidate(reply)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Check if restaurant doesn't exist
	restaurant, err := RestaurantStore.GetByName(params["rname"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("restaurant " + params["rname"] + "doesn't exist")))
		return
	}

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	order := customer.GetOrder(reply.CommentID)
	order.SetState(reply.Answer)

	// Update Comment in User and Update Database
	customer.ReplaceOrder(order)
	err = CustomerStore.UpdateProfile(customer, params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Update Restaurants Database
	restaurant.ReplaceOrder(order)
	err = RestaurantStore.UpdateRestaurant(restaurant, restaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewOrderResponse(order)
	json.NewEncoder(writer).Encode(response)

}

func GetOrdersHistory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	orders := customer.GetOrdersHistory()
	writer.WriteHeader(http.StatusOK)
	response := model.NewOrdersResponse(orders)
	json.NewEncoder(writer).Encode(response)
}

func GetOrderByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var reply *model.CommentReply
	err = json.Unmarshal(reqBody, &reply)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = ReplyValidate(reply)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	order := customer.GetOrder(reply.CommentID)
	// Update DB Again?


	writer.WriteHeader(http.StatusOK)
	response := model.NewOrderResponse(order)
	json.NewEncoder(writer).Encode(response)
}
