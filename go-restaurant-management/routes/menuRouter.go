package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(r *gin.Engine) {
	r.GET("/menus", controller.GetMenus())
	r.GET("/menus/:menu_id", controller.GetMenu())
	r.POST("/menus", controller.CreateMenu())
	r.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
