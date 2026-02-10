package main

import (
	"context"
	"log"
	"os"
	"time"

	"go-mongo-api/internal/handlers"
	"go-mongo-api/internal/repositories"
	"go-mongo-api/internal/routes"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configuración de MongoDB con Réplicas
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://mongo1:27017,mongo2:27017,mongo3:27017/?replicaSet=rs0"
	}

	clientOptions := options.Client().ApplyURI(uri).
		SetConnectTimeout(20 * time.Second).
		SetServerSelectionTimeout(15 * time.Second).
		SetMaxPoolSize(50).
		SetMinPoolSize(10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping verifica la conexión real
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Error: No se pudo conectar a MongoDB: ", err)
	}

	log.Println("✅ Conexión exitosa al Replica Set de MongoDB")

	nodes := client.NumberSessionsInProgress()
	log.Printf("Sesiones activas: %d", nodes)

	// Repositorios
	userRepo := repositories.NewUserRepository(client)
	productRepo := repositories.NewProductRepository(client)

	// Handlers
	userHandler := handlers.NewUserHandler(userRepo)
	productHandler := handlers.NewProductHandler(productRepo)

	// Router (le pasamos ambos)
	r := routes.SetupRouter(userHandler, productHandler)
	log.Println("Servidor Gin iniciado en puerto 8080")
	r.Run(":8080")

}
