package store

import (
	"context"
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

func (rs *RestaurantStore) RemoveFood(restaurant *model.Restaurant, food *model.Food) error {
	var newFoods []*model.Food
	for _, f := range restaurant.Foods {
		if f.Name != food.Name {
			newFoods = append(newFoods, f)
		}
	}
	_, err := rs.db.UpdateOne(context.TODO(), bson.M{"name": restaurant.Name}, bson.M{"$set": bson.M{"foods": newFoods}})
	if err != nil {
		return err
	}
	restaurant.Foods = newFoods
	return nil
}
