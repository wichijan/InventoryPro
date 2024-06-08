package websocket

import (
	"log"
)

// Hub is a struct that holds all the clients and the messages that are sent to them
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool
	// Admin clients.
	admins map[*Client]bool
	// Unregistered clients.
	unregister chan *Client
	// Register requests from the clients.
	register chan *Client
	// Inbound messages from the clients.
	broadcast chan Message
}

// Message struct to hold message data
type Message struct {
	Type     string `json:"type"`
	ForAdmin bool   `json:"forAdmin"`
	Sender   string `json:"sender"`
	Content  string `json:"content"`
	ID       string `json:"id"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		admins:     make(map[*Client]bool),
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
	if client.IsAdmin {
		h.admins[client] = true
	} else {
		connections := h.clients[client.RoomId]
		if connections == nil {
			connections = make(map[*Client]bool)
			h.clients[client.RoomId] = connections
		}
		h.clients[client.RoomId][client] = true
	}
	log.Printf("Client registered: %s, Admin: %v", client.UserId, client.IsAdmin)
}

// Function to remove client from room
func (h *Hub) RemoveClient(client *Client) {
	if client.IsAdmin {
		delete(h.admins, client)
	} else if _, ok := h.clients[client.RoomId]; ok {
		delete(h.clients[client.RoomId], client)
		if len(h.clients[client.RoomId]) == 0 {
			delete(h.clients, client.RoomId)
		}
	}
	close(client.send)
	log.Printf("Client unregistered: %s, Admin: %v", client.UserId, client.IsAdmin)
}

// Function to handle message based on type of message
func (h *Hub) HandleMessage(message Message) {
	h.broadcastToAdmins(message)
	if message.ForAdmin {
		return
	}
	h.broadcastToRoom(message)
}

// Function to broadcast messages to all clients in a room
func (h *Hub) broadcastToRoom(message Message) {
	clients := h.clients[message.ID]
	for client := range clients {
		select {
		case client.send <- message.Content:
		default:
			close(client.send)
			delete(clients, client)
		}
	}
}

// Function to broadcast messages to all admin clients
func (h *Hub) broadcastToAdmins(message Message) {
	for client := range h.admins {
		select {
		case client.send <- message.Content:
		default:
			close(client.send)
			delete(h.admins, client)
		}
	}
}
