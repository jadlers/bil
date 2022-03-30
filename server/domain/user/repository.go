package user

import "github.com/google/uuid"

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(uuid.UUID) (User, error)
	Create(User) error

	// TODO: Add the following
	// Update(User) error
	// Remove(User) error
}
