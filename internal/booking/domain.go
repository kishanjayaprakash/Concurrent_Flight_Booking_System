package booking

import (
	"errors"
	"time"
)

const DefaultHoldTTL = 5 * time.Minute

var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
	ErrUnauthorized = errors.New("unauthorized")
	ErrStoreNotSupported = errors.New("store does not support this operation")
)

type Booking struct {
	ID string
	FlightID  string
	SeatID string
	UserID string
	Status string
	ExpiresAt time.Time
}

type Flight struct {
	ID string
	Origin string
	Destination string
	Rows int
	SeatsPerRow int
}

type BookingStore interface {
	Book(b Booking) (Booking, error)
	ListBookings(flightID string) []Booking
}