package main

import (
	"fmt"
	"log"

	"github.com/fdhliakbar/go-envy/pkg/envy"
)

func main() {
	// Load .env file (opsional)
	err := envy.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system env only")
	}

	// Get with default
	dbURL := envy.Get("DB_URL", "postgres://localhost:5432/mydb")

	// Get int
	port := envy.GetInt("PORT", 8080)

	// Get bool
	debug := envy.GetBool("DEBUG", false)

	// Require (panic kalau gak ada)
	secret := envy.Require("SECRET_KEY")

	fmt.Println("DB:", dbURL)
	fmt.Println("PORT:", port)
	fmt.Println("DEBUG:", debug)
	fmt.Println("SECRET:", secret)
}
