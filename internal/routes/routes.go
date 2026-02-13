package routes

import (
	"go-mongo-api/internal/handlers"
	"go-mongo-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uH *handlers.UserHandler, pH *handlers.ProductHandler) *gin.Engine {
	//r := gin.Default()
	r := gin.New() // Usamos New() para no traer el Logger por defecto y usar el nuestro

	// Aplicamos nuestro middleware personalizado
	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery()) // Para que la API no se caiga ante un pánico

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
		api.PUT("/products/:id", pH.UpdateProduct)
		api.GET("/products/size", pH.GetSize)
	}

	return r
}
