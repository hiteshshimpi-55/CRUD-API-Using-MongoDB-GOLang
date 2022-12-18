package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Rating   int                `json:"rating" bson:"rating"`
	Length   int                `json:"length" bson:"length"`
	Watched  bool               `json:"watched" bson:"watched"`
	Director *Director          `json:"director" bson:"director"`
}

type Director struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}
