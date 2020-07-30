package class

import (
	"time"
)

// Class is a class definition
type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
	Owner     string    `json:"-"`
}

// NewClass is a representation of class creation request
type NewClass struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity" validate:"gt=0"`
	OwnerName string `json:"owner"`
}

// ClassesList is a collection of classes
type ClassesList struct {
	Classes []*Class
}
