package controller

import (
	"log"

	"app/shared/websocket"
)

// SendNotificationByID send new notification to specific user
func SendNotificationByID(id string, content interface{}) error {
	clients := websocket.MainPool.Clients
	for client := range clients {
		if client.ID == id {
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
