package main

import (
	"app/internals/config"
	"app/internals/driver"
	"app/internals/handlers"
	"app/internals/loggers"
	"app/internals/models"
	"app/internals/repository/dbrepo"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	portNumber    = 3000
	templatesPath = "templates/"
)

var app config.AppConfig

func main() {
	conn, err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.SQL.Close()

	db := dbrepo.NewPostgresRepo(conn.SQL, &app)
	handlers.SetUpHandlersConfig(&app, &db)

	host := fmt.Sprintf(":%v", portNumber)

	routes := routes(&app)

	fmt.Printf("Server is up on localhost:%v\n", portNumber)
	err = http.ListenAndServe(host, routes)

	if err != nil {
		fmt.Println("Error:", err)
	}

}

func run() (*driver.DB, error) {
	l := loggers.New()
	app = config.New(templatesPath, l.InfoLogger, l.ErrorLogger)
	app.UseCache = false
	app.InProduction = false

	gob.Register(models.Reservation{})
	app.Session.Lifetime = time.Hour * 24
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	fmt.Println("Connecting to the database")
	conn, err := driver.ConnectSql("host=localhost port=6000 dbname=bookings user=bookings password=bookings")

	if err != nil {
		log.Fatal("can not connect to database", err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	return conn, nil
}
