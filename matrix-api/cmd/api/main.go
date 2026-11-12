package main

import (
	"log"
	"os"

	api "interseguro/matrix-api/internal/adapters/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := api.New().Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
