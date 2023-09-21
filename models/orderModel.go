package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	OrderId   string             `bson:"orderId" json:"orderId"`
	TableId   *string            `bson:"tableId" json:"tableId" validate:"required"`
}
