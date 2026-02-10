package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name" binding:"required"`
	Price float64            `bson:"price" json:"price" binding:"required"`
	Stock int                `bson:"stock" json:"stock"`
}
