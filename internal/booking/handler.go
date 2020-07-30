package booking

import (
	"encoding/json"
	"fmt"
	"glofox/internal/common"
	"net/http"
)

type BookingHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type bookingHandler struct {
	bookingService BookingService
}

func NewBookingHandler(bookingService BookingService) BookingHandler {
	return &bookingHandler{
		bookingService,
	}
}

func (b *bookingHandler) Create(w http.ResponseWriter, r *http.Request) {

	newBooking := NewBooking{}
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	err = b.bookingService.CreateBooking(&newBooking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to create booking: %v", err)
		return
	}

	common.WriteJSON(w, &newBooking, http.StatusCreated)
}
