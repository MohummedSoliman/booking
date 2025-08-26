// Package handlers used for handle reqeusts paths.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/forms"
	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/MohummedSoliman/booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
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
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	log.Println(start, end)

	w.Write([]byte("Posted from page"))
	// render.RenderTemplate(w, "search-availability.page.html")
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, &models.TemplateData{
		Form: forms.New(nil),
	}, "make-reservation.page.html")
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Has("first_name", r)

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.RenderTemplate(w, r, &models.TemplateData{
			Form: form,
			Data: data,
		}, "make-reservation.page.html")
		return
	}
}
