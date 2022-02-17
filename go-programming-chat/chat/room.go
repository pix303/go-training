package chat

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/pix303/go-training/go-chat/trace"
	"github.com/stretchr/objx"
)

// room is struct that rappresent a group of chatting clients
type room struct {
	// chan to broadcast room messages
	forward chan Message
	// chan to manage client add in room
	join chan *client
	// chan to manage client remove from room
	leave chan *client
	// list to track all client in room
	clients map[*client]bool
	// track activity
	Tracker trace.Tracer
}

// define const for socket buffer
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// init and managing websocket
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// NewRoom helper function to instantiate a new room
func NewRoom() *room {
	return &room{
		forward: make(chan Message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		Tracker: trace.New(os.Stdout),
	}
}

// ServeHTTP implement handler interface in order to use as request handler and start up the websocket
func (rm *room) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// start up and add request socket protocol
	socket, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Fatal("Room ServeHTTP error: ", err)
		return
	}

	authCookie, err := r.Cookie("auth")
	if err != nil {
		log.Fatal("Cookie auth error: ", err)
		return
	}

	// create the client (new chat browser page)
	client := &client{
		// pair socket
		socket: socket,
		// create send message channel
		send: make(chan Message, messageBufferSize),
		//pair room
		room: rm,
		// set user data for cookie
		userData: objx.MustFromBase64(authCookie.Value),
	}

	// client add request to room
	rm.join <- client

	// when client closes connection after stop reading it sets to room a leave request
	defer func() {
		rm.leave <- client
	}()

	// add a new goroutine for
	go client.write()
	client.read()
}

// run wait channel message request: join, leave, send messages
func (r *room) Run() {
	for {
		select {
		case joiningClient := <-r.join:
			r.clients[joiningClient] = true
			r.Tracker.Trace("New client joint")
		case leavingClient := <-r.leave:
			delete(r.clients, leavingClient)
			close(leavingClient.send)
			r.Tracker.Trace("Client left")
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				//send message
				case client.send <- msg:
					r.Tracker.Trace(fmt.Sprintf("Message sent to client: %s", msg))
				default:
					//failed to send because of client is closed
					delete(r.clients, client)
					close(client.send)
					r.Tracker.Trace("Message fail to send to client. Client removed from room")
				}
			}
		}
	}
}
