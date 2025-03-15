package util

import (
	"encoding/json"
	"log"
	"social_network/dto"

	"github.com/gorilla/websocket"
)

func HandleUserConnection(req dto.UserConnectionRequest) {
	cnn, err := req.Upgrader.Upgrade(req.Writer, req.Request, nil)

	if err != nil {
		log.Println("WebSocket upgrade failed: ", err.Error())
		return
	}

	defer cnn.Close()

	req.Clients[req.Id] = cnn

	for {
		_, _, err := cnn.ReadMessage()

		if err != nil {
			delete(req.Clients, req.Id)
			break
		}
	}
}

func SendMessage(req dto.WSSendMessageRequest, cnn *websocket.Conn) error {
	msgJson, err := json.Marshal(req)
	if err != nil {

	}

	if err := cnn.WriteMessage(websocket.TextMessage, msgJson); err != nil {

	}

	return nil
}
