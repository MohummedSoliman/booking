package main

import (
	"net/http"

	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Majors)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)

	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
