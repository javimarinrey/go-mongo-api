package main

import (
	"go-mongo-api/internal/repositories"
	"log"
	"os"

	"go-mongo-api/internal/db"
	"go-mongo-api/internal/routes"
)

func main() {
	db.Connect()
	repositories.InitRepositories()

	r := routes.SetupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor Gin corriendo en puerto", port)
	r.Run(":" + port)
}
