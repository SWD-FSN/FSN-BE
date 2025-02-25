package apiroute

import (
	"log"
	"os"
	"social_network/constant/env"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

const (
	backUpPort string = "8080"
)

func InitializeUserAPIRoute(server *gin.Engine, logger *log.Logger) {
	var port string = os.Getenv(env.API_PORT)

	if port == "" {
		port = backUpPort
	}

	var contextPath string = "users"

	server.GET(contextPath, controller.GetAllUsers)
}
