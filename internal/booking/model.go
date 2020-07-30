package booking

import (
	"time"
)

type Booking struct {
	ID     string
	Class  string
	Member string
	Date   time.Time
}

// NewClass is a representation of class creation request
type NewBooking struct {
	BookingDate string `json:"booking_date"`
	ClassName   string `json:"class_name"`
	MemberName  string `json:"member_name"`
}

// ClassesList is a collection of classes
type BookingsList struct {
	Bookings []*Booking
}