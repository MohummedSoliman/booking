package render

import (
	"net/http"
	"testing"

	"github.com/MohummedSoliman/booking/pkg/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("Flash Value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathTemplate = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var mw myWriter

	err = RenderTemplate(&mw, r, &models.TemplateData{}, "home.page.html")
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&mw, r, &models.TemplateData{}, "non-exist.page.html")
	if err == nil {
		t.Error("Render template that doesn not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathTemplate = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
