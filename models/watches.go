package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Watch struct{
	ID primitive.ObjectID `json:"id,omitempty" bson: "_id,omitempty"`
	Name string `json:"name,omitempty" bson: "name,omitempty"`
	Brand string `json:"brand,omitempty" bson: "brand,omitempty"`
	Price int `json:"price,omitempty" bson: "price,omitempty"`
	Image string `json:"image,omitempty" bson: "image,omitempty"`
	Description string `json:"description,omitempty" bson: "description,omitempty"`
}