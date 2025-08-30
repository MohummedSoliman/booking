// Package handlers used for handle reqeusts paths.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MohummedSoliman/booking/driver"
	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/forms"
	"github.com/MohummedSoliman/booking/pkg/helpers"
	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/MohummedSoliman/booking/pkg/render"
	"github.com/MohummedSoliman/booking/repository"
	dbrepo "github.com/MohummedSoliman/booking/repository/dbRepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewRepository(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
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
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]any)
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, "make-reservation.page.html")
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	// example for date want to cenvert the string to it.
	layout := "2006-01-02"

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomId, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomId,
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.RenderTemplate(w, r, &models.TemplateData{
			Form: form,
			Data: data,
		}, "make-reservation.page.html")
		return
	}

	resID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restirction := models.RoomRestriction{
		ReservationID: resID,
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomId,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restirction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		m.App.ErrorLog.Println("Can not get error session")
		m.App.Session.Put(r.Context(), "error", "Can not get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]any)
	data["reservation"] = reservation

	render.RenderTemplate(w, r, &models.TemplateData{
		Data: data,
	}, "reservation-summary.page.html")
}
