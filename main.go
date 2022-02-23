package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"test.com/mutant-go/application/db"
	"test.com/mutant-go/application/routes"
)

func main() {
	loadEnv()
	db.Initialize()
	server := gin.Default()

	routes.Routes(server)
	server.Run(os.Getenv("PORT"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
