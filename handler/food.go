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

var FoodStore *store.FoodStore

func AddFood(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var food *model.Food
	err = json.Unmarshal(reqBody, &food)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = FoodValidate(food)
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

	if _, err := RestaurantStore.GetFoodByNameFromRestaurant(currentRestaurant, food.Name); err == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("this food is already added")))
		return
	}

	writer.WriteHeader(http.StatusOK)

	NewFood := model.NewFood()
	NewFood.SetFields(food)
	NewFood.Restaurant = currentRestaurant
	currentRestaurant.AddFood(NewFood)

	// Add Food to Foods Database
	err = FoodStore.Create(NewFood)
	if err != nil {
		return
	}

	// Update Restaurant at Database
	err = RestaurantStore.UpdateRestaurant(currentRestaurant, currentRestaurant.Name)
	if err != nil {
		return
	}
	response := model.NewFoodResponse(NewFood)
	json.NewEncoder(writer).Encode(response)
}

func DeleteFood(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	restaurant, err := RestaurantStore.GetByName(params["rname"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("restaurant doesn't exist")))
		return
	}

	food, err := RestaurantStore.GetFoodByNameFromRestaurant(restaurant, params["fname"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("food " + params["fname"] + "doesn't exist")))
		return
	}

	// Update Foods Database
	err = FoodStore.Remove("name", params["fname"])
	if err != nil {
		return
	}

	// Update Restaurants Database
	restaurant.RemoveFood(food)
	err = RestaurantStore.UpdateRestaurant(restaurant, restaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	//response := model.NewCoinResponse(coin.Name, coin.Symbol, coin.Amount, coin.Rate, 200, "Coin deleted successfully!")
	response := model.NewFoodResponse(food)
	json.NewEncoder(writer).Encode(response)
}

func GetFoods(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var foods *[]model.Food

	foods, err := FoodStore.GetFoodsList()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewFoodsResponse(foods)
	json.NewEncoder(writer).Encode(response)
}

func setupResponse(writer *http.ResponseWriter, request *http.Request) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	(*writer).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*writer).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

var GetFoodsss = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	setupResponse(&writer, request)
	fmt.Println("GetFoodss")
	writer.Header().Set("Content-Type", "application/json")
	var foods []model.Food

	food := *model.NewFood()
	food.Name = "Qorme Sabzi"
	food.Price = 10000
	foods = append(foods, food)

	food2 := *model.NewFood()
	food2.Name = "Qeyme"
	food2.Price = 1022
	foods = append(foods, food2)

	writer.WriteHeader(http.StatusOK)
	response := model.NewFoodsResponse(&foods)
	json.NewEncoder(writer).Encode(response)
})

func GetFoodss(writer http.ResponseWriter, request *http.Request) {
	setupResponse(&writer, request)
	fmt.Println("GetFoodss")
	writer.Header().Set("Content-Type", "application/json")
	var foods []model.Food

	food := *model.NewFood()
	food.Name = "Qorme Sabzi"
	food.Price = 10000
	foods = append(foods, food)

	food2 := *model.NewFood()
	food2.Name = "Qeyme"
	food2.Price = 1022
	foods = append(foods, food2)

	writer.WriteHeader(http.StatusOK)
	response := model.NewFoodsResponse(&foods)
	json.NewEncoder(writer).Encode(response)
}

func GetFoodsByName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var foods *[]model.Food
	params := mux.Vars(request)

	foods, err := FoodStore.GetFoodsListByName(params["fname"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewFoodsResponse(foods)
	json.NewEncoder(writer).Encode(response)
}

//func GetFoodsByDistrict(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//	var foods *[]model.Food
//
//	foods, err := FoodStore.GetFoodsList()
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//	response := model.NewFoodsResponse(foods)
//	json.NewEncoder(writer).Encode(response)
//}

func UpdateFood(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	// Check if restaurant doesn't exist
	restaurant, err := RestaurantStore.GetByName(params["rname"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("restaurant " + params["rname"] + "doesn't exist")))
		return
	}

	// Check if food doesn't exist
	food, err := RestaurantStore.GetFoodByNameFromRestaurant(restaurant, params["fname"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("food " + params["fname"] + "doesn't exist")))
		return
	}

	// Update Foods Database
	food.SetAvailable(params["available"] == "true")
	err = FoodStore.UpdateFood(food)
	if err != nil {
		return
	}

	// Update Restaurants Database
	restaurant.ReplaceFood(food)
	err = RestaurantStore.UpdateRestaurant(restaurant, restaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewFoodResponse(food)
	json.NewEncoder(writer).Encode(response)

}
