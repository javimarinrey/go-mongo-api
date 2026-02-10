package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// URI con Replica Set
	uri := "mongodb://mongo1:27017,mongo2:27018,mongo3:27019/?replicaSet=rs0"

	// Configuración del pool de conexiones
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(50).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(30 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo hacer ping a MongoDB:", err)
	}

	log.Println("Conectado a MongoDB Replica Set con éxito")
	return client
}
