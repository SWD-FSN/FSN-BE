package service

import (
	"encoding/json"
	"log"
	"net/http"
	"social_network/dto"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
}

var clients = make(map[string]*websocket.Conn) // Map userID to WebSocket connection

func RegisterUserConnection(req dto.UserConnectionRequest) {
	cnn, err := upgrader.Upgrade(req.Writer, req.Request, nil)

	if err != nil {
		log.Println("WebSocket upgrade failed: " + err.Error())
		return
	}

	defer cnn.Close()

	var mu sync.Mutex

	// Store user
	mu.Lock()
	clients[req.UserId] = cnn
	mu.Unlock()

	for {
		if _, _, err := cnn.ReadMessage(); err != nil {
			mu.Lock()
			delete(clients, req.UserId) // Remove user on disconnect
			mu.Unlock()
			break
		}
	}
}

func sendMessage(req dto.WSSendMessageRequest, logger *log.Logger, cnn *websocket.Conn) {
	msgJson, err := json.Marshal(req)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	if err := cnn.WriteMessage(websocket.TextMessage, msgJson); err != nil {
		logger.Println(err.Error())
	}
}

func generateContentAndContentTypeOfMsg(actorUsername, actionType, objectType, orgContent string) (string, string) {
	var content, contentType string

	// Set content type
	if actionType == "like" || actionType == "comment" || actionType == follow_request || actionType == add_friend_request {
		contentType = "notification"
	} else if actionType == "message" {
		contentType = "message"
	}

	// Set content
	if orgContent != "" {
		content = orgContent
	} else {
		switch actionType {
		case follow_request:
			content = actorUsername + " sends a " + actionType + " request to you."
		case add_friend_request:
			content = actorUsername + " sends an " + actionType + " request to you."
		default:
			content = actorUsername + " " + actionType + "s " + "on your " + objectType + "."
			// VD: Nam Nguyen + like + s + on your + post
		}
	}

	return content, contentType
}
