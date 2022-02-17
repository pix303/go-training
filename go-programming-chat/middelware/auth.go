package middelware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pix303/go-training/go-chat/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type authHandler struct {
	next http.Handler
}

var trk trace.Tracer = trace.New(os.Stdout)

func (ah *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authValue, err := r.Cookie("auth")

	trk.Trace("Checks cookies for authorization")

	// manage if not authenticated
	if err == http.ErrNoCookie {
		trk.Trace(fmt.Sprintf("No cookie present for authorization: %s", err.Error()))
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	// manage generic error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	trk.Trace(fmt.Sprintf("Successful authorization for: %s", authValue.Value))

	// move on request
	ah.next.ServeHTTP(w, r)
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	// split url parts login|callback/:action/:provider
	segments := strings.Split(r.URL.Path, "/")
	action := segments[2]
	provider := segments[3]

	switch action {
	case "login":
		// get provider Github (active) | Google | Facebook
		authProvider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error on get provider for %s: %s", provider, err.Error()), http.StatusBadRequest)
		}
		// get url for redirect to provider for login
		loginURL, err := authProvider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error on get url from provider: %s", err.Error()), http.StatusInternalServerError)
		}
		rw.Header().Set("Location", loginURL)
		rw.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		// get provider Github (active) | Google | Facebook
		authProvider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error on get provider for %s: %s", provider, err.Error()), http.StatusBadRequest)
		}
		creds, err := authProvider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error on complete auth for %s: %s", provider, err.Error()), http.StatusInternalServerError)
		}
		user, err := authProvider.GetUser(creds)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error on get user from creds %s: %s", provider, err.Error()), http.StatusInternalServerError)
		}

		candidateName := "Mona"
		if user != nil {
			candidateName = user.Name()
		}
		authCookieValue := objx.New(map[string]interface{}{
			"name": candidateName,
		}).MustBase64()

		http.SetCookie(rw, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})

		rw.Header().Set("Location", "/chat")
		rw.WriteHeader(http.StatusTemporaryRedirect)

	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "Auth for action %s and provider %s not supported", action, provider)
	}
}

// MustAuthenticated used when endopoint require authed user
func MustAuthenticated(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
