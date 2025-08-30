package models

import (
	"time"
)

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restrictions struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartData time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

type RoomRestrictions struct {
	ID            int
	StartData     time.Time
	EndDate       time.Time
	RoomID        int
	RestrictionID int
	ReservationID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Restrictions
	Restriction   Restrictions
}
