package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hamidds/pfood/model"
	"github.com/hamidds/pfood/store"
	"io/ioutil"
	"net/http"
)

var RestaurantStore *store.RestaurantStore

func AddRestaurant(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var restaurant *model.Restaurant
	err = json.Unmarshal(reqBody, &restaurant)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to restaurant fields
	err = RestaurantValidate(restaurant)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Check if Manager doesn't exist
	currentManager, err := ManagerStore.GetByEmail(params["email"]) // jwt needed
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("manager doesn't exist")))
		return
	}

	//if _, err := RestaurantStore.GetFoodByNameFromRestaurant(currentManager, restaurant.Name); err == nil {
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	json.NewEncoder(writer).Encode(model.NewError(errors.New("this food is already added")))
	//	return
	//}

	writer.WriteHeader(http.StatusOK)

	NewRestaurant := model.NewRestaurant()
	NewRestaurant.SetFields(restaurant)
	currentManager.SetRestaurant(NewRestaurant)

	// Add Food to Foods Database
	err = RestaurantStore.Create(NewRestaurant)
	if err != nil {
		return
	}

	// Update Restaurant at Database
	err = ManagerStore.Update(currentManager, currentManager.Email)
	if err != nil {
		return
	}
	response := model.NewRestaurantResponse(NewRestaurant)
	json.NewEncoder(writer).Encode(response)
}

func UpdateRestaurant(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var r *model.Restaurant
	err = json.Unmarshal(reqBody, &r)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	err = RestaurantValidate(r)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Check manager
	manager, err := ManagerStore.GetByEmail(params["email"]) // jwt needed
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Check for having a restaurant
	restaurant, err := RestaurantStore.GetByName(manager.Restaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("manager " + params["email"] + "doesn't have any restaurant")))
		return
	}

	restaurant.SetFields(r)
////////////////////////////////////////////////////////////////////////////
	var newFoods []*model.Food
	for _, food := range restaurant.Foods {
		food.Restaurant = restaurant
		err := FoodStore.UpdateFood(food)
		if err != nil {
			json.NewEncoder(writer).Encode(model.NewError(err))
			return
		}
		newFoods = append(newFoods, food)
	}
	restaurant.Foods = newFoods
/////////////////////////////////////////////////////////////////////////////

	// Update Restaurants Database
	err = RestaurantStore.UpdateRestaurant(r, restaurant.Name)
	if err != nil {
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Update Managers Database
	manager.SetRestaurant(restaurant)
	err = ManagerStore.Update(manager, manager.Name)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewRestaurantResponse(restaurant)
	json.NewEncoder(writer).Encode(response)

}

func GetRestaurantOfManager(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	manager, err := ManagerStore.GetByEmail(params["email"]) // jwt needed
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewRestaurantResponse(manager.Restaurant)
	json.NewEncoder(writer).Encode(response)
}

//func GetRestaurantByName(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(request)
//
//	manager, err := ManagerStore.GetByEmail(params["rname"]) // jwt
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(writer).Encode(model.NewError(err))
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//	response := model.NewRestaurantResponse(manager.Restaurant)
//	json.NewEncoder(writer).Encode(response)
//}


