package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/internals/config"
	"app/internals/models"
	"app/internals/render"
)

func setDefaultData(td *models.TemplateData) {

	td.Data = map[string]interface{}{
		"usernames": []string{
			"notsobrad",
			"notsocrood",
		},
	}
}

var appConfig *config.AppConfig

func SetUpHandlersConfig(c *config.AppConfig) {
	appConfig = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	setDefaultData(&td)
	render.RenderTemplate(w, r, "index.html", appConfig, &td)
}

func About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	setDefaultData(&td)
	render.RenderTemplate(w, r, "about.html", appConfig, &td)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	setDefaultData(&td)

	render.RenderTemplate(w, r, "contact.html", appConfig, &td)
}

func SearchAvailability(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	setDefaultData(&td)

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

	setDefaultData(&td)

	render.RenderTemplate(w, r, "majors-suite.html", appConfig, &td)
}

func GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	setDefaultData(&td)

	render.RenderTemplate(w, r, "generals-quarters.html", appConfig, &td)
}

func MakeReservation(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}

	setDefaultData(&td)

	render.RenderTemplate(w, r, "make-reservation.html", appConfig, &td)
}
