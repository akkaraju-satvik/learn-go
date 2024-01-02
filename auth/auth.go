package auth

import "github.com/gin-gonic/gin"

func Routes(route *gin.RouterGroup) {
	authRouter := route.Group("/auth")
	authRouter.Use(AuthMiddleware)
	authRouter.POST("/login", loginController)
	authRouter.POST("/register", registerController)
}
