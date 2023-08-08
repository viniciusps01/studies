package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/internals/config"
	"app/internals/forms"
	"app/internals/models"
	"app/internals/render"
)

var appConfig *config.AppConfig

func SetUpHandlersConfig(c *config.AppConfig) {
	appConfig = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	render.RenderTemplate(w, r, "index.html", appConfig, &td)
}

func About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	render.RenderTemplate(w, r, "about.html", appConfig, &td)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "contact.html", appConfig, &td)
}

func SearchAvailability(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "search-availability.html", appConfig, &td)
}

func SearchAvailabilityPost(w http.ResponseWriter, r *http.Request) {
	data := fmt.Sprintf(
		"Starting Date: %v \nEnding Date: %v",
		r.Form.Get("start"),
		r.Form.Get("end"),
	)

	bytes := []byte(data)

	w.Write(bytes)
}

type availabilityJsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	data := availabilityJsonResponse{
		Ok:      true,
		Message: "Success",
	}

	json, err := json.Marshal(data)

	if err != nil {
		w.Write([]byte(fmt.Sprint("Error: ", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func MajorsSuite(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "majors-suite.html", appConfig, &td)
}

func GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "generals-quarters.html", appConfig, &td)
}

func MakeReservation(w http.ResponseWriter, r *http.Request) {
	emptyReservation := models.Reservation{}

	td := models.TemplateData{
		Form: forms.New(nil),
		Data: map[string]interface{}{
			"reservation": emptyReservation,
		},
	}

	render.RenderTemplate(w, r, "make-reservation.html", appConfig, &td)
}

func MakeReservationPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")

	form.MinLength("first_name", 3)
	form.MinLength("last_name", 3)
	form.MinLength("email", 7)
	form.ValidateEmail("email")
	form.MinLength("phone", 10)

	td := models.TemplateData{
		Form: form,
		Data: map[string]interface{}{
			"reservation": reservation,
		},
	}

	if !form.Valid() {
		render.RenderTemplate(w, r, "make-reservation.html", appConfig, &td)
		return
	}

	appConfig.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation := appConfig.Session.Get(r.Context(), "reservation")

	if reservation == nil {
		appConfig.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	td := models.TemplateData{
		Data: map[string]interface{}{
			"reservation": reservation,
		},
	}

	render.RenderTemplate(w, r, "reservation-summary.html", appConfig, &td)

}