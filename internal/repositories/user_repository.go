package repositories

import (
	"context"
	"go-mongo-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		collection: db.Database("testdb").Collection("users"),
	}
}

func (r *UserRepository) Create(u models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), u)
	return err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &users)
	return users, err
}

func (r *UserRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
