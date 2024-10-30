package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IsraelTeo/api-jwt-go/db"
	"github.com/IsraelTeo/api-jwt-go/route"
	"github.com/joho/godotenv"
)

func main() {

	r := route.InitRoute()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loanding .env main")
	}

	err = db.Connection()
	if err != nil {
		log.Fatalf("Error trying to connect with database: %v", err)
	}
	err = db.MigrateDB()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("Database migration successful")

	fmt.Println("Starting server on port 8080...")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
