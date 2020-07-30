package class

import (
	"encoding/json"
	"fmt"
	"hexa/internal/common"
	"net/http"
)

type ClassHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type classHandler struct {
	classService ClassService
}

func NewClassHandler(classService ClassService) ClassHandler {
	return &classHandler{
		classService,
	}
}

func (c *classHandler) Create(w http.ResponseWriter, r *http.Request) {

	newClass := NewClass{}
	err := json.NewDecoder(r.Body).Decode(&newClass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	err = c.classService.CreateClass(&newClass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to create class: %v", err)
		return
	}

	common.WriteJSON(w, &newClass, http.StatusCreated)
}

func (c *classHandler) Get(w http.ResponseWriter, r *http.Request) {

	classes, err := c.classService.FindAllClasses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get classes: %v", err)
		return
	}

	common.WriteJSON(w, &classes, http.StatusOK)
}
