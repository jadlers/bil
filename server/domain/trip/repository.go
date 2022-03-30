package trip

import "github.com/google/uuid"

type TripRepository interface {
	GetAll() ([]Trip, error)
	GetByID(uuid.UUID) (Trip, error)
	Create(Trip) error

	// TODO: Add later
	// Update(Trip) error
	// Remove(Trip) error
}
