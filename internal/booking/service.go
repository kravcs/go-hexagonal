package booking

import (
	"errors"
	"fmt"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"

	c "glofox/internal/class"
	u "glofox/internal/user"
)

var (
	ErrBookingInvalid = errors.New("Booking is invalid")
)

type BookingService interface {
	CreateBooking(newBooking *NewBooking) error
	FindAllBookings() (BookingsList, error)
}

type bookingService struct {
	br BookingRepository
	cr c.ClassRepository
	ur u.UserRepository
}

func NewBookingService(br BookingRepository, cr c.ClassRepository, ur u.UserRepository) BookingService {
	return &bookingService{
		br,
		cr,
		ur,
	}
}

func (b *bookingService) CreateBooking(newBooking *NewBooking) error {
	if err := validate.Validate(newBooking); err != nil {
		return errs.Wrap(ErrBookingInvalid, "service.Booking ")
	}

	member, err := b.ur.FindByName(newBooking.MemberName)
	if err != nil {
		return fmt.Errorf("booking creation error: %v", err)
	}
	if member.Role != "member" {
		return fmt.Errorf("invalid role to create a booking")
	}

	class, err := b.cr.FindByName(newBooking.ClassName)
	if err != nil {
		return fmt.Errorf("booking creation error: %v", err)
	}

	var layoutISO = "2006-01-02"
	bookingDate, err := time.Parse(layoutISO, newBooking.BookingDate)
	if err != nil {
		return fmt.Errorf("invalid startDate of a class: %v", err)
	}
	if bookingDate.Before(class.EndDate) == false || bookingDate.After(class.StartDate) == false {
		return fmt.Errorf("choose another date")
	}

	booking := &Booking{
		Class:  class.Name,
		Member: member.Name,
		Date:   bookingDate,
	}

	return b.br.Create(booking)
}

func (b *bookingService) FindAllBookings() (BookingsList, error) {
	return b.br.FindAll()
}
