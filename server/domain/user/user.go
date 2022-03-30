package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserName = errors.New("user missing a name")
	ErrNoDateOfBirth   = errors.New("user missing date of birth")
)

// User represents a user of the bil project.
type User struct {
	person *Person

	// trips is a slice of references to the user's trips
	trips []uuid.UUID
}

// New creates a new `User`.
func New(name string, dob time.Time) (User, error) {
	if name == "" {
		return User{}, ErrInvalidUserName
	}

	if dob.Equal(time.Time{}) {
		return User{}, ErrNoDateOfBirth
	}

	p := &Person{
		ID:          uuid.New(),
		Name:        name,
		DateOfBirth: dob,
	}

	return User{person: p}, nil
}

// GetID returns the ID of the User.
func (u *User) GetID() uuid.UUID {
	return u.person.ID
}

// IsYearsOld returns true if the `Person` is at least `years` old.
func (u *User) IsYearsOld(years int) bool {
	return time.Now().After(u.person.DateOfBirth.AddDate(years, 0, 0))
}
