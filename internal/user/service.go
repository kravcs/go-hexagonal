package user

type UserService interface {
	FindUserByName(name string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo,
	}
}

func (u *userService) FindUserByName(name string) (*User, error) {
	user, err := u.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
