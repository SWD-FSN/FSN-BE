package cmd

import (
	"log"
	api_route "social_network/api_route"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func Execute() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file in main - " + err.Error())
	// }

	var server = gin.Default()
	var logger = &log.Logger{}

	api_route.InitializeUserAPIRoute(server, logger)
	api_route.InitializePostAPIRoute(server, logger)

	if err := server.Run(":8080"); err != nil {
		logger.Println("Error run server " + err.Error())
	}
}
