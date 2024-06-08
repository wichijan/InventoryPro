package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client struct for websocket connection and message sending
type Client struct {
	ID      string
	RoomId  string
	IsAdmin bool
	UserId  string
	Conn    *websocket.Conn
	send    chan string
	hub     *Hub
}

// NewClient creates a new client
func NewClient(roomId string, isAdmin bool, userId string, conn *websocket.Conn, hub *Hub) *Client {
	log.Printf("User %s is admin: %v; Conn: %v", userId, isAdmin, fmt.Sprintf("%p", conn))

	return &Client{
		ID:      fmt.Sprintf("%p", conn),
		RoomId:  roomId,
		UserId:  userId,
		IsAdmin: isAdmin,
		Conn:    conn,
		send:    make(chan string, 256),
		hub:     hub,
	}
}

// Client goroutine to read messages from client
func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		c.hub.broadcast <- msg
	}
}

// Client goroutine to write messages to client
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} else {
				err := c.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println("Error: ", err)
					break
				}
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Client closing channel to unregister client
func (c *Client) Close() {
	close(c.send)
}

// Function to handle websocket connection and register client to hub and start goroutines
func ServeWS(ctx *gin.Context, roomId string, userId string, isAdmin bool, hub *Hub) {
	fmt.Print("ServeWS: ", roomId)
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client := NewClient(roomId, isAdmin, userId, ws, hub)

	hub.register <- client
	go client.Write()
	go client.Read()
}
