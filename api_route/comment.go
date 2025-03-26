package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeCommentAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "comments"

	// var authGroup = server.Group(contextPath, middlewares.Authorize)
	// authGroup.GET("/post/:id", controller.GetCommentsFromPost)
	server.GET(contextPath+"/post", controller.GetCommentsFromPost)
	server.POST(contextPath+"/create", controller.PostComment)
	server.PUT(contextPath+"/edit", controller.EditComment)
}
