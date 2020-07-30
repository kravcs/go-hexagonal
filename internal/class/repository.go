package class

type ClassRepository interface {
	Create(class *Class) error
	FindAll() (ClassesList, error)
	FindByName(name string) (*Class, error)
}
