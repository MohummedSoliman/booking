// Package repository is interface contract for methods convention
package repository

import "github.com/MohummedSoliman/booking/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(r models.Reservation) error
}
