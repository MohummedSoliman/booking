package dbrepo

import (
	"context"
	"time"

	"github.com/MohummedSoliman/booking/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(r models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	statement := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date,
				  room_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				  returning id`

	err := m.DB.QueryRowContext(ctx, statement, r.FirstName, r.LastName, r.Email, r.Phone, r.StartDate,
		r.EndDate, r.RoomID, time.Now(), time.Now()).
		Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id,
				  restriction_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, statement, r.StartDate, r.EndDate, r.RoomID, r.ReservationID,
		r.RestrictionID, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) SearchAvailabilityByDate(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `SELECT COUNT(id) FROM room_restrictions WHERE $1 < end_date and $2 > start_date and room_id = $3`
	row := m.DB.QueryRowContext(ctx, query, start, end, roomID)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}
