package memory

import (
	"fmt"
	"glofox/internal/user"

	"github.com/google/uuid"
)

type userRepository struct {
}

func NewMemoryUserRepository() user.UserRepository {
	return &userRepository{}
}

// userList is a hard coded list of users
var userList = []*user.User{
	&user.User{
		ID:    uuid.New().String(),
		Name:  "Big Boss",
		Email: "bigboss@example.com",
		Role:  "owner",
	},
	&user.User{
		ID:    uuid.New().String(),
		Name:  "VIP Member",
		Email: "vipmember@example.com",
		Role:  "member",
	},
}

var ErrUserNotFound = fmt.Errorf("User not found")

func (c *userRepository) FindByName(name string) (*user.User, error) {
	for _, u := range userList {
		if u.Name == name {
			return u, nil
		}
	}

	return nil, ErrUserNotFound
}
