package class_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	cl "hexa/internal/class"
	mocks "hexa/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	nc := &cl.NewClass{
		Name:      "Test Name",
		StartDate: "2020-01-01",
		EndDate:   "2020-02-01",
		Capacity:  10,
		OwnerName: "Test Owner",
	}

	t.Run("Post with json", func(t *testing.T) {
		// setup
		body, err := json.Marshal(&nc)
		assert.Nil(t, err)

		classService := mocks.ClassService{}
		classService.On("CreateClass", nc).Return(nil)
		handler := cl.NewClassHandler(&classService)
		rr := httptest.NewRecorder()

		// action
		req := httptest.NewRequest("POST", "/classes", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		handler.Create(rr, req)
		responseBody, _ := ioutil.ReadAll(rr.Body)

		// assertion
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, body, responseBody)
	})

	t.Run("Post with wrong content", func(t *testing.T) {
		nc := &cl.NewClass{}
		body, err := json.Marshal(nc)
		assert.Nil(t, err)

		classService := mocks.ClassService{}
		classService.On("CreateClass", nc).Return(errors.New("Service Error"))
		handler := cl.NewClassHandler(&classService)
		rr := httptest.NewRecorder()

		req := httptest.NewRequest("POST", "/classes", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		handler.Create(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
