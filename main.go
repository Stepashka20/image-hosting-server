package main

import (
	// "flag"
	// "fmt"
	// "os"

	"imagecloud/db"
	"imagecloud/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// environment := flag.String("e", "development", "")
	// flag.Usage = func() {
	// 	fmt.Println("Usage: server -e {mode}")
	// 	os.Exit(1)
	// }
	// flag.Parse()
	// config.Init(*environment)
	db.Init()
	server.Init()
}
