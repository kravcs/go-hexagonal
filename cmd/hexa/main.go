package main

import (
	"fmt"
	"glofox/internal/booking"
	"glofox/internal/class"
	"glofox/internal/storage/memory"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	userRepo := memory.NewMemoryUserRepository()

	classRepo := memory.NewMemoryClassRepository()
	classService := class.NewClassService(classRepo, userRepo)
	classHandler := class.NewClassHandler(classService)

	r.HandleFunc("/classes", classHandler.Get).Methods("GET")
	r.HandleFunc("/classes", classHandler.Create).Methods("POST")

	bookingRepo := memory.NewMemoryBookingRepository()
	bookingService := booking.NewBookingService(bookingRepo, classRepo, userRepo)
	bookingHandler := booking.NewBookingHandler(bookingService)
	r.HandleFunc("/bookings", bookingHandler.Create).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    httpPort(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started!")
	log.Fatal(srv.ListenAndServe())
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
