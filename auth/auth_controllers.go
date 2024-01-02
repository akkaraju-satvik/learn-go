package auth

import "github.com/gin-gonic/gin"

func loginController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
		"success": true,
	})
}

func registerController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register",
		"success": true,
	})
}
