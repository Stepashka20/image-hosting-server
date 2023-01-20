package main

import (
	"imagecloud/db"
	"imagecloud/server"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.Load()
	db.Init()
	server.Init()
}
