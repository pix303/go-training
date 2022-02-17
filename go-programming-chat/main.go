package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/pix303/go-training/go-chat/chat"
	"github.com/pix303/go-training/go-chat/middelware"
	"github.com/pix303/go-training/go-chat/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

var trk trace.Tracer

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
		t.templ = *template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	trk.Trace(data)
	t.templ.Execute(rw, data)
}

func main() {

	trk = trace.New(os.Stdout)

	var hostParam = flag.String("host", ":8080", "The address of application server")
	flag.Parse()

	godotenv.Load()

	gomniauth.SetSecurityKey(os.Getenv("SECURE_KEY"))
	gomniauth.WithProviders(
		github.New(os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_CLIENT_SECRET"), os.Getenv("GITHUB_CALLBACK")),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_CALLBACK")),
	)

	// example of inline html response
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`<html>
		<head>
		Welcome in chat
		</head>
		<body>
		<h2>Hi hello! chat with me</h2>
		<a href="chat">Let's go to chat</a>
		<p>bla bla bla...</p>
		</body>
		</html>`))
	})

	// handle the chat page request
	http.Handle("/chat", middelware.MustAuthenticated(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", middelware.LoginHandler)

	// instance a new room and use as handler for websocket request
	r := chat.NewRoom()
	http.Handle("/room", r)
	go r.Run()

	// start up server
	trk.Trace(fmt.Sprintf("Chat server is running at %s", *hostParam))
	if err := http.ListenAndServe(*hostParam, nil); err != nil {
		log.Fatal("Error on server", err)
	}
}
