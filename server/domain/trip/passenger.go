package trip

import (
	"github.com/google/uuid"
	"github.com/jadlers/bil/server/domain/user"
)

type Passenger struct {
	userID uuid.UUID
	paid   uint

	user *user.User
}

// LoadPassenger fetches the user from the user repository
func (p *Passenger) LoadPassenger(ur user.UserRepository) error {
	user, err := ur.GetById(p.userID)
	if err != nil {
		return err
	}

	p.user = user
	return nil
}
