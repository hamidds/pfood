package model

import (
	"github.com/hamidds/pfood/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CustomerResponse struct {
	Customer struct {
		PhoneNumber  string     `json:"phone_number"   bson:"_id"`
		Name         string     `json:"name"           bson:"name"`
		Credit       int        `json:"credit"         bson:"credit"`
		District     int        `json:"district"       bson:"district"`
		Address      string     `json:"address"        bson:"address"`
		OrderHistory []*Order   `json:"order_history"  bson:"order_history"`
		Favorites    []*Food    `json:"favorites"      bson:"favorites"`
		Comments     []*Comment `json:"comments"       bson:"comments"`
		Token        string     `json:"token"`
	} `json:"user"`
}

func NewUserResponse(customer *Customer) *CustomerResponse {
	response := new(CustomerResponse)
	response.Customer.PhoneNumber = customer.PhoneNumber
	response.Customer.Name = customer.Name
	response.Customer.Credit = customer.Credit
	response.Customer.District = customer.District
	response.Customer.Address = customer.Address
	response.Customer.OrderHistory = customer.OrderHistory
	response.Customer.Favorites = customer.Favorites
	response.Customer.Comments = customer.Comments
	response.Customer.Token = utils.GenerateJWT("phone_number", customer.PhoneNumber)
	return response
}

type FoodResponse struct {
	Food struct {
		Name       string      `json:"name"         bson:"name"                              `
		Price      float64     `json:"price"        bson:"price"                             `
		Available  bool        `json:"available"    bson:"available"                         `
		Comments   []*Comment  `json:"comments"     bson:"comments"                          `
		Rating     float64     `json:"rating"       bson:"rating"                            `
		Restaurant *Restaurant `json:"restaurant"   bson:"restaurant"    validate:"isdefault"`
	} `json:"food"`
}

func NewFoodResponse(food *Food) *FoodResponse {
	response := new(FoodResponse)
	response.Food.Name = food.Name
	response.Food.Price = food.Price
	response.Food.Available = food.Available
	response.Food.Comments = food.Comments
	response.Food.Rating = food.Rating
	response.Food.Restaurant = food.Restaurant
	return response
}

type FoodsResponse struct {
	Foods *[]Food `json:"foods" bson:"foods"`
}

func NewFoodsResponse(foods *[]Food) *FoodsResponse {
	return &FoodsResponse{Foods: foods}
}

type ManagerResponse struct {
	Manager struct {
		Email      string      `json:"email"       bson:"email"        validate:"required,email"`
		Restaurant *Restaurant `json:"restaurant"  bson:"restaurant"   validate:"isdefault"`
		Name       string      `json:"name"        bson:"name"         validate:"isdefault"`
	} `json:"manager"`
}

func NewManagerResponse(manager *Manager) *ManagerResponse {
	response := new(ManagerResponse)
	response.Manager.Email = manager.Email
	response.Manager.Name = manager.Name
	response.Manager.Restaurant = manager.Restaurant
	return response
}

type RestaurantResponse struct {
	Restaurant struct {
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
	} `json:"restaurant"`
}

func NewRestaurantResponse(restaurant *Restaurant) *RestaurantResponse {
	response := new(RestaurantResponse)
	response.Restaurant.Name = restaurant.Name
	response.Restaurant.Address = restaurant.Address
	response.Restaurant.District = restaurant.District
	response.Restaurant.Districts = restaurant.Districts
	response.Restaurant.DeliveryPrice = restaurant.DeliveryPrice
	response.Restaurant.DeliveryTime = restaurant.DeliveryTime
	response.Restaurant.OpenTime = restaurant.OpenTime
	response.Restaurant.CloseTime = restaurant.CloseTime
	response.Restaurant.Foods = restaurant.Foods
	response.Restaurant.PendingOrders = restaurant.PendingOrders
	return response
}

type CommentResponse struct {
	Comment struct {
		ID     primitive.ObjectID `json:"_id"     bson:"_id"             `
		Food   *Food              `json:"food"    bson:"food"            `
		Text   string             `json:"text"    bson:"text"            `
		Answer string             `json:"answer"  bson:"answer"          `
		Rating int                `json:"rating"  bson:"rating"          `
	} `json:"comment"`
}

func NewCommentResponse(comment *Comment) *CommentResponse {
	response := new(CommentResponse)
	response.Comment.ID = comment.ID
	response.Comment.Food = comment.Food
	response.Comment.Text = comment.Text
	response.Comment.Answer = comment.Answer
	response.Comment.Rating = comment.Rating
	return response
}

type CommentsResponse struct {
	Comments []*Comment `json:"comments" bson:"comments"`
}

func NewCommentsResponse(comments []*Comment) *CommentsResponse {
	return &CommentsResponse{Comments: comments}
}

type OrderResponse struct {
	Order struct {
		ID                  primitive.ObjectID `json:"_id" bson:"_id"`
		Items               []*Item            `json:"items"                 bson:"items"`
		TotalPrice          float64            `json:"total_price"           bson:"total_price"`
		Available           bool               `json:"available"             bson:"available"`
		Customer            *Customer          `json:"customer"              bson:"customer" `
		State               string             `json:"state"                 bson:"state"`
		PreparationDuration time.Duration      `json:"preparation_duration"  bson:"preparation_duration"`
		DeliveryDuration    time.Duration      `json:"delivery_duration"     bson:"delivery_duration"`
		ConfirmTime         time.Time          `json:"confirm_time"          bson:"confirm_time"`
	} `json:"order"`
}

func NewOrderResponse(order *Order) *OrderResponse {
	response := new(OrderResponse)
	response.Order.ID = order.ID
	response.Order.Items = order.Items
	response.Order.TotalPrice = order.TotalPrice
	response.Order.Customer = order.Customer
	response.Order.State = order.GetState()
	response.Order.PreparationDuration = order.PreparationDuration
	response.Order.DeliveryDuration = order.DeliveryDuration
	response.Order.ConfirmTime = order.ConfirmTime
	return response
}

type OrdersResponse struct {
	Comments []*Order `json:"orders" bson:"orders"`
}

func NewOrdersResponse(orders []*Order) *OrdersResponse {
	return &OrdersResponse{Comments: orders}
}
