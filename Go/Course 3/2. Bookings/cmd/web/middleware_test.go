package main

import (
	"testing"
)

func TestCSRFCheck(t *testing.T) {
	next := myHandler{}

	h := CSRFcheck(next)

	if h == nil {
		t.Errorf("failed CSRFCheck(): expect http.Handler, but got %T", h)
	}
}

func TestHandleSession(t *testing.T) {
	mh := myHandler{}

	h := handleSession(mh)

	if h == nil {
		t.Errorf("failed handleSession(): expect http.Handler, but got %T", h)
	}
}
