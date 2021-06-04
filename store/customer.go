package store

import (
	"context"
	"github.com/hamidds/pfood/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerStore struct {
	db *mongo.Collection
}

func NewCustomerStore(db *mongo.Collection) *CustomerStore {
	return &CustomerStore{db: db}
}

func (cs *CustomerStore) Create(customer *model.Customer) error {
	_, err := cs.db.InsertOne(context.TODO(), customer)
	return err
}

func (cs *CustomerStore) Remove(field, value string) error {
	_, err := cs.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (cs *CustomerStore) GetByPhone(phone string) (*model.Customer, error) {
	var customer model.Customer
	err := cs.db.FindOne(context.TODO(), bson.M{"phone_number": phone}).Decode(&customer)
	return &customer, err
}

func (cs *CustomerStore) GetCustomersList() (*[]model.Customer, error) {
	var customers []model.Customer
	query := bson.M{}
	res, err := cs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &customers); err != nil {
		return nil, err
	}
	return &customers, err
}

func (cs *CustomerStore) UpdateProfile(customer *model.Customer, phone string) error {
	old, err := cs.GetByPhone(phone)
	if err != nil {
		return err
	}
	_, err = cs.db.UpdateOne(context.TODO(),
		bson.M{"phone_number": old.PhoneNumber},
		bson.M{"$set": bson.M{
			"phone_number":  old.PhoneNumber,
			"password":      customer.Password,
			"name":          customer.Name,
			"credit":        customer.Credit,
			"district":      customer.District,
			"address":       customer.Address,
			"order_history": customer.OrderHistory,
			"favorites":     customer.Favorites,
			"comments":      customer.Comments,
		},
		})
	return err
}
