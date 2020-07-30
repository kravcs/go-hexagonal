package class

import (
	"errors"
	"fmt"
	"time"

	u "hexa/internal/user"

	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrNewClassInvalid = errors.New("request data for Class creation is invalid")
)

type ClassService interface {
	CreateClass(newClass *NewClass) error
	FindAllClasses() (ClassesList, error)
	FindClassByName(name string) (*Class, error)
}

type classService struct {
	cr ClassRepository
	ur u.UserRepository
}

func NewClassService(cr ClassRepository, ur u.UserRepository) ClassService {
	return &classService{
		cr,
		ur,
	}
}

func (c *classService) CreateClass(newClass *NewClass) error {
	if err := validate.Validate(newClass); err != nil {
		return fmt.Errorf("class validation error: %v", err)
	}

	owner, err := c.ur.FindByName(newClass.OwnerName)
	if err != nil {
		return fmt.Errorf("class creation error: %v", err)
	}
	if owner.Role != "owner" {
		return fmt.Errorf("invalid role to create a class")
	}

	var layoutISO = "2006-01-02"
	startDate, err := time.Parse(layoutISO, newClass.StartDate)
	if err != nil {
		return fmt.Errorf("invalid startDate of a class: %v", err)
	}
	endDate, err := time.Parse(layoutISO, newClass.EndDate)
	if err != nil {
		return fmt.Errorf("invalid endDate of a class: %v", err)
	}

	class := &Class{
		Owner:     owner.Name,
		Name:      newClass.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  newClass.Capacity,
	}

	return c.cr.Create(class)
}

func (c *classService) FindAllClasses() (ClassesList, error) {
	return c.cr.FindAll()
}

func (c *classService) FindClassByName(name string) (*Class, error) {
	class, err := c.cr.FindByName(name)
	if err != nil {
		return nil, err
	}
	return class, nil
}
