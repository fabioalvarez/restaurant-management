package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"
)

var orderCollection = database.OpenCollection(database.Client, "order")

func OrderItemOrderCreator(order models.Order, ctx context.Context) (string, error) {

	order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.OrderId = order.ID.Hex()

	if _, err := orderCollection.InsertOne(ctx, order); err != nil {
		return "", err
	}

	return order.OrderId, nil
}
