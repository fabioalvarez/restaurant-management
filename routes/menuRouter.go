package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menus", controllers.GetMenus())
	router.GET("/menus/:menuId", controllers.GetMenu())
	router.POST("/menus", controllers.CreateMenu())
	router.PATCH("/menus/:menuId", controllers.UpdateMenu())
}
