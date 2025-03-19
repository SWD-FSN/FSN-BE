package cmd

import (
	"log"
	"os"
	api_route "social_network/api_route"
	"social_network/constant/env"
	"social_network/util"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func Execute() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file in main - " + err.Error())
	// }

	var server = gin.Default()
	var logger = util.GetLogConfig()
	var apiPort string = os.Getenv(env.API_PORT)

	config(logger)
	setUpApiRoutes(server, logger, apiPort)

	if err := server.Run(":" + apiPort); err != nil {
		logger.Println("Error run server " + err.Error())
	}
}

func setUpApiRoutes(server *gin.Engine, logger *log.Logger, port string) {
	// User
	api_route.InitializeUserAPIRoute(server, logger, port)
	// Post
	api_route.InitializePostAPIRoute(server, logger, port)
}
