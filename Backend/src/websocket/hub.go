package websocket

import (
	"log"
	"time"

	"github.com/wichijan/InventoryPro/src/utils"
)

// Hub is a struct that holds all the clients and the messages that are sent to them
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Unregistered clients.
	unregister chan *Client
	// Register requests from the clients.
	register chan *Client
	// Inbound messages from the clients.
	broadcast chan Message
}

// Message struct to hold message data
type Message struct {
	Type         string `json:"type"`
	SentToUserId string `json:"sentToUserId"`
	Sender       string `json:"sender"`
	Content      string `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		unregister: make(chan *Client),
		register:   make(chan *Client),
		broadcast:  make(chan Message),
	}
}

// Core function to run the hub
func (h *Hub) Run() {
	for {
		select {
		// Register a client.
		case client := <-h.register:
			h.RegisterNewClient(client)
		// Unregister a client.
		case client := <-h.unregister:
			h.RemoveClient(client)
		// Broadcast a message to all clients.
		case message := <-h.broadcast:
			h.HandleMessage(message)
		}
	}
}

// Function to check if room exists and if not create it and add client to it
func (h *Hub) RegisterNewClient(client *Client) {
	connections := h.clients
	if connections == nil {
		connections = make(map[*Client]bool)
		h.clients = connections
	}
	h.clients[client] = true
	log.Printf("Client registered: %s, Admin: %v", client.UserId, client.IsAdmin)
}

// Function to remove client from room
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
		if len(h.clients) == 0 {
			delete(h.clients, client)
		}
	}
	close(client.send)
	log.Printf("Client unregistered: %s, Admin: %v", client.UserId, client.IsAdmin)
}

// Function to handle message based on type of message
func (h *Hub) HandleMessage(message Message) {
	messageStr := `{ "Message": "` + message.Content + `", "TimeStamp": "` + time.Now().Format(time.RFC3339) + `" }`
	log.Print("Message: ", messageStr)

	if len(h.clients) == 0 {
		log.Print("No clients to send message to")
		return
	}

	if message.Type == utils.MESSAGE_TYPE_TO_USER {
		log.Printf("Send to user %v", message.SentToUserId)
		for client := range h.clients {
			if client.UserId == message.SentToUserId {
				select {
				case client.send <- `{"SentToUserId": "`+ message.SentToUserId + `", "Sender": "`+ message.Sender + `", "Message": "` + message.Content + `", "TimeStamp": "` + time.Now().Format(time.RFC3339) + `" }`:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}

	}

	if message.Type == utils.MESSAGE_TYPE_TO_ADMINS {
		log.Print("Send to admins")
		for client := range h.clients {
			log.Print("Client ID: ", client.ID)
			if client.IsAdmin {
				select {
				case client.send <- messageStr:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}

	if message.Type == utils.MESSAGE_TYPE_EVERYONE {
		log.Print("Send to everyone")
		for client := range h.clients {
			select {
			case client.send <- messageStr:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}

}
