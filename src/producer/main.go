package main

import (
	"log"
	"producer/api"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system variables
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app := api.GetApiApp()
	app.Listen(":3000")
}
