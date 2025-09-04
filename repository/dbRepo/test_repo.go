package dbrepo

import (
	"errors"
	"time"

	"github.com/MohummedSoliman/booking/pkg/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(r models.Reservation) (int, error) {
	return 1, nil
}

func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

func (m *testDBRepo) SearchAvailabilityByDate(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns slice of available rooms if any for given date range.
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

// GetUserByID return user by ID.
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}

// UpdateUser update user data
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate authenticate user.
func (m *testDBRepo) Authenticate(email, password string) (int, string, error) {
	return 0, "", nil
}

func (m *testDBRepo) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}

func (m *testDBRepo) UpdateReservation(updatedRes models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservationByID(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessed(id, proccess int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}
