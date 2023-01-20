package db

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func Init() {
	Client, _ = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	Client.Connect(nil)
	fmt.Println("Connected to MongoDB")
	DB = Client.Database("imagecloud")
}

func GetDB() *mongo.Database {
	return DB
}
