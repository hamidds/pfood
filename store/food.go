package store

import (
	"context"
	"github.com/hamidds/pfood/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodStore struct {
	db *mongo.Collection
}

func NewFoodStore(db *mongo.Collection) *FoodStore {
	return &FoodStore{db: db}
}

func (fs *FoodStore) Create(food *model.Food) error {
	_, err := fs.db.InsertOne(context.TODO(), food)
	return err
}

func (fs *FoodStore) Remove(field, value string) error {
	_, err := fs.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (fs *FoodStore) GetByName(name string) (*model.Food, error) {
	var food model.Food
	err := fs.db.FindOne(context.TODO(), bson.M{"name": name}).Decode(&food)
	return &food, err
}

func (fs *FoodStore) GetFoodsList() (*[]model.Food, error) {
	var foods []model.Food
	query := bson.M{}
	res, err := fs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &foods); err != nil {
		return nil, err
	}
	return &foods, err
}

func (fs *FoodStore) GetFoodsListByName(name string) (*[]model.Food, error) {
	var foods []model.Food
	query := bson.M{"name": name}
	res, err := fs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &foods); err != nil {
		return nil, err
	}
	return &foods, err
}

func (fs *FoodStore) UpdateFood(food *model.Food) error {
	_, err := fs.db.UpdateOne(context.TODO(),
		bson.M{"_id": food.ID},
		bson.M{"$set": bson.M{
			"name":       food.Name,
			"price":      food.Price,
			"available":  food.Available,
			"comments":   food.Comments,
			"rating":     food.Rating,
			"restaurant": food.Restaurant,
		},
		})
	return err
}

func (fs *FoodStore) AddComment(food *model.Food, comment *model.Comment) error {
	food.AddComment(comment)
	_, err := fs.db.UpdateOne(context.TODO(), bson.M{"_id": food.ID}, bson.M{"$set": bson.M{"comments": food.Comments}})
	if err != nil {
		return err
	}
	return nil
}
