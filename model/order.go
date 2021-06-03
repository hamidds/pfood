package model

import "time"

type Item struct {
	Food  *Food `json:"food"        bson:"food"`
	Count int   `json:"count"       bson:"count"`
}

const Pending = "Pending!"
const Preparing = "Preparing!"
const Delivering = "Delivering!"
const Delivered = "Delivered!"

type Order struct {
	//Foods      Food   `json:"foods"       bson:"foods"`
	//Foods      map[string]int `json:"foods"       bson:"foods"`
	Items               []*Item       `json:"items"                 bson:"items"`
	TotalPrice          float64       `json:"total_price"           bson:"total_price"`
	Available           bool          `json:"available"             bson:"available"`
	Customer            *Customer     `json:"customer"              bson:"customer" `
	State               string        `json:"state"                 bson:"state"`
	PreparationDuration time.Duration `json:"preparation_duration"  bson:"preparation_duration"`
	DeliveryDuration    time.Duration `json:"delivery_duration"     bson:"delivery_duration"`
	ConfirmTime         time.Time     `json:"confirm_time"          bson:"confirm_time"`
}

func (order *Order) GetState() string {
	t := order.ConfirmTime.Add(order.PreparationDuration)
	if time.Now().Before(t) {
		order.SetState(Preparing)
	} else {
		t = t.Add(order.DeliveryDuration)
		if time.Now().Before(t) {
			order.SetState(Delivering)
		} else {
			order.SetState(Delivered)
		}
	}
	return order.State
}

func (order *Order) SetState(state string) {
	order.State = state
}
func (order *Order) SetConfirmTime(time time.Time) {
	order.ConfirmTime = time
}
func (order *Order) SetPreparationDuration(duration time.Duration) {
	order.PreparationDuration = duration
}
func (order *Order) SetDeliveryDuration(duration time.Duration) {
	order.DeliveryDuration = duration
}

func (order *Order) AddItem(food *Food, count int) {
	order.Items = append(order.Items, &Item{Food: food, Count: count})
}

func (order *Order) CalculatePrice() {
	var balance float64
	for _, item := range order.Items {
		balance += item.Food.Price * float64(item.Count)
	}
	order.TotalPrice = balance
}

//func (order *Order) Remove(food *Food, count int) {
//	order.Items = append(order.Items, Item{Food: food, Count: count})
//	// Update Food Rating
//	//food.UpdateRating(comment)
//	// Add Comment
//	// Update DB ?
//}
//
//func (order *Order) EditItem(food *Food, count int) {
//	order.Items = append(order.Items, Item{Food: food, Count: count})
//	// Update Food Rating
//	//food.UpdateRating(comment)
//	// Add Comment
//	// Update DB ?
//}
