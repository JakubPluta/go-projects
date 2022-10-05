package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(r *gin.Engine) {
	r.GET("/tables", controller.GetTables())
	r.GET("/tables/:table_id", controller.GetTable())
	r.POST("/tables", controller.CreateTable())
	r.PATCH("/tables/:table_id", controller.UpdateTable())

}
