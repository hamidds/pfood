package model

type Manager struct {
	Email      string      `json:"email"       bson:"email"        validate:"required,email"`
	Password   string      `json:"password"    bson:"password"     validate:"required"`
	Restaurant *Restaurant `json:"restaurant"  bson:"restaurant"   validate:"isdefault"`
	Name       *Restaurant `json:"name"        bson:"name"         validate:"isdefault"`
}
