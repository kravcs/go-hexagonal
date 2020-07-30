package class_test

import (
	"testing"
	"time"

	cl "hexa/internal/class"
	mocks "hexa/internal/mocks"
	u "hexa/internal/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	var layoutISO = "2006-01-02"
	nc := &cl.NewClass{
		Name:      "Test Name",
		StartDate: "2020-01-01",
		EndDate:   "2020-02-01",
		Capacity:  10,
		OwnerName: "Test Owner",
	}
	t1, _ := time.Parse(layoutISO, nc.StartDate)
	t2, _ := time.Parse(layoutISO, nc.EndDate)
	class := &cl.Class{
		Owner:     "Test Owner",
		Name:      "Test Name",
		StartDate: t1,
		EndDate:   t2,
		Capacity:  10,
	}

	user := &u.User{
		ID:    uuid.New().String(),
		Name:  "Test Owner",
		Email: "testowner@example.com",
		Role:  "owner",
	}

	t.Run("CreateClass", func(t *testing.T) {

		classRepository := mocks.ClassRepository{}
		classRepository.On("Create", class).Return(nil)

		userRepository := mocks.UserRepository{}
		userRepository.On("FindByName", nc.OwnerName).Return(user, nil)

		service := cl.NewClassService(&classRepository, &userRepository)
		err := service.CreateClass(nc)

		assert.Nil(t, err)
		assert.NotEmpty(t, class.Capacity)
	})

	t.Run("FindClassByName", func(t *testing.T) {
		classRepository := mocks.ClassRepository{}
		classRepository.On("FindByName", "Test Name").Return(class, nil)
		userRepository := mocks.UserRepository{}
		userRepository.On("FindByName", nc.OwnerName).Return(user, nil)
		service := cl.NewClassService(&classRepository, &userRepository)
		classResult, err := service.FindClassByName("Test Name")
		assert.Nil(t, err)
		assert.Equal(t, class, classResult)
	})
}
