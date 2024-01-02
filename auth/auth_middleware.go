package auth

import (
	"encoding/json"
	"fmt"

	"github.com/akkaraju-satvik/learn-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthRequestBody struct {
	EmailId  string `json:"email_id" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func AuthMiddleware(c *gin.Context) {
	body := c.Request.Body
	x := json.NewDecoder(body)
	var data AuthRequestBody
	x.Decode(&data)
	err := utils.Validate.Struct(data)
	if err != nil {
		errMsg := err.(validator.ValidationErrors)
		var errors []string
		for _, v := range errMsg {
			if v.Tag() == "required" {
				errors = append(errors, fmt.Sprintf("%s is required", v.Field()))
			}
			if v.Tag() == "email" {
				errors = append(errors, fmt.Sprintf("%s is not valid", v.Field()))
			}
		}
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"success": false,
			"errors":  errors,
		})
		c.Abort()
		return
	}
	c.Set("email_id", data.EmailId)
	c.Set("password", data.Password)
	c.Next()
}
