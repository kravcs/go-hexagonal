package user

type UserRepository interface {
	FindByName(name string) (*User, error)
}
