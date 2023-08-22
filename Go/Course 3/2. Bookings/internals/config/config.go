package config

import (
	"app/pkg/cache"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	Session        *scs.SessionManager
	UseCache       bool
	TemplatesCache cache.Cache[*template.Template]
	InfoLogger     *log.Logger
	ErrorLogger    *log.Logger
	InProduction   bool
	TemplatesPath  string
}

func New(templatesPath string, infoLogger, errorLogger *log.Logger) AppConfig {
	config := AppConfig{
		Session:        scs.New(),
		UseCache:       true,
		InProduction:   true,
		TemplatesCache: *cache.NewTemplatesCache(),
		TemplatesPath:  templatesPath,
		InfoLogger:     infoLogger,
		ErrorLogger:    infoLogger,
	}

	return config
}
