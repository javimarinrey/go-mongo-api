package routes

import (
	"go-mongo-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.Default())

	// Middleware logging + recovery ya viene con gin.Default()

	// Usuarios
	users := r.Group("/users")
	{
		users.GET("", handlers.GetUsersHandler)
		users.POST("", handlers.CreateUserHandler)
	}

	// Productos
	products := r.Group("/products")
	{
		products.GET("", handlers.GetProductsHandler)
		products.POST("", handlers.CreateProductHandler)
	}

	return r
}
