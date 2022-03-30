package trip

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jadlers/bil/server/domain/user"
)

type Trip struct {
	tripID      uuid.UUID
	date        time.Time
	description *string

	// passengers reference users by their ID
	passengers []uuid.UUID
}

// NewTrip creates a new trip with a date and description
func NewTrip(date time.Time, description string) (*Trip, error) {
	trip := Trip{
		tripID:      uuid.New(),
		passengers:  make([]uuid.UUID, 0),
		date:        date,
		description: &description,
	}

	return &trip, nil
}

// GetID returns the ID of the Trip
func (t *Trip) GetID() uuid.UUID {
	return t.tripID
}

// Description is the description of the trip, if not set the second return
// value will be false
func (t *Trip) Description() (string, bool) {
	if t.description == nil {
		return "", false
	}
	return *t.description, true
}

// SetDescription updates the date of the Trip
func (t *Trip) SetDate(date time.Time) {
	t.date = date
}

// SetDescription updates the description of the Trip
func (t *Trip) SetDescription(desc string) {
	t.description = &desc
}

// AddPassenger takes a User and adds it to the list of passengers on the Trip.
// If the User is already a passenger an error is returned.
func (t *Trip) AddPassenger(user user.User) error {
	for _, u := range t.passengers {
		if u == user.GetID() {
			return errors.New("user is already a passenger of this trip")
		}
	}

	return nil
}

// TODO: Implement
func (t *Trip) RemovePassenger(user user.User) error {
	return errors.New("TODO")
}
