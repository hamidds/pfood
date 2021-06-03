package model

import (
	"time"
)

type Restaurant struct {
	Name          string    `json:"name"             bson:"name"            validate:"required"`
	Address       string    `json:"address"          bson:"address"         validate:"required"`
	District      int       `json:"district"         bson:"district"        validate:"required"`
	Districts     []int     `json:"districts"        bson:"districts"       validate:"required"`
	DeliveryPrice int       `json:"delivery_price"   bson:"delivery_price"  validate:"required"`
	Foods         []*Food   `json:"foods"            bson:"foods"           validate:"required"`
	DeliveryTime  time.Time `json:"delivery_time"    bson:"delivery_time"   validate:"required"`
	OpenTime      time.Time `json:"open_time"        bson:"open_time"       validate:"required"`
	CloseTime     time.Time `json:"close_time"       bson:"close_time"      validate:"required"`
	PendingOrders []*Order  `json:"pending_orders"   bson:"pending_orders"`
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
