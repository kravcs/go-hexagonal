package booking

type BookingRepository interface {
	Create(booking *Booking) error
	FindAll() (BookingsList, error)
}
