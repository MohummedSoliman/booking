// Package handlers used for handle reqeusts paths.
package handlers

import (
	"net/http"

	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/MohummedSoliman/booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "about.page.html")
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "generals.page.html")
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "majors.page.html")
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "contact.page.html")
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "search-availability.page.html")
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted from page"))
	// render.RenderTemplate(w, "search-availability.page.html")
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{}, "make-reservation.page.html")
}
