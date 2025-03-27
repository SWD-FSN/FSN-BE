package cmd

import (
	"log"
	"os"
	api_route "social_network/api_route"
	"social_network/constant/env"
	"social_network/util"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func Execute() {
	// Load configuration
	var logger = util.GetLogConfig()
	config(logger)

	var server = gin.Default()
	var apiPort string = os.Getenv(env.API_PORT)

	// Enable CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, or specify ["http://example.com"]
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	// Comment
	api_route.InitializeCommentAPIRoute(server, logger, port)
	// Conversation
	api_route.InitializeConversationAPIRoute(server, logger, port)
	// Like
	api_route.InitializeLikeAPIRoute(server, logger, port)
	// Social Request
	api_route.InitializeSocialRequestAPIRoute(server, logger, port)
	// Notification
	api_route.InitializeNotificationRoute(server, logger, port)
	// Search object
	api_route.InitializeSearchObjectAPIRoute(server, logger, port)
}
