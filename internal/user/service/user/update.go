package user

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/utils"
)

// update existing user
func (s *userssrvc) Update(ctx context.Context, p *users.UpdatePayload) (err error) {
	var userObj = &models.User{}
	// check if the user is the same as the login one
	if p.ID != int64(s.auth.GetUserID()) {
		return fmt.Errorf("invalid credential")
	}
	s.db.Rest()
	count, err := s.db.Model(userObj).
		Where(fmt.Sprintf(" lower(trim(%s)) like lower(trim('%s'))", userI.Name(), p.User.UserName)).
		Where(fmt.Sprintf(" %s != %d", userI.ID(), p.ID)).
		Count()
	if err != nil {
		return err
	}
	if *count > 0 {
		return fmt.Errorf("This user exists before")
	}
	// check the username
	if err = utils.ValidString(p.User.UserName); err != nil {
		return err
	}
	// check the password
	if err = utils.ValidPassword(p.User.UserPassword); err != nil {
		return err
	}
	// get the user
	s.db.Rest()
	err = s.db.Model(userObj).
		Where(fmt.Sprintf(" %s = %d", userI.ID(), p.ID)).
		Find(userObj)
	if err != nil {
		return nil
	}
	if userObj == nil {
		return fmt.Errorf("Id not found")
	}
	if userObj.UserID == 0 {
		return fmt.Errorf("Id not found")
	}
	// update the user data
	userObj.UserName = utils.PrepareString(&p.User.UserName, true)
	userObj.UserPassword = utils.ConvertToSha(p.User.UserPassword)
	userObj.UserRole, err = userI.ConvertStringToRole(p.User.UserRole)
	if err != nil {
		return err
	}
	s.db.Save(userObj)
	return
}
