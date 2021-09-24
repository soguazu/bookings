package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/soguazu/bookings/pkg/config"
	"github.com/soguazu/bookings/pkg/handlers"
	"github.com/soguazu/bookings/pkg/render"
)

const portNumber = ":4000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplatesCache()

	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)

	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting application on port %s\n", portNumber)

	log.Fatal(srv.ListenAndServe())
}
