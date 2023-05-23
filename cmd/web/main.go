package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ryanhoangt/go-web-bookings/pkg/config"
	"github.com/ryanhoangt/go-web-bookings/pkg/handlers"
	"github.com/ryanhoangt/go-web-bookings/pkg/render"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var sessionMgr *scs.SessionManager

func main() {
	appConfig.InProduction = false

	// set up session manager
	sessionMgr = scs.New()
	sessionMgr.Lifetime = 24 * time.Hour
	sessionMgr.Cookie.Persist = true
	sessionMgr.Cookie.SameSite = http.SameSiteLaxMode
	sessionMgr.Cookie.Secure = appConfig.InProduction

	appConfig.SessionMgr = sessionMgr

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	}
	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = appConfig.InProduction

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTemplateRenderer(&appConfig)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server.", err)
	}
}
