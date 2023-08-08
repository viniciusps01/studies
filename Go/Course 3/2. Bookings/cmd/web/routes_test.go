package main

import (
	"app/internals/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	app := config.AppConfig{}
	r := routes(&app)

	if r == nil {
		t.Errorf("failed routes(): expect htt.Handler, but got %T", r)
	}
}
