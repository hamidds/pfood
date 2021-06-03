package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	PATH = "mongodb://localhost:27017"
	//PATH = "mongodb+srv://hamidds:PjjjLCZMR5Q6qCLE@cluster0.gaykr.mongodb.net/restapi?retryWrites=true&w=majority"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(PATH)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		} else {
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				clientInstanceError = err
			}
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func SetupRestaurantsDb(mongoClient *mongo.Client) *mongo.Collection {
	restaurantsDb := mongoClient.Database("pfood_db").Collection("restaurants")
	createUniqueIndices(restaurantsDb, "name")
	return restaurantsDb
}

func SetupFoodsDb(mongoClient *mongo.Client) *mongo.Collection {
	FoodsDb := mongoClient.Database("pfood_db").Collection("foods")
	//createUniqueIndices(FoodsDb, "name")
	return FoodsDb
}

func SetupCustomerDb(mongoClient *mongo.Client) *mongo.Collection {
	customersDb := mongoClient.Database("pfood_db").Collection("customers")
	createUniqueIndices(customersDb, "phone_number")
	return customersDb
}

func SetupManagerDb(mongoClient *mongo.Client) *mongo.Collection {
	managersDb := mongoClient.Database("pfood_db").Collection("managers")
	createUniqueIndices(managersDb, "email")
	return managersDb
}


func createUniqueIndices(db *mongo.Collection, field string) {
	_, err := db.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: field, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
