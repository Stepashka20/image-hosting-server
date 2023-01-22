package models

import (
	"context"

	"fmt"
	"imagecloud/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Image struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
	Path string `json:"path"`
}

type GroupImage struct {
	MainKey string   `json:"key"`
	Keys    []string `json:"keys"`
}

func (i Image) NewImage(key string, hash string, path string) Image {
	db := db.GetDB()
	collection := db.Collection("images")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.InsertOne(ctx, Image{key, hash, path})
	if err != nil {
		log.Fatal(err)
	}
	return Image{key, fmt.Sprintf("%x", hash), path}
}

func (i Image) NewGroupImages(key string, keys []string) {
	db := db.GetDB()
	collection := db.Collection("image_groups")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.InsertOne(ctx, GroupImage{key, keys})
	if err != nil {
		log.Fatal(err)
	}
}

func (i Image) GetImageByHash(hash string) Image {
	db := db.GetDB()
	collection := db.Collection("images")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var image Image
	res := collection.FindOne(ctx, bson.M{"hash": hash})
	if res.Err() != nil {
		return image
	}
	res.Decode(&image)
	return image
}
