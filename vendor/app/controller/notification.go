package controller

import (
	"log"
	"strconv"

	"app/shared/websocket"
)

// SendNotificationByID send new notification to specific user
func SendNotificationByID(id uint32, content interface{}) error {
	clients := websocket.MainPool.Clients
	for client := range clients {
		ci, _ := strconv.Atoi(client.ID)
		if ci == int(id) {
			err := client.Conn.WriteJSON(content)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

// SendNotificationToAll send new notification to all online users
func SendNotificationToAll(content interface{}) error {
	var err error
	clients := websocket.MainPool.Clients
	for client := range clients {
		err = client.Conn.WriteJSON(content)
		if err != nil {
			log.Println(err)
		}
	}
	return err
}
