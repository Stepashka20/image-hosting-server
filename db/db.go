package db

// init mongoDB

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func Init() {
	Client, _ = mongo.NewClient(options.Client().ApplyURI())
	Client.Connect(nil)
	DB = Client.Database("imagecloud")
}
