package main

import (
	"app/internals/config"
	"app/internals/handlers"
	"app/internals/models"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"
)

const portNumber = 3000

var app config.AppConfig

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	host := fmt.Sprintf(":%v", portNumber)

	routes := routes(&app)

	fmt.Printf("Server is up on localhost:%v\n", portNumber)
	err = http.ListenAndServe(host, routes)

	if err != nil {
		fmt.Println("Error:", err)
	}

}

func run() error {
	app = config.New()
	app.UseCache = false
	app.InProduction = false
	handlers.SetUpHandlersConfig(&app)

	gob.Register(models.Reservation{})
	app.Session.Lifetime = time.Hour * 24
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	return nil
}
