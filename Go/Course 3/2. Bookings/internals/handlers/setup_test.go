package handlers

import (
	"app/internals/config"
	"app/internals/driver"
	"app/internals/loggers"
	"app/internals/models"
	"app/internals/repository/dbrepo"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	l := loggers.New()
	app = config.New("../../templates/", l.InfoLogger, l.ErrorLogger)
	app.UseCache = false
	app.InProduction = false

	gob.Register(models.Reservation{})
	app.Session.Lifetime = time.Hour * 24
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	conn, err := driver.ConnectSql("host=localhost port=6000 dbname=bookings user=bookings password=bookings")

	if err != nil {
		log.Fatal("can not connect to database", err)
	}

	db := dbrepo.NewPostgresRepo(conn.SQL, &app)

	SetUpHandlersConfig(&app, &db)

	router := chi.NewRouter()

	router.Use(handleSession)
	// router.Use(middleware.Logger)
	// router.Use(CSRFcheck)
	router.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("../static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", http.HandlerFunc(Home))

	router.Get("/about", http.HandlerFunc(About))

	router.Get("/contact", http.HandlerFunc(Contact))

	router.Get("/search-availability", http.HandlerFunc(SearchAvailability))
	router.Post("/search-availability", http.HandlerFunc(SearchAvailabilityPost))
	router.Post("/search-availability-json", http.HandlerFunc(AvailabilityJson))

	router.Get("/make-reservation", http.HandlerFunc(MakeReservation))
	router.Post("/make-reservation", http.HandlerFunc(MakeReservationPost))
	router.Get("/reservation-summary", http.HandlerFunc(ReservationSummary))

	router.Get("/majors-suite", http.HandlerFunc(MajorsSuite))

	router.Get("/generals-quarters", http.HandlerFunc(GeneralsQuarters))

	return router
}

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
