package routes

import (
	"go-mongo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/users", h.GetUsers)
		api.POST("/users", h.CreateUser)
		api.DELETE("/users/:id", h.DeleteUser)
	}

	return r
}
