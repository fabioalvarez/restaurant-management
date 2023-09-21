package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Query the menu collection
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listening the menu items"})
			return
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get menuId from parameter
		menuId := c.Param("menuId")
		var menu models.Menu

		// Query the collection using the menuId
		err := menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return menu when succeed
		c.JSON(http.StatusOK, menu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Initiate menu variable as menu model type
		var menu models.Menu

		// Receive the data a transform it to a Menu struct
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the shared information
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Update fields
		menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.MenuId = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			log.Println("Error:", insertErr)
			msg := fmt.Sprintf("Menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context that is canceled when the specified timeout elapses
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Initiate menu variable as menu model type
		var menu models.Menu

		// Receive the data a transform it to a Menu struct
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get parameters and create filter
		menuId := c.Param("menuId")
		filter := bson.M{"menuId": menuId}

		var updateObj primitive.D

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}

		menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updatedAt", Value: menu.UpdatedAt})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := "Menu update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
