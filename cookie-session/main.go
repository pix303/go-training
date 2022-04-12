package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

const sessionKey = "test-cookie-session"
const authKey = "authenticated"

func init() {
	store = sessions.NewCookieStore([]byte("test"))
	store.Options.HttpOnly = true
	store.Options.Path = "/"
	store.Options.MaxAge = 30
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("GET")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(authMiddleware)
	apiRouter.HandleFunc("/auth", authHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func homeHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("testing cookie session"))
}

func authHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("testing cookie session - authed pass"))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		as, err := store.Get(rq, sessionKey)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		if as.Values[authKey] != nil {
			ok := as.Values[authKey].(bool)
			if !ok {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("no auth"))
				return
			}
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("no auth"))
			return
		}

		next.ServeHTTP(rw, rq)
	})
}

func loginHandler(rw http.ResponseWriter, rq *http.Request) {
	err := rq.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(err.Error()))
		return
	}

	formValues := rq.Form
	u := formValues.Get("u")
	p := formValues.Get("p")
	if u != "test" || p != "test" {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("user and pass wrong"))
		return
	}

	as, err := store.Get(rq, sessionKey)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(err.Error()))
		return
	}

	as.Values[authKey] = true
	err = as.Save(rq, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte("ok!"))
}

func logoutHandler(rw http.ResponseWriter, rq *http.Request) {
	as, err := store.Get(rq, sessionKey)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	as.Options.MaxAge = -1
	as.Values[authKey] = false
	as.Save(rq, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte("ok removed from sessions"))

}
