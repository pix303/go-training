package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templateHandler is a struct for handle chat request
type templateHandler struct {
	// for loading only once the template first time it is requested
	once sync.Once
	// template filename
	filename string
	// template reference
	templ template.Template
}

// ServeHTTP implement Handler interface for templateHandler type
func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = *template.Must(template.ParseFiles(filepath.Join("chat/template", t.filename)))
	})
	t.templ.Execute(rw, nil)

}

func main() {
	log.Println("Start up a chat server")

	// example of inline html response
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`<html>
		<head>
		Welcome in chat
		</head>
		<body>
		<h2>Hi hello! chat with me</h2>
		<p>bla bla bla...</p>
		</body>
		</html>`))
	})

	// handle the chat page request
	http.Handle("/chat", &templateHandler{filename: "chat.html"})

	// instance a new room and use as handler for websocket request
	r := NewRoom()
	http.Handle("/room", r)
	go r.run()

	// start up server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error on server", err)
	}
}
