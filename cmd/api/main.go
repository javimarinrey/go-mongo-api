package main

import (
	"context"
	"log"
	"time"

	"go-mongo-api/internal/handlers"
	"go-mongo-api/internal/repositories"
	"go-mongo-api/internal/routes"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configuración de MongoDB con Réplicas
	uri := "mongodb://mongo1:27017,mongo2:27018,mongo3:27019/?replicaSet=rs0"

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(50).
		SetMinPoolSize(10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Inyección de dependencias
	userRepo := repositories.NewUserRepository(client)
	userHandler := handlers.NewUserHandler(userRepo)

	// Iniciar Servidor
	r := routes.SetupRouter(userHandler)
	log.Println("Servidor Gin iniciado en puerto 8080")
	r.Run(":8080")
}
