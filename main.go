package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/xmriz/go-scraper/cmd"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Execute the CLI command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
