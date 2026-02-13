package repositories

import (
	"context"
	"fmt"
	"go-mongo-api/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Client) *ProductRepository {
	return &ProductRepository{
		// Usamos la colección "products"
		collection: db.Database("testdb").Collection("products"),
	}
}

func (r *ProductRepository) Create(p *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	// Extraemos el ID generado y lo asignamos al modelo
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid
	}

	return nil
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var products []models.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &products)
	return products, err
}

func (r *ProductRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (r *ProductRepository) Update(id string, p models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err // ID mal formado
	}

	// Definimos los campos a actualizar
	update := bson.M{
		"$set": bson.M{
			"name":  p.Name,
			"price": p.Price,
			"stock": p.Stock,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments // El ID no existe en la DB
	}

	return nil
}

func (r *ProductRepository) GetSize() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{"$group", bson.D{
				{"_id", nil},
				{"avgSize", bson.D{{"$avg", bson.D{{"$bsonSize", "$$ROOT"}}}}},
				{"maxSize", bson.D{{"$max", bson.D{{"$bsonSize", "$$ROOT"}}}}},
				{"minSize", bson.D{{"$min", bson.D{{"$bsonSize", "$$ROOT"}}}}},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println("Promedio (bytes):", result["avgSize"])
		fmt.Println("Máximo (bytes):", result["maxSize"])
		fmt.Println("Mínimo (bytes):", result["minSize"])
	}
}
