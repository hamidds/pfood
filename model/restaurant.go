package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Restaurant struct {
	Name          string        `json:"name"             bson:"name"            validate:"required"`
	Address       string        `json:"address"          bson:"address"         validate:"required"`
	District      int           `json:"district"         bson:"district"        validate:"required"`
	Districts     []int         `json:"districts"        bson:"districts"       validate:"required"`
	DeliveryPrice int           `json:"delivery_price"   bson:"delivery_price"  validate:"required"`
	DeliveryTime  time.Duration `json:"delivery_time"    bson:"delivery_time"   validate:"required"`
	OpenTime      time.Time     `json:"open_time"        bson:"open_time"       validate:"required"`
	CloseTime     time.Time     `json:"close_time"       bson:"close_time"      validate:"required"`
	Foods         []*Food       `json:"foods"            bson:"foods"           validate:"required"`
	PendingOrders []*Order      `json:"pending_orders"   bson:"pending_orders"`
}

func NewRestaurant() *Restaurant {
	var restaurant Restaurant
	restaurant.Name = ""
	restaurant.Address = ""
	restaurant.District = -1
	restaurant.Districts = []int{}
	restaurant.DeliveryPrice = -1
	restaurant.DeliveryTime = 1 * time.Hour
	restaurant.OpenTime = time.Now()
	restaurant.CloseTime = time.Now()
	return &restaurant
}

func (restaurant *Restaurant) SetFields(r *Restaurant) {
	restaurant.Name = r.Name
	restaurant.Address = r.Address
	restaurant.District = r.District
	restaurant.Districts = r.Districts
	restaurant.DeliveryPrice = r.DeliveryPrice
	restaurant.DeliveryTime = r.DeliveryTime
	restaurant.OpenTime = r.OpenTime
	restaurant.CloseTime = r.CloseTime
}

func (restaurant *Restaurant) AddFood(food *Food) {
	// Add Food
	restaurant.Foods = append(restaurant.Foods, food)
	// Update DB ?
}

func (restaurant *Restaurant) RemoveFood(food *Food) {
	var newFoods []*Food
	for _, f := range restaurant.Foods {
		if f.Name != food.Name {
			newFoods = append(newFoods, f)
		}
	}
	restaurant.Foods = newFoods
}

func (restaurant *Restaurant) ReplaceFood(food *Food) {
	var newFoods []*Food
	for _, f := range restaurant.Foods {
		if f.Name != food.Name {
			newFoods = append(newFoods, f)
		} else {
			newFoods = append(newFoods, food)
		}
	}
	restaurant.Foods = newFoods
}

func (restaurant *Restaurant) UpdateFoods() {
	var newFoods []*Food
	for _, food := range restaurant.Foods {
		food.Restaurant = restaurant
		newFoods = append(newFoods, food)
	}
	restaurant.Foods = newFoods
}

func (restaurant *Restaurant) AddOrder(order *Order) {
	// Add Food
	restaurant.PendingOrders = append(restaurant.PendingOrders, order)
	// Update DB ?
}

func (restaurant *Restaurant) GetOrder(id primitive.ObjectID) *Order {
	for _, o := range restaurant.PendingOrders {
		if o.ID == id {
			return o
		}
	}
	return nil
}

func (restaurant *Restaurant) ReplaceOrder(order *Order) {
	var newPendingOrders []*Order
	for _, c := range restaurant.PendingOrders {
		if c.ID != order.ID {
			newPendingOrders = append(newPendingOrders, c)
		} else {
			newPendingOrders = append(newPendingOrders, order)
		}
	}
	restaurant.PendingOrders = newPendingOrders
}
