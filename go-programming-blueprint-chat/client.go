package main

import (
	"github.com/gorilla/websocket"
)

// client is a struct that rappresents user/browser using chat page
type client struct {
	// reference to room socket
	socket *websocket.Conn
	// channel to send message to socket
	send chan []byte
	// belonging room
	room *room
}

// read waits messages from socket and when they arrive broadcasting in room
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// write wait socket send text (submit button next textarea in chat page)
func (c *client) write() {
	for msg := range c.send {
		if len(msg) > 0 {
			if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}
	c.socket.Close()
}
