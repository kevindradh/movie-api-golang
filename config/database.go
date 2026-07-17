package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var MovieCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(Env.MongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	// Check connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping failed:", err)
	}

	log.Println("Connected to MongoDB!")

	DB = client
	MovieCollection = DB.Database(Env.MongoDBName).Collection("movies")
	UserCollection = DB.Database(Env.MongoDBName).Collection("users")
}
