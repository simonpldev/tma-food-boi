package db

import (
	"context"
	"log"
	"tma-food-boi/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func Connect() {
	env := config.GetENV()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(env.DatabaseConnection))
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(env.DatabaseName)
}

func UserCol() *mongo.Collection {
	return db.Collection("users")
}
