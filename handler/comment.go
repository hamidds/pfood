package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hamidds/pfood/model"
	"io/ioutil"
	"net/http"
)

func AddComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to Json encoding
	var comment *model.Comment
	err = json.Unmarshal(reqBody, &comment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = CommentValidate(comment)
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

	currentFood, err := RestaurantStore.GetFoodByNameFromRestaurant(currentRestaurant, params["fname"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("this food doesn't exist")))
		return
	}

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)

	NewComment := model.NewComment()
	NewComment.SetFields(comment)
	NewComment.Food = currentFood
	customer.AddComment(NewComment)
	currentFood.AddComment(NewComment)

	// Add Comment to User and Update Database
	err = CustomerStore.UpdateProfile(customer, params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Add Food to Foods Database
	err = FoodStore.UpdateFood(currentFood)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Update Restaurant at Database
	currentRestaurant.ReplaceFood(currentFood)
	err = RestaurantStore.UpdateRestaurant(currentRestaurant, currentRestaurant.Name)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}
	response := model.NewCommentResponse(NewComment)
	json.NewEncoder(writer).Encode(response)
}

func UpdateComment(writer http.ResponseWriter, request *http.Request) {
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

	// Check if food doesn't exist
	food, err := RestaurantStore.GetFoodByNameFromRestaurant(restaurant, params["fname"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("food " + params["fname"] + "doesn't exist")))
		return
	}

	// Check if comment doesn't exist
	comment := food.GetComment(reply.CommentID)
	if comment == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("comment " + reply.CommentID.String() + "doesn't exist")))
		return
	}

	customer, err := CustomerStore.GetByPhone(params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	comment.SetAnswer(reply.Answer)
	customer.ReplaceComment(comment)

	// Update Comment in User and Update Database
	err = CustomerStore.UpdateProfile(customer, params["phone_number"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Update Foods Database
	food.ReplaceComment(comment)
	err = FoodStore.UpdateFood(food)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
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
	response := model.NewCommentResponse(comment)
	json.NewEncoder(writer).Encode(response)

}

func GetComments(writer http.ResponseWriter, request *http.Request) {
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

	comments := food.GetComments()
	writer.WriteHeader(http.StatusOK)
	response := model.NewCommentsResponse(comments)
	json.NewEncoder(writer).Encode(response)
}
