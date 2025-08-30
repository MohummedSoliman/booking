package dbrepo

import (
	"context"
	"time"

	"github.com/MohummedSoliman/booking/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(r models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date,
				  room_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, statement, r.FirstName, r.LastName, r.Email, r.Phone, r.StartData,
		r.EndDate, r.RoomID, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
