package render

import (
	"app/internals/config"
	"app/internals/models"
	"app/pkg/cache"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, filename string, config *config.AppConfig, templateData *models.TemplateData) {
	filepath := "../../templates/" + filename
	templateData.CSRFToken = nosurf.Token(r)
	var template *template.Template

	if config.UseCache {
		template = readTemplateFromCache(filepath, &config.TemplatesCache)
	}

	if template != nil {
		renderTemplate(w, template, templateData)
		return
	}

	log.Println("Reading template from disk")
	template, err := readTemplateFromDisk(filepath)

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

func readTemplateFromDisk(filepath string) (*template.Template, error) {
	template, err := template.ParseFiles(filepath)

	if err != nil {
		return template, err
	}

	hasTemplates, err := hasTemplateFiles()

	if err != nil {
		return template, err
	}

	if !hasTemplates {
		return template, nil
	}

	template, err = template.ParseGlob("../../templates/*-template.html")

	if err != nil {
		return template, err
	}

	return template, nil
}

func hasTemplateFiles() (bool, error) {
	matches, err := filepath.Glob("../../templates/*-template.html")

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
