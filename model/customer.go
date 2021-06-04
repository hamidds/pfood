package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	PhoneNumber  string     `json:"phone_number"   bson:"phone_number"       validate:"required"`
	Password     string     `json:"password"       bson:"password"           validate:"required"`
	Name         string     `json:"name"           bson:"name"                                  `
	Credit       int        `json:"credit"         bson:"credit"                                `
	District     int        `json:"district"       bson:"district"                              `
	Address      string     `json:"address"        bson:"address"                               `
	OrderHistory []*Order   `json:"order_history"  bson:"order_history"                         `
	Favorites    []*Food    `json:"favorites"      bson:"favorites"                             `
	Comments     []*Comment `json:"comments"       bson:"comments"                              `
}

func (customer *Customer) AddComment(comment *Comment) {
	customer.Comments = append(customer.Comments, comment)
	// Update Favorites
	//customer.UpdateFavorites()
}

func (customer *Customer) AddOrder(order *Order) {
	customer.OrderHistory = append(customer.OrderHistory, order)
	// Update Favorites
	//customer.UpdateFavorites()
}

//func (customer *Customer) UpdateFavorites() {
//	var newFavorites []*Food
//
//	OrderCounts := customer.GetOrderCounts()
//	for foodID, foodCount := range OrderCounts {
//		if foodCount > 5 {
//			// Get food from database
//			//food, err := handler.FoodStore.GetByID(foodID)
//			if err != nil {
//				fmt.Println(err)
//				continue
//			}
//			// Add food to newFavorite
//			newFavorites = append(newFavorites, food)
//		}
//	}
//
//	OrderRates := customer.GetOrderRates()
//	for foodID, foodRate := range OrderRates {
//		if foodRate > 3 {
//			// Get food from database
//			//food, err := handler.FoodStore.GetByID(foodID)
//			if err != nil {
//				fmt.Println(err)
//				continue
//			}
//			// Add food to newFavorite
//			newFavorites = append(newFavorites, food)
//		}
//	}
//
//	customer.Favorites = newFavorites
//}

func (customer *Customer) GetOrderRates() map[primitive.ObjectID]float64 {
	orderRates := make(map[primitive.ObjectID]float64)
	commentCounts := make(map[primitive.ObjectID]int)
	for _, comment := range customer.Comments {
		key := comment.Food.ID
		if count, ok := commentCounts[key]; ok {
			newRate := float64(orderRates[key]*float64(count)+float64(comment.Rating)) / float64(count+1)
			orderRates[key] = newRate
			commentCounts[key] = count + 1
		} else {
			commentCounts[key] = 1
			orderRates[key] = float64(comment.Rating)
		}
	}
	return orderRates
}

func (customer *Customer) GetOrderCounts() map[primitive.ObjectID]int {
	orderCounts := make(map[primitive.ObjectID]int)
	for _, order := range customer.OrderHistory {
		for _, item := range order.Items {
			key := item.Food.ID
			if oldCount, ok := orderCounts[key]; ok {
				orderCounts[key] = oldCount + item.Count
			} else {
				orderCounts[key] = item.Count
			}
		}
	}
	return orderCounts
}

func (customer *Customer) CheckPassword(password string) bool {
	return customer.Password == password
}

func (customer *Customer) ReplaceComment(comment *Comment) {
	var newComments []*Comment
	for _, c := range customer.Comments {
		if c.ID != comment.ID {
			newComments = append(newComments, c)
		} else {
			newComments = append(newComments, comment)
		}
	}
	customer.Comments = newComments
}

func (customer *Customer) GetOrder(id primitive.ObjectID) *Order {
	for _, o := range customer.OrderHistory {
		if o.ID == id {
			return o
		}
	}
	return nil
}

func (customer *Customer) GetOrdersHistory() []*Order {
	return customer.OrderHistory
}

func (customer *Customer) ReplaceOrder(order *Order) {
	var newPendingOrders []*Order
	for _, o := range customer.OrderHistory {
		if o.ID != order.ID {
			newPendingOrders = append(newPendingOrders, o)
		} else {
			newPendingOrders = append(newPendingOrders, order)
		}
	}
	customer.OrderHistory = newPendingOrders
}
