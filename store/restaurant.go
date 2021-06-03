package store

import (
	"context"
	"errors"
	"github.com/hamidds/pfood/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantStore struct {
	db *mongo.Collection
}

func NewRestaurantStore(db *mongo.Collection) *RestaurantStore {
	return &RestaurantStore{db: db}
}

func (rs *RestaurantStore) Create(restaurant *model.Restaurant) error {
	_, err := rs.db.InsertOne(context.TODO(), restaurant)
	return err
}

func (rs *RestaurantStore) Remove(field, value string) error {
	_, err := rs.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (rs *RestaurantStore) UpdateRestaurant(restaurant *model.Restaurant, name string) error {
	old, err := rs.GetByName(name)
	if err != nil {
		return err
	}
	_, err = rs.db.UpdateOne(context.TODO(),
		bson.M{"name": old.Name},
		bson.M{"$set": bson.M{
			"name":           restaurant.Name,
			"address":        restaurant.Address,
			"district":       restaurant.District,
			"districts":      restaurant.Districts,
			"delivery_price": restaurant.DeliveryPrice,
			"foods":          restaurant.Foods,
			"delivery_time":  restaurant.DeliveryTime,
			"open_time":      restaurant.OpenTime,
			"close_time":     restaurant.CloseTime,
			"pending_orders": restaurant.PendingOrders,
		},
		})
	return err
}

func (rs *RestaurantStore) GetByName(name string) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	err := rs.db.FindOne(context.TODO(), bson.M{"name": name}).Decode(&restaurant)
	return &restaurant, err
}

func (rs *RestaurantStore) GetFoodByNameFromRestaurant(restaurant *model.Restaurant, name string) (*model.Food, error) {
	for _, food := range restaurant.Foods {
		if food.Name == name {
			return food, nil
		}
	}
	return &model.Food{}, errors.New("coin doesn't exist")
}

func (rs *RestaurantStore) GetRestaurantsListByName(name string) (*[]model.Food, error) {
	var foods []model.Food
	query := bson.M{"name": name}
	res, err := rs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &foods); err != nil {
		return nil, err
	}
	return &foods, err
}

func (rs *RestaurantStore) GetRestaurantsListByDistrict(district int) (*[]model.Food, error) {
	var foods []model.Food
	query := bson.M{"district": district}
	res, err := rs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &foods); err != nil {
		return nil, err
	}
	return &foods, err
}


