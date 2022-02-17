package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

// client is a struct that rappresents user/browser using chat page
type client struct {
	// reference to room socket
	socket *websocket.Conn

	// channel to send message to socket
	send chan Message

	// belonging room
	room *room

	//user data
	userData map[string]interface{}
}

// read waits messages from socket and when they arrive broadcasting in room
func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		c.room.forward <- *msg
	}
}

// write wait socket send text (submit button next textarea in chat page)
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
