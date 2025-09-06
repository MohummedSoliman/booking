// Package repository is interface contract for methods convention
package repository

import (
	"time"

	"github.com/MohummedSoliman/booking/pkg/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(r models.Reservation) (int, error)

	InsertRoomRestriction(r models.RoomRestriction) error

	SearchAvailabilityByDate(start, end time.Time, roomID int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)

	GetRoomByID(id int) (models.Room, error)

	GetUserByID(id int) (models.User, error)

	UpdateUser(u models.User) error

	Authenticate(email, password string) (int, string, error)

	GetAllReservations() ([]models.Reservation, error)

	AllNewReservations() ([]models.Reservation, error)

	GetReservationByID(id int) (models.Reservation, error)

	UpdateReservation(updatedRes models.Reservation) error

	DeleteReservationByID(id int) error

	UpdateProcessed(id, proccess int) error

	AllRooms() ([]models.Room, error)

	GetRestrictionsForRoomsByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
}
