package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler

	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
	// passed case
	default:
		t.Errorf("Type Correct Type is %T, but is %T", "http.Handler", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler

	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:
	// passed case
	default:
		t.Errorf("Type Correct Type is %T, but is %T", "http.Handler", v)
	}
}
