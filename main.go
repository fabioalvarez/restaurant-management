package main

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/routes"
	"restaurant-management/utils"
)

func main() {

	// Initiate variables struct and get environment variables
	enVar := utils.GetVariables()

	// Creates an instance of *gin.Engine. It's typically used
	// when you want to create a fresh router with default settings.
	router := gin.New()

	// Setup Routes
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	// Run Router
	err := router.Run(":" + enVar.Port)
	if err != nil {
		return
	}

}
