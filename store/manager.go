package store

import (
	"context"
	"github.com/hamidds/pfood/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManagerStore struct {
	db *mongo.Collection
}

func NewManagerStore(db *mongo.Collection) *ManagerStore {
	return &ManagerStore{db: db}
}

func (ms *ManagerStore) Create(manager *model.Manager) error {
	_, err := ms.db.InsertOne(context.TODO(), manager)
	return err
}

func (ms *ManagerStore) Remove(field, value string) error {
	_, err := ms.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (ms *ManagerStore) Update(manager *model.Manager, email string) error {
	old, err := ms.GetByEmail(email)
	if err != nil {
		return err
	}
	_, err = ms.db.UpdateOne(context.TODO(),
		bson.M{"email": old.Email},
		bson.M{"$set": bson.M{
			"email":      old.Email,
			"name":       manager.Name,
			"password":   manager.Password,
			"restaurant": manager.Restaurant,
		},
		})
	return err
}

func (ms *ManagerStore) GetByEmail(email string) (*model.Manager, error) {
	var manager model.Manager
	err := ms.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&manager)
	return &manager, err
}
