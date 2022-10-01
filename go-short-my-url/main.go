package main

import (
	"fmt"

	"github.com/JakubPluta/go-projects/go-short-my-url/handlers"
	"github.com/JakubPluta/go-projects/go-short-my-url/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitStore()

	err := r.Run(":9000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
