package users

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(api *gin.RouterGroup) {
	{
		api.GET("/users", GetAllUser)
		api.POST("/users", CreateUser)
		api.GET("/users/:id", GetUser)
		api.PUT("/users/:id", UpdateUser)
		api.DELETE("/users/:id", DeleteUser)
	}
}
