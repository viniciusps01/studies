package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/internals/config"
	"app/internals/error_response"
	"app/internals/forms"
	"app/internals/models"
	"app/internals/render"
	"app/internals/repository"
	"app/internals/repository/dbrepo"
)

var repo repository.Repository

func SetUpHandlersConfig(c *config.AppConfig, dbrepo *dbrepo.DatabaseRepo) {
	repo.App = c
	repo.DB = dbrepo
}

func Home(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	render.RenderTemplate(w, r, "index.html", repo.App, &td)
}

func About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	render.RenderTemplate(w, r, "about.html", repo.App, &td)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "contact.html", repo.App, &td)
}

func SearchAvailability(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "search-availability.html", repo.App, &td)
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
		error_response.ServerError(repo.App, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func MajorsSuite(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "majors-suite.html", repo.App, &td)
}

func GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	render.RenderTemplate(w, r, "generals-quarters.html", repo.App, &td)
}

func MakeReservation(w http.ResponseWriter, r *http.Request) {
	emptyReservation := models.Reservation{}

	td := models.TemplateData{
		Form: forms.New(nil),
		Data: map[string]interface{}{
			"reservation": emptyReservation,
		},
	}

	render.RenderTemplate(w, r, "make-reservation.html", repo.App, &td)
}

func MakeReservationPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		error_response.ServerError(repo.App, w, err)
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
		render.RenderTemplate(w, r, "make-reservation.html", repo.App, &td)
		return
	}

	repo.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation := repo.App.Session.Get(r.Context(), "reservation")

	if reservation == nil {
		repo.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	td := models.TemplateData{
		Data: map[string]interface{}{
			"reservation": reservation,
		},
	}

	render.RenderTemplate(w, r, "reservation-summary.html", repo.App, &td)

}
