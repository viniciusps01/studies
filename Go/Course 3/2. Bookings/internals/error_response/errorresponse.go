package error_response

import (
	"app/internals/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

func ClientError(app *config.AppConfig, w http.ResponseWriter, status int) {
	app.ErrorLogger.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(app *config.AppConfig, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())

	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)

	app.ErrorLogger.Println(trace)
}
