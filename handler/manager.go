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

var ManagerStore *store.ManagerStore

func ManagerSignUp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var manager *model.Manager

	// Errors which are related to Json encoding
	err = json.Unmarshal(reqBody, &manager)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to customer fields
	err = ManagerValidate(manager)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Username already exists!
	if _, err := ManagerStore.GetByEmail(manager.Email); err == nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("email already exists")))
		return
	}

	//newCustomer := model.(wallet.Name, 0.0, []model.Coin{}, time.Now())
	err = ManagerStore.Create(manager)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewManagerResponse(manager)
	json.NewEncoder(writer).Encode(response)
}

func ManagerLogin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	var m *model.Manager

	// Errors which are related to Json encoding
	err = json.Unmarshal(reqBody, &m)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	// Errors which are related to manager fields
	err = ManagerValidate(m)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	manager, err := ManagerStore.GetByEmail(m.Email)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(err))
		return
	}

	if manager == nil {
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("wrong email")))
		return
	}

	if !manager.CheckPassword(manager.Password) {
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("wrong password")))
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := model.NewManagerResponse(manager)
	json.NewEncoder(writer).Encode(response)
}

func ManagerEmailCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	fmt.Println(params["email"])

	if _, err := ManagerStore.GetByEmail(params["email"]); err != nil {
		writer.WriteHeader(http.StatusOK)
		return
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.NewError(errors.New("email already exists")))
		return
	}

}
