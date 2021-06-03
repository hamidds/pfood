package main

import (
	"github.com/gorilla/mux"
	"github.com/hamidds/pfood/db"
	"github.com/hamidds/pfood/handler"
	"github.com/hamidds/pfood/store"
	"log"
	"net/http"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	mongoClient, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	customersDb := db.SetupCustomerDb(mongoClient)
	handler.CustomerStore = store.NewCustomerStore(customersDb)
	//foodsDb := db.SetupFoodsDb(mongoClient)
	//handler.CustomerStore = store.NewCustomerStore(customersDb)
	//managersDb := db.SetupManagerDb(mongoClient)
	//handler.CustomerStore = store.NewCustomerStore(customersDb)
	//restaurantsDb := db.SetupRestaurantsDb(mongoClient)
	//handler.CustomerStore = store.NewCustomerStore(customersDb)

	// Customer Handlers
	r.HandleFunc("/signup", handler.CustomerSignUp).Methods("POST")
	r.HandleFunc("/login", handler.AuthMiddleware(handler.CustomerLogin)).Methods("POST")
	//r.HandleFunc("/customers/{phone_number}", handler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/signup", handler.GetCustomers).Methods("GET")

	// Food Handlers => By Manager
	r.HandleFunc("/{rname}/foods", handler.AddFood).Methods("POST")
	r.HandleFunc("/{rname}/{fname}", handler.DeleteFood).Methods("DELETE")
	r.HandleFunc("/{rname}/{fname}/{available}", handler.UpdateFood).Methods("PUT")

	// Food Filters => By Customer
	r.HandleFunc("/foods", handler.GetFoods).Methods("GET")
	r.HandleFunc("/foods/{fname}", handler.GetFoodsByName).Methods("GET")
	//r.HandleFunc("/foods/{rname}", handler.GetFoodsByRestaurant).Methods("GET")
	//r.HandleFunc("/foods/{district}", handler.GetFoodsByDistrict).Methods("GET")
	// r.HandleFunc("/foods/{rname}/{fname}", handler.GetFoodsByRestaurantAndName).Methods("GET")
	// r.HandleFunc("/foods/{fname}/{district}", handler.GetFoodsByNameAndDistrict).Methods("GET")

	// Favorite Foods
	// r.HandleFunc("/foods/favorites", handler.GetFavoriteFoods).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
