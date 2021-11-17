package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    template.Template
}

func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = *template.Must(template.ParseFiles(filepath.Join("chat/template", t.filename)))
	})
	t.templ.Execute(rw, nil)

}

func main() {
	log.Println("Start up a chat server")

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

	http.Handle("/chat", &templateHandler{filename: "chat.html"})

	r := NewRoom()
	http.Handle("/room", r)
	go r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error on server", err)
	}
}
