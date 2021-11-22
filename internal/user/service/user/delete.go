package user

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
)

// delete existing user
func (s *userssrvc) Delete(ctx context.Context, p *users.DeletePayload) (err error) {
	var userObj = &models.User{}
	// check if the user is the same as the login one
	if p.ID != int64(s.auth.GetUserID()) {
		return fmt.Errorf("invalid credential")
	}
	s.db.Rest()
	count, err := s.db.Model(userObj).
		Where(fmt.Sprintf(" %s = %d", userI.ID(), p.ID)).
		Count()
	if err != nil {
		return err
	}
	if *count == 0 {
		return fmt.Errorf("This user is not existing")
	}
	fmt.Println(userObj)

	err = s.db.Delete(userObj, p.ID)

	return err
}
