package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	r.GET("/foods", controller.GetFoods())
	r.GET("/foods/:food_id", controller.GetFood())
	r.POST("/foods", controller.CreateFood())
	r.PATCH("/foods/:food_id", controller.UpdateFood())

}
