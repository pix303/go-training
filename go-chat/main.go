package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/pix303/go-training/go-chat/trace"
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
	t.templ.Execute(rw, r)
}

func main() {

	trk := trace.New(os.Stdout)

	var hostParam = flag.String("host", ":8080", "The address of application server")
	flag.Parse()

	// example of inline html response
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`<html>
		<head>
		Welcome in chat
		</head>s
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
	trk.Trace("Chat server is running")
	if err := http.ListenAndServe(*hostParam, nil); err != nil {
		log.Fatal("Error on server", err)
	}
}
