package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/akkaraju-satvik/learn-go/auth"
)

func main() {
	godotenv.Load(".env")
	listenAddr := flag.String("listenAddr", "", "HTTP listen address")
	flag.Parse()
	if *listenAddr == "" {
		log.Fatal("Listen address is required")
	}
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.LoadHTMLGlob("static/*.html")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/assets", "./static/assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Satvik",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
			"title": "404",
		})
	})
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is healthy",
			"success": true,
		})
	})
	auth.Routes(apiV1)
	router.Run(":" + *listenAddr)
}
