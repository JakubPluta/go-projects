package main

import (
	"os"
	// 	"github.com/gin-gonic/gin"
	// "github.com/JakubPluta/go-projects/go-restaurant-management/routes"
	// "github.com/JakubPluta/go-projects/go-restaurant-management/database"
	// "github.com/JakubPluta/go-projects/go-restaurant-management/middleware"
	// "go.mongodb.org/mongo-driver/mongo"

	"github.com/JakubPluta/go-projects/go-restaurant-management/database"
	"github.com/JakubPluta/go-projects/go-restaurant-management/middleware"
	"github.com/JakubPluta/go-projects/go-restaurant-management/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(
	database.Client, "food",
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Authentication())

	routes.UserRoutes(r)
	routes.FoodRoutes(r)
	routes.MenuRoutes(r)
	routes.TableRoutes(r)
	routes.OrderRoutes(r)
	routes.OrderItemRoutes(r)
	routes.InvoiceRoutes(r)

	r.Run(port)
}
