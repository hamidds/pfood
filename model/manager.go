package model

type Manager struct {
	Email      string      `json:"email"       bson:"email"        validate:"required,email"`
	Password   string      `json:"password"    bson:"password"     validate:"required"`
	Name       string `json:"name"        bson:"name"         validate:"isdefault"`
	Restaurant *Restaurant `json:"restaurant"  bson:"restaurant"   validate:"isdefault"`
}

func (manager Manager) CheckPassword(password string) bool {
	return manager.Password == password
}

func (manager Manager) SetRestaurant(restaurant *Restaurant) {
	manager.Restaurant = restaurant
}
