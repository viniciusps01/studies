package cache

import (
	"html/template"
)

type Cache[T any] interface {
	Get(id string) T
	Add(id string, data T)
}

type TemplatesCache struct {
	data map[string]*template.Template
}

func (t TemplatesCache) Get(id string) *template.Template {
	value := t.data[id]
	return value
}

func (t TemplatesCache) Add(id string, template *template.Template) {
	t.data[id] = template
}

func NewTemplatesCache() *TemplatesCache {
	cache := &TemplatesCache{
		data: map[string]*template.Template{},
	}

	return cache
}
