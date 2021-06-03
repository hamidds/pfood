package model

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
	customer.UpdateFavorites()
}

func (customer *Customer) AddOrder(order *Order) {
	customer.OrderHistory = append(customer.OrderHistory, order)
	// Update Favorites
	customer.UpdateFavorites()
}

func (customer *Customer) UpdateFavorites() {

	//var newFavorite []Food
	//
	//OrderCounts := customer.GetOrderCounts()
	//for foodName, foodCount := range OrderCounts {
	//	restaurant = strings.Split(foodName, "+")[0]
	//	food = strings.Split(foodName, "+")[1]
	//	if foodCount > 5 {
	//		// Get food from database
	//		// Add it to newFavorite
	//	}
	//}
	//
	//OrderRates := customer.GetOrderRates()
	//for foodName, foodRate := range OrderRates {
	//	restaurant = strings.Split(foodName, "+")[0]
	//	food = strings.Split(foodName, "+")[1]
	//	if foodRate > 3 {
	//		// Get food from database
	//		// Add it to newFavorite
	//	}
	//}
}

func (customer *Customer) GetOrderRates() map[string]float64 {
	orderRates := make(map[string]float64)
	commentCounts := make(map[string]int)
	for _, comment := range customer.Comments {
		key := comment.Food.Name + "+" + comment.Food.Restaurant.Name
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

func (customer *Customer) GetOrderCounts() map[string]int {
	orderCounts := make(map[string]int)
	for _, order := range customer.OrderHistory {
		for _, item := range order.Items {
			key := item.Food.Name + "+" + item.Food.Restaurant.Name
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
