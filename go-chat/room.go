package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (rm *room) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP error: ", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   rm,
	}

	rm.join <- client
	defer func() {
		rm.leave <- client
	}()
	go client.write()
	client.read()
}

func (r *room) run() {
	for {
		select {
		case joiningClient := <-r.join:
			r.clients[joiningClient] = true
		case leavingClient := <-r.leave:
			delete(r.clients, leavingClient)
			close(leavingClient.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send message
				default:
					//failed to send because of client is closed
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
