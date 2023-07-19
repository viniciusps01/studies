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
	InProduction   bool
}

func New() AppConfig {
	config := AppConfig{
		Session:        scs.New(),
		UseCache:       true,
		InProduction:   true,
		TemplatesCache: *cache.NewTemplatesCache(),
	}

	return config
}
