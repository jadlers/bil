package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("person: could not create new person")
)

type Person struct {
	// ID is the identifier of the person
	ID uuid.UUID

	// Name is the name of the person
	Name string

	// DateOfBirth is the birthdate of the Person
	DateOfBirth time.Time
}
