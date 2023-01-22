package models

import (
	"context"

	"imagecloud/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGroup(key string) ([]Image, error) {
	db := db.GetDB()
	group_collection := db.Collection("image_groups")
	images_collection := db.Collection("images")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var group GroupImage
	res := group_collection.FindOne(ctx, bson.M{"mainkey": key})
	if res.Err() != nil {
		return []Image{}, res.Err()
	}
	res.Decode(&group)
	var images []Image
	var imgHash = map[string]Image{}
	cur, err := images_collection.Find(ctx, bson.M{"key": bson.M{"$in": group.Keys}})
	if err != nil {
		log.Fatal(err)
		return []Image{}, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var image Image
		cur.Decode(&image)
		imgHash[image.Key] = image
	}
	for _, key := range group.Keys {
		images = append(images, imgHash[key])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return []Image{}, err
	}
	return images, nil

}
