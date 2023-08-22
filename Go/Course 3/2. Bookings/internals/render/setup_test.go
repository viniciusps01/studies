package render

import (
	"app/internals/config"
	"app/internals/loggers"
	"app/internals/models"
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"
)

var app *config.AppConfig

func TestMain(m *testing.M) {
	l := loggers.New()
	config := config.New("/templates", l.InfoLogger, l.ErrorLogger)
	app = &config
	app.UseCache = false
	app.InProduction = false

	gob.Register(models.Reservation{})
	app.Session.Lifetime = time.Hour * 24
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	os.Exit(m.Run())
}

type myWriter struct{}

func (myWriter) Header() http.Header {
	return http.Header{}
}
func (myWriter) Write([]byte) (int, error) {
	return 200, nil
}
func (myWriter) WriteHeader(statusCode int) {

}
