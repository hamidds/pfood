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
	//r.HandleFunc("/customers/{phone}", handler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/signup", handler.GetCustomers).Methods("GET")
	//r.HandleFunc("/{rname}/{fname}", handler.DeleteCoin).Methods("DELETE")

	// Coin Handlers
	//r.HandleFunc("/{wname}/coins", handler.CreateCoin).Methods("POST")
	//r.HandleFunc("/{wname}", handler.GetCoins).Methods("GET")
	//r.HandleFunc("/{wname}/{symbol}", handler.UpdateCoin).Methods("PUT")
	//r.HandleFunc("/{wname}/{symbol}", handler.DeleteCoin).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
	//var customer model.Customer
	//customer.Name = "hamid"
	//var restaurant model.Restaurant
	//restaurant.Name = "Shabdiz"
	//var food model.Food
	//food.Name = "Burger"
	//food.Restaurant = &restaurant
	//var comment model.Comment
	//comment.Text = "Before"
	//comment.Food = &food
	//comment.Rating = 5
	//customer.AddComment(&comment)
	//fmt.Println(customer.Comments[0].Text)
	//comment.Text = "After"
	//fmt.Println(customer.Comments[0].Text)
	//fmt.Println(customer.Favorites)
	//fmt.Println(customer.GetOrderRates())
	//food.Name = "burger after"

}
