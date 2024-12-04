package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheusfrteixeira/crud-go/src/configuration/database/mongodb"
	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/controller/routes"
)

func main() {
	logger.Info("Starting Application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	userController := InitDependencies(database)

	router := gin.Default()

	routes.IniRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
