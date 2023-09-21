package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Category  string             `bson:"category" json:"category" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	MenuId    string             `bson:"menuId" json:"menuId"`
}
