package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func CSRFcheck(next http.Handler) http.Handler {
	csrf := nosurf.New(next)

	cookie := http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: !app.InProduction,
		SameSite: http.SameSiteLaxMode,
	}
	csrf.SetBaseCookie(cookie)

	return csrf

}

func handleSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
