package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/akkaraju-satvik/learn-go/auth"
)

func main() {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("xmx"))
	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Satvik",
		})
	})
	router.SetTrustedProxies([]string{"localhost"})
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is healthy",
			"success": true,
		})
	})
	auth.Routes(apiV1)
	router.Run(":8080")
}
