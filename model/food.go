package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"                     `
	Name       string             `json:"name"         bson:"name"                              `
	Price      float64            `json:"price"        bson:"price"                             `
	Available  bool               `json:"available"    bson:"available"                         `
	Comments   []*Comment         `json:"comments"     bson:"comments"                          `
	Rating     float64            `json:"rating"       bson:"rating"                            `
	Restaurant *Restaurant        `json:"restaurant"   bson:"restaurant"    validate:"isdefault"`
}

func NewFood() *Food {
	var food Food
	food.ID = primitive.NewObjectID()
	food.Name = ""
	food.Price = 0.0
	food.Available = true
	food.Comments = []*Comment{}
	food.Rating = 1
	food.Restaurant = &Restaurant{}
	return &food
}

func (food *Food) SetFields(f *Food) {
	food.Name = f.Name
	food.Price = f.Price
	food.Available = f.Available
	food.Comments = f.Comments
	food.Rating = f.Rating
	food.Restaurant = f.Restaurant
}

func (food *Food) AddComment(comment *Comment) {
	// Update Food Rating
	food.UpdateRating(comment)
	// Add Comment
	food.Comments = append(food.Comments, comment)
	// Update DB ?
}

func (food *Food) UpdateRating(comment *Comment) {
	newRating := float64(food.Rating*float64(len(food.Comments))+float64(comment.Rating)) / float64(len(food.Comments))
	food.Rating = newRating
}

func (food *Food) SetAvailable(status bool) {
	food.Available = status
}
