package model

type Comment struct {
	Food   *Food   `json:"food"    bson:"food"`
	Text   string `json:"text"    bson:"text"`
	Answer string `json:"answer"  bson:"answer"`
	Rating int    `json:"rating"  bson:"rating"`
}
