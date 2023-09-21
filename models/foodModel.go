package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      *string            `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Price     *float64           `bson:"price" json:"price" validate:"required"`
	FoodImage *string            `bson:"foodImage" json:"foodImage" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	FoodId    string             `bson:"foodId" json:"foodId"`
	MenuId    *string            `bson:"menuId" json:"menuId" validate:"required"`
}
