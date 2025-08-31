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
}
