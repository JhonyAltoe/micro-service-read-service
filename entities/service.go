package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Services []Service `bson:"services,omitempty" json:"result,omitempty"`
}

type Service struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string	`bson:"name,omitempty" json:"name,omitempty"`
	MinDuration int `bson:"min_duration,omitempty" json:"min_duration,omitempty"`
	MinPrice float32 `bson:"min_price,omitempty" json:"min_price,omitempty"`
}
