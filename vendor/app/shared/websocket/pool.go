package websocket

import "fmt"

var MainPool *Pool

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPoolStart() {
	MainPool = &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
	go MainPool.Start()
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			// fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// for client, _ := range pool.Clients {
			// 	fmt.Println("client----------", client.ID)
			// 	// client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			// }
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			// fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// for client, _ := range pool.Clients {
			// 	client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			// }
			break
		case message := <-pool.Broadcast:
			fmt.Println("Treat Incoming message", message)
		}
	}
}
