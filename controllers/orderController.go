package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"
)

var orderCollection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Query all orders to db
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing order items"})
		}

		// Create variable to add results
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatalln(err)
		}

		// Return result if not error was gotten
		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get parameter
		orderId := c.Param("orderId")

		// Initiate order struct
		var order models.Order

		// Query order by id
		err := orderCollection.FindOne(ctx, bson.M{"orderId": orderId}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the orders"})
		}

		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Initialize variables
		var order models.Order
		var table models.Table

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Bind body
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create order in table
		if order.TableId != nil {
			err := tableCollection.FindOne(ctx, bson.M{"tableId": order.TableId}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.OrderId = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Initiate variables
		var order models.Order
		var table models.Table
		var updateObj primitive.D

		// Get parameter value
		orderId := c.Param("orderId")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.TableId != nil {
			err := orderCollection.FindOne(ctx, bson.M{"tableId": order.TableId}).Decode(&table)
			if err != nil {
				msg := fmt.Sprintf("table has no orders")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "tableId", Value: order.TableId})
		}

		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updatedAt", Value: order.UpdatedAt})

		upsert := true

		filter := bson.M{"oderId": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$st", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}

}
