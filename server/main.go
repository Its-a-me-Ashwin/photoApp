package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	database "github.com/Its-a-me-Ashwin/photoApp/server/database"
)

const (
	mongoUri string "127.0.0.1:27017"
)

func main () {
	fmt.Print("Starting")

	client, ctx, _, err := database.ConnectToMongoDb(mongoUri)
	if err != nil {
		fmt.Prinln("Unable to connect to DB", err)
	}
	err = database.Ping(ctx, client)
	if err != nil {
		fmt.Prinln("Failed to ping DB", err)
	}
	fmt.Prinln("Pinged DB", err)
}