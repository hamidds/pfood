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


	// Manager Handlers
	r.HandleFunc("/signup/manager", handler.ManagerSignUp).Methods("POST")
	r.HandleFunc("/login/manager", handler.ManagerLogin).Methods("POST")

	// Restaurant Handlers
	r.HandleFunc("/manager", handler.AddRestaurant).Methods("POST") // jwt
	r.HandleFunc("/manager/restaurant", handler.UpdateRestaurant).Methods("POST") // jwt
	// also can be used for pending orders
	r.HandleFunc("/manager/restaurant", handler.GetRestaurantOfManager).Methods("GET") // jwt

	// Customer Handlers
	r.HandleFunc("/signup/user", handler.CustomerSignUp).Methods("POST")
	r.HandleFunc("/login/user", handler.AuthMiddleware(handler.CustomerLogin)).Methods("POST")
	r.HandleFunc("/customers", handler.UpdateCustomer).Methods("PUT") // jwt
	r.HandleFunc("/signup", handler.GetCustomers).Methods("GET")

	// Food Handlers => By Manager
	r.HandleFunc("/{rname}/foods", handler.AddFood).Methods("POST") // jwt
	r.HandleFunc("/{rname}/{fname}", handler.DeleteFood).Methods("DELETE") // jwt
	r.HandleFunc("/{rname}/{fname}/{available}", handler.UpdateFood).Methods("PUT") // jwt

	// Food Filters => By Customer
	r.HandleFunc("/foods", handler.GetFoods).Methods("GET") // jwt
	r.HandleFunc("/foods/{fname}", handler.GetFoodsByName).Methods("GET") // jwt
	// r.HandleFunc("/foods/{rname}", handler.GetFoodsByRestaurant).Methods("GET")
	// r.HandleFunc("/foods/{district}", handler.GetFoodsByDistrict).Methods("GET")
	// r.HandleFunc("/foods/{rname}/{fname}", handler.GetFoodsByRestaurantAndName).Methods("GET")
	// r.HandleFunc("/foods/{fname}/{district}", handler.GetFoodsByNameAndDistrict).Methods("GET")

	// Favorite Foods
	// r.HandleFunc("/foods/favorites", handler.GetFavoriteFoods).Methods("GET")


	// Comment Handlers
	r.HandleFunc("/{rname}/{fname}/comments", handler.AddComment).Methods("POST")
	// e.g. reply a comment
	r.HandleFunc("/{rname}/{fname}/comments", handler.UpdateComment).Methods("PUT")
	r.HandleFunc("/{rname}/{fname}/comments", handler.GetComments).Methods("GET")


	// Order Handlers
	r.HandleFunc("/{rname}/order", handler.AddOrder).Methods("POST")
	// e.g. change state by manager
	r.HandleFunc("/{rname}/order", handler.UpdateOrder).Methods("PUT")
	// Get order state
	r.HandleFunc("/customer/orders", handler.GetOrdersHistory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
