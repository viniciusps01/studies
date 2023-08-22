package render

import (
	"app/internals/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	w := myWriter{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	filename := "index.html"
	td := models.TemplateData{}

	RenderTemplate(w, r, filename, app, &td)
}
