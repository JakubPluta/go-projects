package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(r *gin.Engine) {
	r.GET("/orderItems", controller.GetOrderItems())
	r.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	r.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	r.POST("/orderItems", controller.CreateOrderItem())
	r.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
