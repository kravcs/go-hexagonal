package memory

import (
	bk "hexa/internal/booking"
	"sync"

	"github.com/google/uuid"
)

type bookingRepository struct {
	bookings bk.BookingsList
	mu       *sync.Mutex
}

func NewMemoryBookingRepository() bk.BookingRepository {
	return &bookingRepository{
		mu: &sync.Mutex{},
	}
}

func (b *bookingRepository) Create(booking *bk.Booking) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	booking.ID = uuid.New().String()

	b.bookings.Bookings = append(b.bookings.Bookings, booking)
	return nil
}

func (b *bookingRepository) FindAll() (bk.BookingsList, error) {
	bookings := bk.BookingsList{
		Bookings: append([]*bk.Booking(nil), b.bookings.Bookings...),
	}
	return bookings, nil
}
