package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/orders", controller.GetOrders())
	r.GET("/orders/:order_id", controller.GetOrder())
	r.POST("/orders", controller.CreateOrder())
	r.PATCH("/orders/:order_id", controller.UpdateOrder())

}
