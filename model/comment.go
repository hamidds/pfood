package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentReply struct {
	CommentID primitive.ObjectID `json:"_id"     bson:"_id"             `
	Answer    string             `json:"answer"    bson:"answer"            `
}

type Comment struct {
	ID     primitive.ObjectID `json:"_id"     bson:"_id"             `
	Food   *Food              `json:"food"    bson:"food"            `
	Text   string             `json:"text"    bson:"text"            `
	Answer string             `json:"answer"  bson:"answer"          `
	Rating int                `json:"rating"  bson:"rating"          `
}

func NewComment() *Comment {
	var comment Comment
	comment.ID = primitive.NewObjectID()
	comment.Text = ""
	comment.Answer = ""
	comment.Food = &Food{}
	comment.Rating = 1
	return &comment
}

func (comment *Comment) SetFields(c *Comment) {
	comment.ID = c.ID
	comment.Text = c.Text
	comment.Answer = c.Answer
	comment.Food = c.Food
	comment.Rating = c.Rating
}

func (comment *Comment) SetAnswer(answer string) {
	comment.Answer = answer
}
