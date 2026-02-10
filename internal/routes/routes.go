package routes

import (
	"go-mongo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uH *handlers.UserHandler, pH *handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Rutas de Usuarios
		api.GET("/users", uH.GetUsers)
		api.POST("/users", uH.CreateUser)
		api.DELETE("/users/:id", uH.DeleteUser)

		// Rutas de Productos
		api.GET("/products", pH.GetProducts)
		api.POST("/products", pH.CreateProduct)
		api.DELETE("/products/:id", pH.DeleteProduct)
	}

	return r
}
