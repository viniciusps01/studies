package main

import (
	"app/pkg/config"
	"app/pkg/handlers"
	"fmt"
	"net/http"
	"time"
)

const portNumber = 3000

var app config.AppConfig

func main() {
	app = config.New()
	app.UseCache = false
	app.InProduction = false
	handlers.SetUpHandlersConfig(&app)

	app.Session.Lifetime = time.Hour * 24
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	host := fmt.Sprintf(":%v", portNumber)

	routes := routes(&app)

	fmt.Printf("Server is up on localhost:%v\n", portNumber)
	err := http.ListenAndServe(host, routes)

	if err != nil {
		fmt.Println("Error:", err)
	}

}
