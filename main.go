package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakharmaurya/gin_test/users"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")

	users.AddRoutes(api)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
