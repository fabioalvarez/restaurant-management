package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders", controllers.GetOrders())
	router.GET("/orders/:orderId", controllers.GetOrder())
	router.POST("orders/", controllers.CreateOrder())
	router.PATCH("orders/:orderId", controllers.UpdateOrder())
}
