package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf the CSRF protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return sessionMgr.LoadAndSave(next)
}
