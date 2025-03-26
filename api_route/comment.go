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
	server.DELETE(contextPath+"/delete/actor/:actor_id/comment/:commentId", controller.RemoveComment)
}

// {
//     "username": "NamLord",
//     "full_name": "Nam Nguyen",
//     "email": "externalauthdemo1234@gmail.com",
//     "password": "@Aa12345678",
//     "date_of_birth": "1990-05-15T00:00:00Z",
//     "profile_avatar": "avatar_url.jpg"
// }{
//     "username": "NamLord",
//     "full_name": "Nam Nguyen",
//     "email": "externalauthdemo1234@gmail.com",
//     "password": "@Aa12345678",
//     "date_of_birth": "1990-05-15T00:00:00Z",
//     "profile_avatar": "avatar_url.jpg"
// }
