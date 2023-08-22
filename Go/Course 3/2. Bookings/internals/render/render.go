package render

import (
	"app/internals/config"
	"app/internals/models"
	"app/pkg/cache"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, filename string, config *config.AppConfig, templateData *models.TemplateData) {
	filepath := config.TemplatesPath + filename
	templateData.CSRFToken = nosurf.Token(r)
	templateData.Flash = config.Session.PopString(r.Context(), "flash")
	templateData.Warning = config.Session.PopString(r.Context(), "warning")
	templateData.Error = config.Session.PopString(r.Context(), "error")
	var template *template.Template

	if config.UseCache {
		template = readTemplateFromCache(filepath, &config.TemplatesCache)
	}

	if template != nil {
		renderTemplate(w, template, templateData)
		return
	}

	template, err := readTemplateFromDisk(config.TemplatesPath, filepath)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if config.UseCache {
		saveTemplateToCache(template, filepath, &config.TemplatesCache)
	}

	renderTemplate(w, template, templateData)
}

func renderTemplate(w http.ResponseWriter, template *template.Template, templateData *models.TemplateData) {
	err := template.Execute(w, templateData)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func readTemplateFromDisk(templatesPath string, filepath string) (*template.Template, error) {
	template, err := template.ParseFiles(filepath)

	if err != nil {
		return template, err
	}

	hasTemplates, err := hasTemplateFiles(templatesPath)

	if err != nil {
		return template, err
	}

	if !hasTemplates {
		return template, nil
	}

	template, err = template.ParseGlob(templatesPath + "*-template.html")

	if err != nil {
		return template, err
	}

	return template, nil
}

func hasTemplateFiles(templatesPath string) (bool, error) {
	matches, err := filepath.Glob(templatesPath + "*-template.html")

	if err != nil {
		return false, err
	}

	if len(matches) == 0 {
		return false, nil
	}

	return true, nil
}

func readTemplateFromCache(filepath string, cache *cache.Cache[*template.Template]) *template.Template {
	value := (*cache).Get(filepath)
	return value
}

func saveTemplateToCache(template *template.Template, filepath string, cache *cache.Cache[*template.Template]) {
	(*cache).Add(filepath, template)
}
