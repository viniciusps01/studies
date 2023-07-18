package handlers

import (
	"net/http"

	"app/pkg/config"
	"app/pkg/models"
	"app/pkg/render"
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
	ip := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "remote_ip", ip)
	td := models.TemplateData{}
	setDefaultData(&td)
	render.RenderTemplate(w, "index.html", appConfig, &td)
}

func About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	ip := appConfig.Session.GetString(r.Context(), "remote_ip")

	setDefaultData(&td)

	td.Data["remote_ip"] = ip
	render.RenderTemplate(w, "about.html", appConfig, &td)
}

func HelloFelix(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	setDefaultData(&td)

	age := r.URL.Query().Get("age")
	td.Data["felix_age"] = age

	render.RenderTemplate(w, "hello-felix.html", appConfig, &td)
}
