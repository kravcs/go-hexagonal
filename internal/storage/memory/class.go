package memory

import (
	"errors"
	"glofox/internal/class"
	"sync"

	"github.com/google/uuid"
)

type classRepository struct {
	classes class.ClassesList
	mu      *sync.Mutex
}

func NewMemoryClassRepository() class.ClassRepository {
	return &classRepository{
		mu: &sync.Mutex{},
	}
}

func (c *classRepository) Create(class *class.Class) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	class.ID = uuid.New().String()

	c.classes.Classes = append(c.classes.Classes, class)
	return nil
}

func (c *classRepository) FindAll() (class.ClassesList, error) {
	classes := class.ClassesList{
		Classes: append([]*class.Class(nil), c.classes.Classes...),
	}
	return classes, nil
}

func (c *classRepository) FindByName(name string) (*class.Class, error) {
	for _, cs := range c.classes.Classes {
		if cs.Name == name {
			return cs, nil
		}
	}
	return nil, errors.New("class not found")
}
