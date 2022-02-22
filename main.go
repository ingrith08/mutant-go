package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"test.com/mutant-go/application/db"
	"test.com/mutant-go/application/routes"
)

func main() {
	loadEnv()
	db.Initialize()
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("ALLOW_URL"), ";")
	config.AllowCredentials = true
	config.AllowHeaders = []string{"content-type", "captcha"}
	server.Use(cors.New(config))

	routes.Routes(server)
	server.Run(os.Getenv("PORT"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
