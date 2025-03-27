package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializePersonalProfileAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "personal-profile"
	server.GET(contextPath+"/user/:userId/actor/:actorId", controller.GetPersonalProfile)
}
